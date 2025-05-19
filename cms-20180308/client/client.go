// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AccessKeyGetRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId   *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AccessKeyGetRequest) String() string {
	return tea.Prettify(s)
}

func (s AccessKeyGetRequest) GoString() string {
	return s.String()
}

func (s *AccessKeyGetRequest) SetRegionId(v string) *AccessKeyGetRequest {
	s.RegionId = &v
	return s
}

func (s *AccessKeyGetRequest) SetUserId(v int64) *AccessKeyGetRequest {
	s.UserId = &v
	return s
}

type AccessKeyGetResponseBody struct {
	AccessKey    *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretKey    *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	UserId       *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AccessKeyGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AccessKeyGetResponseBody) GoString() string {
	return s.String()
}

func (s *AccessKeyGetResponseBody) SetAccessKey(v string) *AccessKeyGetResponseBody {
	s.AccessKey = &v
	return s
}

func (s *AccessKeyGetResponseBody) SetErrorCode(v int32) *AccessKeyGetResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AccessKeyGetResponseBody) SetErrorMessage(v string) *AccessKeyGetResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AccessKeyGetResponseBody) SetRequestId(v string) *AccessKeyGetResponseBody {
	s.RequestId = &v
	return s
}

func (s *AccessKeyGetResponseBody) SetSecretKey(v string) *AccessKeyGetResponseBody {
	s.SecretKey = &v
	return s
}

func (s *AccessKeyGetResponseBody) SetSuccess(v bool) *AccessKeyGetResponseBody {
	s.Success = &v
	return s
}

func (s *AccessKeyGetResponseBody) SetUserId(v int64) *AccessKeyGetResponseBody {
	s.UserId = &v
	return s
}

type AccessKeyGetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AccessKeyGetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AccessKeyGetResponse) String() string {
	return tea.Prettify(s)
}

func (s AccessKeyGetResponse) GoString() string {
	return s.String()
}

func (s *AccessKeyGetResponse) SetHeaders(v map[string]*string) *AccessKeyGetResponse {
	s.Headers = v
	return s
}

func (s *AccessKeyGetResponse) SetStatusCode(v int32) *AccessKeyGetResponse {
	s.StatusCode = &v
	return s
}

func (s *AccessKeyGetResponse) SetBody(v *AccessKeyGetResponseBody) *AccessKeyGetResponse {
	s.Body = v
	return s
}

type AddMyGroupInstancesRequest struct {
	// This parameter is required.
	GroupId   *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Instances *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddMyGroupInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMyGroupInstancesRequest) GoString() string {
	return s.String()
}

func (s *AddMyGroupInstancesRequest) SetGroupId(v int64) *AddMyGroupInstancesRequest {
	s.GroupId = &v
	return s
}

func (s *AddMyGroupInstancesRequest) SetInstances(v string) *AddMyGroupInstancesRequest {
	s.Instances = &v
	return s
}

func (s *AddMyGroupInstancesRequest) SetRegionId(v string) *AddMyGroupInstancesRequest {
	s.RegionId = &v
	return s
}

type AddMyGroupInstancesResponseBody struct {
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddMyGroupInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMyGroupInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *AddMyGroupInstancesResponseBody) SetErrorCode(v int32) *AddMyGroupInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddMyGroupInstancesResponseBody) SetErrorMessage(v string) *AddMyGroupInstancesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddMyGroupInstancesResponseBody) SetRequestId(v string) *AddMyGroupInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMyGroupInstancesResponseBody) SetSuccess(v bool) *AddMyGroupInstancesResponseBody {
	s.Success = &v
	return s
}

type AddMyGroupInstancesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMyGroupInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMyGroupInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMyGroupInstancesResponse) GoString() string {
	return s.String()
}

func (s *AddMyGroupInstancesResponse) SetHeaders(v map[string]*string) *AddMyGroupInstancesResponse {
	s.Headers = v
	return s
}

func (s *AddMyGroupInstancesResponse) SetStatusCode(v int32) *AddMyGroupInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMyGroupInstancesResponse) SetBody(v *AddMyGroupInstancesResponseBody) *AddMyGroupInstancesResponse {
	s.Body = v
	return s
}

type CreateAlarmRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// >
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	ContactGroups      *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"instanceId":"i-2zecrzcri3d6fhd2ff7j "}]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// 24
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CPUUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 1
	NotifyType *int32 `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// example:
	//
	// 60
	Period   *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// example:
	//
	// 0
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// example:
	//
	// http://www.net.cn/example_callback.htm
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s CreateAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRequest) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequest) SetComparisonOperator(v string) *CreateAlarmRequest {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateAlarmRequest) SetContactGroups(v string) *CreateAlarmRequest {
	s.ContactGroups = &v
	return s
}

func (s *CreateAlarmRequest) SetDimensions(v string) *CreateAlarmRequest {
	s.Dimensions = &v
	return s
}

func (s *CreateAlarmRequest) SetDryRun(v bool) *CreateAlarmRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAlarmRequest) SetEndTime(v int32) *CreateAlarmRequest {
	s.EndTime = &v
	return s
}

func (s *CreateAlarmRequest) SetEvaluationCount(v int32) *CreateAlarmRequest {
	s.EvaluationCount = &v
	return s
}

func (s *CreateAlarmRequest) SetMetricName(v string) *CreateAlarmRequest {
	s.MetricName = &v
	return s
}

func (s *CreateAlarmRequest) SetName(v string) *CreateAlarmRequest {
	s.Name = &v
	return s
}

func (s *CreateAlarmRequest) SetNamespace(v string) *CreateAlarmRequest {
	s.Namespace = &v
	return s
}

func (s *CreateAlarmRequest) SetNotifyType(v int32) *CreateAlarmRequest {
	s.NotifyType = &v
	return s
}

func (s *CreateAlarmRequest) SetPeriod(v int32) *CreateAlarmRequest {
	s.Period = &v
	return s
}

func (s *CreateAlarmRequest) SetRegionId(v string) *CreateAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAlarmRequest) SetSilenceTime(v int32) *CreateAlarmRequest {
	s.SilenceTime = &v
	return s
}

func (s *CreateAlarmRequest) SetStartTime(v int32) *CreateAlarmRequest {
	s.StartTime = &v
	return s
}

func (s *CreateAlarmRequest) SetStatistics(v string) *CreateAlarmRequest {
	s.Statistics = &v
	return s
}

func (s *CreateAlarmRequest) SetThreshold(v string) *CreateAlarmRequest {
	s.Threshold = &v
	return s
}

func (s *CreateAlarmRequest) SetWebhook(v string) *CreateAlarmRequest {
	s.Webhook = &v
	return s
}

type CreateAlarmResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 0c4af0f1-a864-468b-bed3-15c7deff75ee
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 910ABE4E-DC9D-4231-9DC0-C96835553327
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlarmResponseBody) SetCode(v string) *CreateAlarmResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAlarmResponseBody) SetData(v string) *CreateAlarmResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAlarmResponseBody) SetMessage(v string) *CreateAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAlarmResponseBody) SetRequestId(v string) *CreateAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlarmResponseBody) SetSuccess(v bool) *CreateAlarmResponseBody {
	s.Success = &v
	return s
}

type CreateAlarmResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmResponse) GoString() string {
	return s.String()
}

func (s *CreateAlarmResponse) SetHeaders(v map[string]*string) *CreateAlarmResponse {
	s.Headers = v
	return s
}

func (s *CreateAlarmResponse) SetStatusCode(v int32) *CreateAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlarmResponse) SetBody(v *CreateAlarmResponseBody) *CreateAlarmResponse {
	s.Body = v
	return s
}

type CreateHybridDoubleWriteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-target
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-source
	SourceNamespace *string `json:"SourceNamespace,omitempty" xml:"SourceNamespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12706766********
	SourceUserId *int64 `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12706766********
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateHybridDoubleWriteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridDoubleWriteRequest) GoString() string {
	return s.String()
}

func (s *CreateHybridDoubleWriteRequest) SetNamespace(v string) *CreateHybridDoubleWriteRequest {
	s.Namespace = &v
	return s
}

func (s *CreateHybridDoubleWriteRequest) SetSourceNamespace(v string) *CreateHybridDoubleWriteRequest {
	s.SourceNamespace = &v
	return s
}

func (s *CreateHybridDoubleWriteRequest) SetSourceUserId(v int64) *CreateHybridDoubleWriteRequest {
	s.SourceUserId = &v
	return s
}

func (s *CreateHybridDoubleWriteRequest) SetUserId(v int64) *CreateHybridDoubleWriteRequest {
	s.UserId = &v
	return s
}

type CreateHybridDoubleWriteResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FAAC3C5D-00BF-543A-9E08-FAAC3C5D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateHybridDoubleWriteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridDoubleWriteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHybridDoubleWriteResponseBody) SetCode(v string) *CreateHybridDoubleWriteResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHybridDoubleWriteResponseBody) SetMessage(v string) *CreateHybridDoubleWriteResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHybridDoubleWriteResponseBody) SetRequestId(v string) *CreateHybridDoubleWriteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHybridDoubleWriteResponseBody) SetSuccess(v bool) *CreateHybridDoubleWriteResponseBody {
	s.Success = &v
	return s
}

type CreateHybridDoubleWriteResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHybridDoubleWriteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHybridDoubleWriteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridDoubleWriteResponse) GoString() string {
	return s.String()
}

func (s *CreateHybridDoubleWriteResponse) SetHeaders(v map[string]*string) *CreateHybridDoubleWriteResponse {
	s.Headers = v
	return s
}

func (s *CreateHybridDoubleWriteResponse) SetStatusCode(v int32) *CreateHybridDoubleWriteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHybridDoubleWriteResponse) SetBody(v *CreateHybridDoubleWriteResponseBody) *CreateHybridDoubleWriteResponse {
	s.Body = v
	return s
}

type CreateMonitoringTemplateRequest struct {
	// This parameter is required.
	AlertTemplatesJson       *string `json:"AlertTemplatesJson,omitempty" xml:"AlertTemplatesJson,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HostAvailabilityTemplate *string `json:"HostAvailabilityTemplate,omitempty" xml:"HostAvailabilityTemplate,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Namespace               *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ProcessMonitorTemplates *string `json:"ProcessMonitorTemplates,omitempty" xml:"ProcessMonitorTemplates,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SystemEventTemplates    *string `json:"SystemEventTemplates,omitempty" xml:"SystemEventTemplates,omitempty"`
}

func (s CreateMonitoringTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitoringTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitoringTemplateRequest) SetAlertTemplatesJson(v string) *CreateMonitoringTemplateRequest {
	s.AlertTemplatesJson = &v
	return s
}

func (s *CreateMonitoringTemplateRequest) SetDescription(v string) *CreateMonitoringTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateMonitoringTemplateRequest) SetHostAvailabilityTemplate(v string) *CreateMonitoringTemplateRequest {
	s.HostAvailabilityTemplate = &v
	return s
}

func (s *CreateMonitoringTemplateRequest) SetName(v string) *CreateMonitoringTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateMonitoringTemplateRequest) SetNamespace(v string) *CreateMonitoringTemplateRequest {
	s.Namespace = &v
	return s
}

func (s *CreateMonitoringTemplateRequest) SetProcessMonitorTemplates(v string) *CreateMonitoringTemplateRequest {
	s.ProcessMonitorTemplates = &v
	return s
}

func (s *CreateMonitoringTemplateRequest) SetRegionId(v string) *CreateMonitoringTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMonitoringTemplateRequest) SetSystemEventTemplates(v string) *CreateMonitoringTemplateRequest {
	s.SystemEventTemplates = &v
	return s
}

type CreateMonitoringTemplateResponseBody struct {
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMonitoringTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitoringTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitoringTemplateResponseBody) SetErrorCode(v int32) *CreateMonitoringTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateMonitoringTemplateResponseBody) SetErrorMessage(v string) *CreateMonitoringTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateMonitoringTemplateResponseBody) SetId(v int64) *CreateMonitoringTemplateResponseBody {
	s.Id = &v
	return s
}

func (s *CreateMonitoringTemplateResponseBody) SetRequestId(v string) *CreateMonitoringTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMonitoringTemplateResponseBody) SetSuccess(v bool) *CreateMonitoringTemplateResponseBody {
	s.Success = &v
	return s
}

type CreateMonitoringTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMonitoringTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitoringTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitoringTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitoringTemplateResponse) SetHeaders(v map[string]*string) *CreateMonitoringTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitoringTemplateResponse) SetStatusCode(v int32) *CreateMonitoringTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitoringTemplateResponse) SetBody(v *CreateMonitoringTemplateResponseBody) *CreateMonitoringTemplateResponse {
	s.Body = v
	return s
}

type CreateMyGroupsRequest struct {
	BindUrl       *string `json:"BindUrl,omitempty" xml:"BindUrl,omitempty"`
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Options       *string `json:"Options,omitempty" xml:"Options,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceId     *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateMyGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMyGroupsRequest) GoString() string {
	return s.String()
}

func (s *CreateMyGroupsRequest) SetBindUrl(v string) *CreateMyGroupsRequest {
	s.BindUrl = &v
	return s
}

func (s *CreateMyGroupsRequest) SetContactGroups(v string) *CreateMyGroupsRequest {
	s.ContactGroups = &v
	return s
}

func (s *CreateMyGroupsRequest) SetGroupName(v string) *CreateMyGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *CreateMyGroupsRequest) SetOptions(v string) *CreateMyGroupsRequest {
	s.Options = &v
	return s
}

func (s *CreateMyGroupsRequest) SetRegionId(v string) *CreateMyGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMyGroupsRequest) SetServiceId(v int64) *CreateMyGroupsRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateMyGroupsRequest) SetType(v string) *CreateMyGroupsRequest {
	s.Type = &v
	return s
}

type CreateMyGroupsResponseBody struct {
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GroupId      *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMyGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMyGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMyGroupsResponseBody) SetErrorCode(v int32) *CreateMyGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateMyGroupsResponseBody) SetErrorMessage(v string) *CreateMyGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateMyGroupsResponseBody) SetGroupId(v int64) *CreateMyGroupsResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateMyGroupsResponseBody) SetRequestId(v string) *CreateMyGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMyGroupsResponseBody) SetSuccess(v bool) *CreateMyGroupsResponseBody {
	s.Success = &v
	return s
}

type CreateMyGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMyGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMyGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMyGroupsResponse) GoString() string {
	return s.String()
}

func (s *CreateMyGroupsResponse) SetHeaders(v map[string]*string) *CreateMyGroupsResponse {
	s.Headers = v
	return s
}

func (s *CreateMyGroupsResponse) SetStatusCode(v int32) *CreateMyGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMyGroupsResponse) SetBody(v *CreateMyGroupsResponseBody) *CreateMyGroupsResponse {
	s.Body = v
	return s
}

type CreateTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://www.aliyun.com
	Address   *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AlertIds  *string `json:"AlertIds,omitempty" xml:"AlertIds,omitempty"`
	AlertRule *string `json:"AlertRule,omitempty" xml:"AlertRule,omitempty"`
	// example:
	//
	// 1
	Interval     *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IntervalUnit *string `json:"IntervalUnit,omitempty" xml:"IntervalUnit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]
	IspCity *string `json:"IspCity,omitempty" xml:"IspCity,omitempty"`
	// example:
	//
	// {"http_method":"get","header":"xx=bb","cookie":"test=aa","time_out":30000,"match_rule":0,"response_content":"aa"}
	Options  *string `json:"Options,omitempty" xml:"Options,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aliyunTest
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 1.http
	//
	// 2.ping
	//
	// 3.tcp
	//
	// 4.udp
	//
	// 5.dns
	//
	// 6.smtp
	//
	// 7.pop3
	//
	// 8.ftp
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	Caller   *string `json:"caller,omitempty" xml:"caller,omitempty"`
}

func (s CreateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskRequest) SetAddress(v string) *CreateTaskRequest {
	s.Address = &v
	return s
}

func (s *CreateTaskRequest) SetAlertIds(v string) *CreateTaskRequest {
	s.AlertIds = &v
	return s
}

func (s *CreateTaskRequest) SetAlertRule(v string) *CreateTaskRequest {
	s.AlertRule = &v
	return s
}

func (s *CreateTaskRequest) SetInterval(v string) *CreateTaskRequest {
	s.Interval = &v
	return s
}

func (s *CreateTaskRequest) SetIntervalUnit(v string) *CreateTaskRequest {
	s.IntervalUnit = &v
	return s
}

func (s *CreateTaskRequest) SetIspCity(v string) *CreateTaskRequest {
	s.IspCity = &v
	return s
}

func (s *CreateTaskRequest) SetOptions(v string) *CreateTaskRequest {
	s.Options = &v
	return s
}

func (s *CreateTaskRequest) SetRegionId(v string) *CreateTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTaskRequest) SetTaskName(v string) *CreateTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateTaskRequest) SetTaskType(v string) *CreateTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateTaskRequest) SetCaller(v string) *CreateTaskRequest {
	s.Caller = &v
	return s
}

type CreateTaskResponseBody struct {
	AlertRule *string `json:"AlertRule,omitempty" xml:"AlertRule,omitempty"`
	// example:
	//
	// 200
	Code             *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateResultList *CreateTaskResponseBodyCreateResultList `json:"CreateResultList,omitempty" xml:"CreateResultList,omitempty" type:"Struct"`
	Data             *string                                 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successfull
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// a4e6f550-7eac-4a13-b11f-6742aa2d42d1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBody) SetAlertRule(v string) *CreateTaskResponseBody {
	s.AlertRule = &v
	return s
}

func (s *CreateTaskResponseBody) SetCode(v string) *CreateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskResponseBody) SetCreateResultList(v *CreateTaskResponseBodyCreateResultList) *CreateTaskResponseBody {
	s.CreateResultList = v
	return s
}

func (s *CreateTaskResponseBody) SetData(v string) *CreateTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTaskResponseBody) SetMessage(v string) *CreateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskResponseBody) SetRequestId(v string) *CreateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskResponseBody) SetSuccess(v string) *CreateTaskResponseBody {
	s.Success = &v
	return s
}

type CreateTaskResponseBodyCreateResultList struct {
	Contact []*CreateTaskResponseBodyCreateResultListContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s CreateTaskResponseBodyCreateResultList) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponseBodyCreateResultList) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBodyCreateResultList) SetContact(v []*CreateTaskResponseBodyCreateResultListContact) *CreateTaskResponseBodyCreateResultList {
	s.Contact = v
	return s
}

type CreateTaskResponseBodyCreateResultListContact struct {
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateTaskResponseBodyCreateResultListContact) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponseBodyCreateResultListContact) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBodyCreateResultListContact) SetTaskId(v string) *CreateTaskResponseBodyCreateResultListContact {
	s.TaskId = &v
	return s
}

func (s *CreateTaskResponseBodyCreateResultListContact) SetTaskName(v string) *CreateTaskResponseBodyCreateResultListContact {
	s.TaskName = &v
	return s
}

type CreateTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskResponse) SetHeaders(v map[string]*string) *CreateTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskResponse) SetStatusCode(v int32) *CreateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskResponse) SetBody(v *CreateTaskResponseBody) *CreateTaskResponse {
	s.Body = v
	return s
}

type DeleteAlarmRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 576fbae7-2fd1-411a-ae13-6f09f4fafdde
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlarmRequest) SetId(v string) *DeleteAlarmRequest {
	s.Id = &v
	return s
}

func (s *DeleteAlarmRequest) SetRegionId(v string) *DeleteAlarmRequest {
	s.RegionId = &v
	return s
}

type DeleteAlarmResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A9371CD8-369D-49FA-BED9-35050A0DC6A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlarmResponseBody) SetCode(v string) *DeleteAlarmResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAlarmResponseBody) SetMessage(v string) *DeleteAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAlarmResponseBody) SetRequestId(v string) *DeleteAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlarmResponseBody) SetSuccess(v bool) *DeleteAlarmResponseBody {
	s.Success = &v
	return s
}

type DeleteAlarmResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlarmResponse) SetHeaders(v map[string]*string) *DeleteAlarmResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlarmResponse) SetStatusCode(v int32) *DeleteAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlarmResponse) SetBody(v *DeleteAlarmResponseBody) *DeleteAlarmResponse {
	s.Body = v
	return s
}

type DeleteCustomMetricRequest struct {
	// This parameter is required.
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Md5        *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UUID       *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
}

func (s DeleteCustomMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomMetricRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomMetricRequest) SetGroupId(v string) *DeleteCustomMetricRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteCustomMetricRequest) SetMd5(v string) *DeleteCustomMetricRequest {
	s.Md5 = &v
	return s
}

func (s *DeleteCustomMetricRequest) SetMetricName(v string) *DeleteCustomMetricRequest {
	s.MetricName = &v
	return s
}

func (s *DeleteCustomMetricRequest) SetRegionId(v string) *DeleteCustomMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomMetricRequest) SetUUID(v string) *DeleteCustomMetricRequest {
	s.UUID = &v
	return s
}

type DeleteCustomMetricResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteCustomMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomMetricResponseBody) SetCode(v string) *DeleteCustomMetricResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomMetricResponseBody) SetMessage(v string) *DeleteCustomMetricResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomMetricResponseBody) SetRequestId(v string) *DeleteCustomMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomMetricResponseBody) SetResult(v string) *DeleteCustomMetricResponseBody {
	s.Result = &v
	return s
}

type DeleteCustomMetricResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomMetricResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomMetricResponse) SetHeaders(v map[string]*string) *DeleteCustomMetricResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomMetricResponse) SetStatusCode(v int32) *DeleteCustomMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomMetricResponse) SetBody(v *DeleteCustomMetricResponseBody) *DeleteCustomMetricResponse {
	s.Body = v
	return s
}

type DeleteHybridDoubleWriteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-source
	SourceNamespace *string `json:"SourceNamespace,omitempty" xml:"SourceNamespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12706766********
	SourceUserId *int64 `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
}

func (s DeleteHybridDoubleWriteRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHybridDoubleWriteRequest) GoString() string {
	return s.String()
}

func (s *DeleteHybridDoubleWriteRequest) SetSourceNamespace(v string) *DeleteHybridDoubleWriteRequest {
	s.SourceNamespace = &v
	return s
}

func (s *DeleteHybridDoubleWriteRequest) SetSourceUserId(v int64) *DeleteHybridDoubleWriteRequest {
	s.SourceUserId = &v
	return s
}

type DeleteHybridDoubleWriteResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6F815BDC-9063-5417-BA88-E1BBD84BAA1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteHybridDoubleWriteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHybridDoubleWriteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHybridDoubleWriteResponseBody) SetCode(v string) *DeleteHybridDoubleWriteResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHybridDoubleWriteResponseBody) SetMessage(v string) *DeleteHybridDoubleWriteResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHybridDoubleWriteResponseBody) SetRequestId(v string) *DeleteHybridDoubleWriteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHybridDoubleWriteResponseBody) SetSuccess(v bool) *DeleteHybridDoubleWriteResponseBody {
	s.Success = &v
	return s
}

type DeleteHybridDoubleWriteResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHybridDoubleWriteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHybridDoubleWriteResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHybridDoubleWriteResponse) GoString() string {
	return s.String()
}

func (s *DeleteHybridDoubleWriteResponse) SetHeaders(v map[string]*string) *DeleteHybridDoubleWriteResponse {
	s.Headers = v
	return s
}

func (s *DeleteHybridDoubleWriteResponse) SetStatusCode(v int32) *DeleteHybridDoubleWriteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHybridDoubleWriteResponse) SetBody(v *DeleteHybridDoubleWriteResponseBody) *DeleteHybridDoubleWriteResponse {
	s.Body = v
	return s
}

type DeleteMetricRuleTargetsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	RuleId    *string   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	TargetIds []*string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty" type:"Repeated"`
}

func (s DeleteMetricRuleTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleTargetsRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTargetsRequest) SetRegionId(v string) *DeleteMetricRuleTargetsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMetricRuleTargetsRequest) SetRuleId(v string) *DeleteMetricRuleTargetsRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteMetricRuleTargetsRequest) SetTargetIds(v []*string) *DeleteMetricRuleTargetsRequest {
	s.TargetIds = v
	return s
}

type DeleteMetricRuleTargetsResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	FailIds   *DeleteMetricRuleTargetsResponseBodyFailIds `json:"FailIds,omitempty" xml:"FailIds,omitempty" type:"Struct"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMetricRuleTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTargetsResponseBody) SetCode(v string) *DeleteMetricRuleTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMetricRuleTargetsResponseBody) SetFailIds(v *DeleteMetricRuleTargetsResponseBodyFailIds) *DeleteMetricRuleTargetsResponseBody {
	s.FailIds = v
	return s
}

func (s *DeleteMetricRuleTargetsResponseBody) SetMessage(v string) *DeleteMetricRuleTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMetricRuleTargetsResponseBody) SetRequestId(v string) *DeleteMetricRuleTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetricRuleTargetsResponseBody) SetSuccess(v bool) *DeleteMetricRuleTargetsResponseBody {
	s.Success = &v
	return s
}

type DeleteMetricRuleTargetsResponseBodyFailIds struct {
	TargetIds *DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds `json:"TargetIds,omitempty" xml:"TargetIds,omitempty" type:"Struct"`
}

func (s DeleteMetricRuleTargetsResponseBodyFailIds) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleTargetsResponseBodyFailIds) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTargetsResponseBodyFailIds) SetTargetIds(v *DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds) *DeleteMetricRuleTargetsResponseBodyFailIds {
	s.TargetIds = v
	return s
}

type DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds struct {
	TargetId []*string `json:"TargetId,omitempty" xml:"TargetId,omitempty" type:"Repeated"`
}

func (s DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds) SetTargetId(v []*string) *DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds {
	s.TargetId = v
	return s
}

type DeleteMetricRuleTargetsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMetricRuleTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMetricRuleTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleTargetsResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTargetsResponse) SetHeaders(v map[string]*string) *DeleteMetricRuleTargetsResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetricRuleTargetsResponse) SetStatusCode(v int32) *DeleteMetricRuleTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMetricRuleTargetsResponse) SetBody(v *DeleteMetricRuleTargetsResponseBody) *DeleteMetricRuleTargetsResponse {
	s.Body = v
	return s
}

type DeleteMetricRulesRequest struct {
	// This parameter is required.
	Id       []*string `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteMetricRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetricRulesRequest) SetId(v []*string) *DeleteMetricRulesRequest {
	s.Id = v
	return s
}

func (s *DeleteMetricRulesRequest) SetRegionId(v string) *DeleteMetricRulesRequest {
	s.RegionId = &v
	return s
}

type DeleteMetricRulesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMetricRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetricRulesResponseBody) SetCode(v string) *DeleteMetricRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMetricRulesResponseBody) SetMessage(v string) *DeleteMetricRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMetricRulesResponseBody) SetRequestId(v string) *DeleteMetricRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetricRulesResponseBody) SetSuccess(v bool) *DeleteMetricRulesResponseBody {
	s.Success = &v
	return s
}

type DeleteMetricRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMetricRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMetricRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetricRulesResponse) SetHeaders(v map[string]*string) *DeleteMetricRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetricRulesResponse) SetStatusCode(v int32) *DeleteMetricRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMetricRulesResponse) SetBody(v *DeleteMetricRulesResponseBody) *DeleteMetricRulesResponse {
	s.Body = v
	return s
}

type DeleteMyGroupInstancesRequest struct {
	// This parameter is required.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// This parameter is required.
	GroupId        *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceIdList *string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	InstanceIds    *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteMyGroupInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMyGroupInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMyGroupInstancesRequest) SetCategory(v string) *DeleteMyGroupInstancesRequest {
	s.Category = &v
	return s
}

func (s *DeleteMyGroupInstancesRequest) SetGroupId(v int64) *DeleteMyGroupInstancesRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteMyGroupInstancesRequest) SetInstanceIdList(v string) *DeleteMyGroupInstancesRequest {
	s.InstanceIdList = &v
	return s
}

func (s *DeleteMyGroupInstancesRequest) SetInstanceIds(v string) *DeleteMyGroupInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DeleteMyGroupInstancesRequest) SetRegionId(v string) *DeleteMyGroupInstancesRequest {
	s.RegionId = &v
	return s
}

type DeleteMyGroupInstancesResponseBody struct {
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMyGroupInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMyGroupInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMyGroupInstancesResponseBody) SetErrorCode(v int32) *DeleteMyGroupInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMyGroupInstancesResponseBody) SetErrorMessage(v string) *DeleteMyGroupInstancesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteMyGroupInstancesResponseBody) SetRequestId(v string) *DeleteMyGroupInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMyGroupInstancesResponseBody) SetSuccess(v bool) *DeleteMyGroupInstancesResponseBody {
	s.Success = &v
	return s
}

type DeleteMyGroupInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMyGroupInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMyGroupInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMyGroupInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMyGroupInstancesResponse) SetHeaders(v map[string]*string) *DeleteMyGroupInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMyGroupInstancesResponse) SetStatusCode(v int32) *DeleteMyGroupInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMyGroupInstancesResponse) SetBody(v *DeleteMyGroupInstancesResponseBody) *DeleteMyGroupInstancesResponse {
	s.Body = v
	return s
}

type DeleteMyGroupsRequest struct {
	GroupId  *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteMyGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMyGroupsRequest) GoString() string {
	return s.String()
}

func (s *DeleteMyGroupsRequest) SetGroupId(v int64) *DeleteMyGroupsRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteMyGroupsRequest) SetRegionId(v string) *DeleteMyGroupsRequest {
	s.RegionId = &v
	return s
}

type DeleteMyGroupsResponseBody struct {
	ErrorCode    *int32                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Group        *DeleteMyGroupsResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMyGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMyGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMyGroupsResponseBody) SetErrorCode(v int32) *DeleteMyGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMyGroupsResponseBody) SetErrorMessage(v string) *DeleteMyGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteMyGroupsResponseBody) SetGroup(v *DeleteMyGroupsResponseBodyGroup) *DeleteMyGroupsResponseBody {
	s.Group = v
	return s
}

func (s *DeleteMyGroupsResponseBody) SetRequestId(v string) *DeleteMyGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMyGroupsResponseBody) SetSuccess(v bool) *DeleteMyGroupsResponseBody {
	s.Success = &v
	return s
}

type DeleteMyGroupsResponseBodyGroup struct {
	BindUrls      *string                                       `json:"BindUrls,omitempty" xml:"BindUrls,omitempty"`
	ContactGroups *DeleteMyGroupsResponseBodyGroupContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
	GroupId       *int64                                        `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string                                       `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ServiceId     *string                                       `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	Type          *string                                       `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteMyGroupsResponseBodyGroup) String() string {
	return tea.Prettify(s)
}

func (s DeleteMyGroupsResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *DeleteMyGroupsResponseBodyGroup) SetBindUrls(v string) *DeleteMyGroupsResponseBodyGroup {
	s.BindUrls = &v
	return s
}

func (s *DeleteMyGroupsResponseBodyGroup) SetContactGroups(v *DeleteMyGroupsResponseBodyGroupContactGroups) *DeleteMyGroupsResponseBodyGroup {
	s.ContactGroups = v
	return s
}

func (s *DeleteMyGroupsResponseBodyGroup) SetGroupId(v int64) *DeleteMyGroupsResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *DeleteMyGroupsResponseBodyGroup) SetGroupName(v string) *DeleteMyGroupsResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *DeleteMyGroupsResponseBodyGroup) SetServiceId(v string) *DeleteMyGroupsResponseBodyGroup {
	s.ServiceId = &v
	return s
}

func (s *DeleteMyGroupsResponseBodyGroup) SetType(v string) *DeleteMyGroupsResponseBodyGroup {
	s.Type = &v
	return s
}

type DeleteMyGroupsResponseBodyGroupContactGroups struct {
	ContactGroup []*DeleteMyGroupsResponseBodyGroupContactGroupsContactGroup `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s DeleteMyGroupsResponseBodyGroupContactGroups) String() string {
	return tea.Prettify(s)
}

func (s DeleteMyGroupsResponseBodyGroupContactGroups) GoString() string {
	return s.String()
}

func (s *DeleteMyGroupsResponseBodyGroupContactGroups) SetContactGroup(v []*DeleteMyGroupsResponseBodyGroupContactGroupsContactGroup) *DeleteMyGroupsResponseBodyGroupContactGroups {
	s.ContactGroup = v
	return s
}

type DeleteMyGroupsResponseBodyGroupContactGroupsContactGroup struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteMyGroupsResponseBodyGroupContactGroupsContactGroup) String() string {
	return tea.Prettify(s)
}

func (s DeleteMyGroupsResponseBodyGroupContactGroupsContactGroup) GoString() string {
	return s.String()
}

func (s *DeleteMyGroupsResponseBodyGroupContactGroupsContactGroup) SetName(v string) *DeleteMyGroupsResponseBodyGroupContactGroupsContactGroup {
	s.Name = &v
	return s
}

type DeleteMyGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMyGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMyGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMyGroupsResponse) GoString() string {
	return s.String()
}

func (s *DeleteMyGroupsResponse) SetHeaders(v map[string]*string) *DeleteMyGroupsResponse {
	s.Headers = v
	return s
}

func (s *DeleteMyGroupsResponse) SetStatusCode(v int32) *DeleteMyGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMyGroupsResponse) SetBody(v *DeleteMyGroupsResponseBody) *DeleteMyGroupsResponse {
	s.Body = v
	return s
}

type DeleteTasksRequest struct {
	// example:
	//
	// 1
	IsDeleteAlarms *int32  `json:"IsDeleteAlarms,omitempty" xml:"IsDeleteAlarms,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["2b5e6f7d-108f-4117-85fb-b202ba033468"]
	TaskIds *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s DeleteTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTasksRequest) GoString() string {
	return s.String()
}

func (s *DeleteTasksRequest) SetIsDeleteAlarms(v int32) *DeleteTasksRequest {
	s.IsDeleteAlarms = &v
	return s
}

func (s *DeleteTasksRequest) SetRegionId(v string) *DeleteTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTasksRequest) SetTaskIds(v string) *DeleteTasksRequest {
	s.TaskIds = &v
	return s
}

type DeleteTasksResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"count":1}
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// a4e6f550-7eac-4a13-b11f-6742aa2d42d1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTasksResponseBody) SetCode(v string) *DeleteTasksResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTasksResponseBody) SetData(v string) *DeleteTasksResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteTasksResponseBody) SetMessage(v string) *DeleteTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTasksResponseBody) SetRequestId(v string) *DeleteTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTasksResponseBody) SetSuccess(v string) *DeleteTasksResponseBody {
	s.Success = &v
	return s
}

type DeleteTasksResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTasksResponse) GoString() string {
	return s.String()
}

func (s *DeleteTasksResponse) SetHeaders(v map[string]*string) *DeleteTasksResponse {
	s.Headers = v
	return s
}

func (s *DeleteTasksResponse) SetStatusCode(v int32) *DeleteTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTasksResponse) SetBody(v *DeleteTasksResponseBody) *DeleteTasksResponse {
	s.Body = v
	return s
}

type DescribeAlarmHistoryRequest struct {
	AlertName  *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Ascending  *bool   `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Namespace  *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OnlyCount  *bool   `json:"OnlyCount,omitempty" xml:"OnlyCount,omitempty"`
	Page       *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State      *string `json:"State,omitempty" xml:"State,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAlarmHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlarmHistoryRequest) SetAlertName(v string) *DescribeAlarmHistoryRequest {
	s.AlertName = &v
	return s
}

func (s *DescribeAlarmHistoryRequest) SetAscending(v bool) *DescribeAlarmHistoryRequest {
	s.Ascending = &v
	return s
}

func (s *DescribeAlarmHistoryRequest) SetEndTime(v string) *DescribeAlarmHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlarmHistoryRequest) SetGroupId(v string) *DescribeAlarmHistoryRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAlarmHistoryRequest) SetMetricName(v string) *DescribeAlarmHistoryRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeAlarmHistoryRequest) SetNamespace(v string) *DescribeAlarmHistoryRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeAlarmHistoryRequest) SetOnlyCount(v bool) *DescribeAlarmHistoryRequest {
	s.OnlyCount = &v
	return s
}

func (s *DescribeAlarmHistoryRequest) SetPage(v int32) *DescribeAlarmHistoryRequest {
	s.Page = &v
	return s
}

func (s *DescribeAlarmHistoryRequest) SetPageSize(v int32) *DescribeAlarmHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlarmHistoryRequest) SetRegionId(v string) *DescribeAlarmHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlarmHistoryRequest) SetRuleName(v string) *DescribeAlarmHistoryRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeAlarmHistoryRequest) SetStartTime(v string) *DescribeAlarmHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlarmHistoryRequest) SetState(v string) *DescribeAlarmHistoryRequest {
	s.State = &v
	return s
}

func (s *DescribeAlarmHistoryRequest) SetStatus(v string) *DescribeAlarmHistoryRequest {
	s.Status = &v
	return s
}

type DescribeAlarmHistoryResponseBody struct {
	AlarmHistoryList *DescribeAlarmHistoryResponseBodyAlarmHistoryList `json:"AlarmHistoryList,omitempty" xml:"AlarmHistoryList,omitempty" type:"Struct"`
	Code             *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message          *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Total            *string                                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAlarmHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlarmHistoryResponseBody) SetAlarmHistoryList(v *DescribeAlarmHistoryResponseBodyAlarmHistoryList) *DescribeAlarmHistoryResponseBody {
	s.AlarmHistoryList = v
	return s
}

func (s *DescribeAlarmHistoryResponseBody) SetCode(v string) *DescribeAlarmHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBody) SetMessage(v string) *DescribeAlarmHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBody) SetRequestId(v string) *DescribeAlarmHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBody) SetSuccess(v bool) *DescribeAlarmHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBody) SetTotal(v string) *DescribeAlarmHistoryResponseBody {
	s.Total = &v
	return s
}

type DescribeAlarmHistoryResponseBodyAlarmHistoryList struct {
	AlarmHistory []*DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory `json:"AlarmHistory,omitempty" xml:"AlarmHistory,omitempty" type:"Repeated"`
}

func (s DescribeAlarmHistoryResponseBodyAlarmHistoryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmHistoryResponseBodyAlarmHistoryList) GoString() string {
	return s.String()
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryList) SetAlarmHistory(v []*DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) *DescribeAlarmHistoryResponseBodyAlarmHistoryList {
	s.AlarmHistory = v
	return s
}

type DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory struct {
	AlertName       *string                                                                    `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	AlertTime       *int64                                                                     `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	ContactALIIMs   *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs `json:"ContactALIIMs,omitempty" xml:"ContactALIIMs,omitempty" type:"Struct"`
	ContactGroups   *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
	ContactMails    *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactMails  `json:"ContactMails,omitempty" xml:"ContactMails,omitempty" type:"Struct"`
	ContactSmses    *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactSmses  `json:"ContactSmses,omitempty" xml:"ContactSmses,omitempty" type:"Struct"`
	Contacts        *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContacts      `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Struct"`
	Dimensions      *string                                                                    `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EvaluationCount *int32                                                                     `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Expression      *string                                                                    `json:"Expression,omitempty" xml:"Expression,omitempty"`
	GroupId         *string                                                                    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id              *string                                                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceName    *string                                                                    `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	LastTime        *int64                                                                     `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	Level           *string                                                                    `json:"Level,omitempty" xml:"Level,omitempty"`
	MetricName      *string                                                                    `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Namespace       *string                                                                    `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PreLevel        *string                                                                    `json:"PreLevel,omitempty" xml:"PreLevel,omitempty"`
	RuleName        *string                                                                    `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	State           *string                                                                    `json:"State,omitempty" xml:"State,omitempty"`
	Status          *int32                                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId          *string                                                                    `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Value           *string                                                                    `json:"Value,omitempty" xml:"Value,omitempty"`
	Webhooks        *string                                                                    `json:"Webhooks,omitempty" xml:"Webhooks,omitempty"`
}

func (s DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) GoString() string {
	return s.String()
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetAlertName(v string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.AlertName = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetAlertTime(v int64) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.AlertTime = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetContactALIIMs(v *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactALIIMs = v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetContactGroups(v *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactGroups) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactGroups = v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetContactMails(v *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactMails) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactMails = v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetContactSmses(v *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactSmses) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactSmses = v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetContacts(v *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContacts) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.Contacts = v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetDimensions(v string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.Dimensions = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetEvaluationCount(v int32) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetExpression(v string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.Expression = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetGroupId(v string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.GroupId = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetId(v string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.Id = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetInstanceName(v string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.InstanceName = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetLastTime(v int64) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.LastTime = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetLevel(v string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.Level = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetMetricName(v string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.MetricName = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetNamespace(v string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.Namespace = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetPreLevel(v string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.PreLevel = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetRuleName(v string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.RuleName = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetState(v string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.State = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetStatus(v int32) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.Status = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetUserId(v string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.UserId = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetValue(v string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.Value = &v
	return s
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetWebhooks(v string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.Webhooks = &v
	return s
}

type DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs struct {
	ContactALIIM []*string `json:"ContactALIIM,omitempty" xml:"ContactALIIM,omitempty" type:"Repeated"`
}

func (s DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) GoString() string {
	return s.String()
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) SetContactALIIM(v []*string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs {
	s.ContactALIIM = v
	return s
}

type DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactGroups struct {
	ContactGroup []*string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactGroups) GoString() string {
	return s.String()
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactGroups) SetContactGroup(v []*string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactGroups {
	s.ContactGroup = v
	return s
}

type DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactMails struct {
	ContactMail []*string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty" type:"Repeated"`
}

func (s DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactMails) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactMails) GoString() string {
	return s.String()
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactMails) SetContactMail(v []*string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactMails {
	s.ContactMail = v
	return s
}

type DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactSmses struct {
	ContactSms []*string `json:"ContactSms,omitempty" xml:"ContactSms,omitempty" type:"Repeated"`
}

func (s DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactSmses) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactSmses) GoString() string {
	return s.String()
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactSmses) SetContactSms(v []*string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContactSmses {
	s.ContactSms = v
	return s
}

type DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContacts struct {
	Contact []*string `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContacts) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContacts) GoString() string {
	return s.String()
}

func (s *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContacts) SetContact(v []*string) *DescribeAlarmHistoryResponseBodyAlarmHistoryListAlarmHistoryContacts {
	s.Contact = v
	return s
}

type DescribeAlarmHistoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlarmHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlarmHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlarmHistoryResponse) SetHeaders(v map[string]*string) *DescribeAlarmHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlarmHistoryResponse) SetStatusCode(v int32) *DescribeAlarmHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlarmHistoryResponse) SetBody(v *DescribeAlarmHistoryResponseBody) *DescribeAlarmHistoryResponse {
	s.Body = v
	return s
}

type DescribeAlarmsRequest struct {
	AlertState  *string `json:"AlertState,omitempty" xml:"AlertState,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EnableState *bool   `json:"EnableState,omitempty" xml:"EnableState,omitempty"`
	GroupBy     *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MetricName  *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	NameKeyword *string `json:"NameKeyword,omitempty" xml:"NameKeyword,omitempty"`
	Names       *string `json:"Names,omitempty" xml:"Names,omitempty"`
	Namespace   *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Page        *string `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAlarmsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsRequest) SetAlertState(v string) *DescribeAlarmsRequest {
	s.AlertState = &v
	return s
}

func (s *DescribeAlarmsRequest) SetDisplayName(v string) *DescribeAlarmsRequest {
	s.DisplayName = &v
	return s
}

func (s *DescribeAlarmsRequest) SetEnableState(v bool) *DescribeAlarmsRequest {
	s.EnableState = &v
	return s
}

func (s *DescribeAlarmsRequest) SetGroupBy(v string) *DescribeAlarmsRequest {
	s.GroupBy = &v
	return s
}

func (s *DescribeAlarmsRequest) SetGroupId(v string) *DescribeAlarmsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAlarmsRequest) SetMetricName(v string) *DescribeAlarmsRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeAlarmsRequest) SetNameKeyword(v string) *DescribeAlarmsRequest {
	s.NameKeyword = &v
	return s
}

func (s *DescribeAlarmsRequest) SetNames(v string) *DescribeAlarmsRequest {
	s.Names = &v
	return s
}

func (s *DescribeAlarmsRequest) SetNamespace(v string) *DescribeAlarmsRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeAlarmsRequest) SetPage(v string) *DescribeAlarmsRequest {
	s.Page = &v
	return s
}

func (s *DescribeAlarmsRequest) SetPageSize(v string) *DescribeAlarmsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlarmsRequest) SetRegionId(v string) *DescribeAlarmsRequest {
	s.RegionId = &v
	return s
}

type DescribeAlarmsResponseBody struct {
	Code       *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Datapoints *DescribeAlarmsResponseBodyDatapoints `json:"Datapoints,omitempty" xml:"Datapoints,omitempty" type:"Struct"`
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Total      *string                               `json:"Total,omitempty" xml:"Total,omitempty"`
	TraceId    *string                               `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeAlarmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBody) SetCode(v int32) *DescribeAlarmsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetDatapoints(v *DescribeAlarmsResponseBodyDatapoints) *DescribeAlarmsResponseBody {
	s.Datapoints = v
	return s
}

func (s *DescribeAlarmsResponseBody) SetMessage(v string) *DescribeAlarmsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetRequestId(v string) *DescribeAlarmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetSuccess(v bool) *DescribeAlarmsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetTotal(v string) *DescribeAlarmsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetTraceId(v string) *DescribeAlarmsResponseBody {
	s.TraceId = &v
	return s
}

type DescribeAlarmsResponseBodyDatapoints struct {
	Alarm []*DescribeAlarmsResponseBodyDatapointsAlarm `json:"Alarm,omitempty" xml:"Alarm,omitempty" type:"Repeated"`
}

func (s DescribeAlarmsResponseBodyDatapoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBodyDatapoints) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyDatapoints) SetAlarm(v []*DescribeAlarmsResponseBodyDatapointsAlarm) *DescribeAlarmsResponseBodyDatapoints {
	s.Alarm = v
	return s
}

type DescribeAlarmsResponseBodyDatapointsAlarm struct {
	ComparisonOperator  *string                                               `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	ContactGroups       *string                                               `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	Dimensions          *string                                               `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	DisplayName         *string                                               `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EffectiveInterval   *string                                               `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	Enable              *bool                                                 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Escalations         *DescribeAlarmsResponseBodyDatapointsAlarmEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	EvaluationCount     *string                                               `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	GroupId             *string                                               `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName           *string                                               `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Level               *string                                               `json:"Level,omitempty" xml:"Level,omitempty"`
	MetricName          *string                                               `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Name                *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace           *string                                               `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NoEffectiveInterval *string                                               `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	Period              *string                                               `json:"Period,omitempty" xml:"Period,omitempty"`
	Resources           *string                                               `json:"Resources,omitempty" xml:"Resources,omitempty"`
	RuleName            *string                                               `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType            *string                                               `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	SilenceTime         *string                                               `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	State               *string                                               `json:"State,omitempty" xml:"State,omitempty"`
	Statistics          *string                                               `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Subject             *string                                               `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Threshold           *string                                               `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Uuid                *string                                               `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Webhook             *string                                               `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s DescribeAlarmsResponseBodyDatapointsAlarm) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBodyDatapointsAlarm) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetComparisonOperator(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetContactGroups(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.ContactGroups = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetDimensions(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.Dimensions = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetDisplayName(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.DisplayName = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetEffectiveInterval(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.EffectiveInterval = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetEnable(v bool) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.Enable = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetEscalations(v *DescribeAlarmsResponseBodyDatapointsAlarmEscalations) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.Escalations = v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetEvaluationCount(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetGroupId(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.GroupId = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetGroupName(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.GroupName = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetLevel(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.Level = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetMetricName(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.MetricName = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetName(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.Name = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetNamespace(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.Namespace = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetNoEffectiveInterval(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.NoEffectiveInterval = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetPeriod(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.Period = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetResources(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.Resources = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetRuleName(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.RuleName = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetRuleType(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.RuleType = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetSilenceTime(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.SilenceTime = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetState(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.State = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetStatistics(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.Statistics = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetSubject(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.Subject = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetThreshold(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.Threshold = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetUuid(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.Uuid = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarm) SetWebhook(v string) *DescribeAlarmsResponseBodyDatapointsAlarm {
	s.Webhook = &v
	return s
}

type DescribeAlarmsResponseBodyDatapointsAlarmEscalations struct {
	Critical *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s DescribeAlarmsResponseBodyDatapointsAlarmEscalations) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBodyDatapointsAlarmEscalations) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalations) SetCritical(v *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical) *DescribeAlarmsResponseBodyDatapointsAlarmEscalations {
	s.Critical = v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalations) SetInfo(v *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo) *DescribeAlarmsResponseBodyDatapointsAlarmEscalations {
	s.Info = v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalations) SetWarn(v *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn) *DescribeAlarmsResponseBodyDatapointsAlarmEscalations {
	s.Warn = v
	return s
}

type DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical struct {
	AlertSensitivity   *string `json:"AlertSensitivity,omitempty" xml:"AlertSensitivity,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	HistoryDataRange   *string `json:"HistoryDataRange,omitempty" xml:"HistoryDataRange,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *string `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical) SetAlertSensitivity(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical {
	s.AlertSensitivity = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical) SetComparisonOperator(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical) SetHistoryDataRange(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical {
	s.HistoryDataRange = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical) SetPreCondition(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical {
	s.PreCondition = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical) SetStatistics(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical) SetThreshold(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical) SetTimes(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsCritical {
	s.Times = &v
	return s
}

type DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo struct {
	AlertSensitivity   *string `json:"AlertSensitivity,omitempty" xml:"AlertSensitivity,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	HistoryDataRange   *string `json:"HistoryDataRange,omitempty" xml:"HistoryDataRange,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *string `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo) SetAlertSensitivity(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo {
	s.AlertSensitivity = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo) SetComparisonOperator(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo) SetHistoryDataRange(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo {
	s.HistoryDataRange = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo) SetPreCondition(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo {
	s.PreCondition = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo) SetStatistics(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo) SetThreshold(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo) SetTimes(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsInfo {
	s.Times = &v
	return s
}

type DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn struct {
	AlertSensitivity   *string `json:"AlertSensitivity,omitempty" xml:"AlertSensitivity,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	HistoryDataRange   *string `json:"HistoryDataRange,omitempty" xml:"HistoryDataRange,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *string `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn) SetAlertSensitivity(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn {
	s.AlertSensitivity = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn) SetComparisonOperator(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn) SetHistoryDataRange(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn {
	s.HistoryDataRange = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn) SetPreCondition(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn {
	s.PreCondition = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn) SetStatistics(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn) SetThreshold(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn) SetTimes(v string) *DescribeAlarmsResponseBodyDatapointsAlarmEscalationsWarn {
	s.Times = &v
	return s
}

type DescribeAlarmsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlarmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlarmsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponse) SetHeaders(v map[string]*string) *DescribeAlarmsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlarmsResponse) SetStatusCode(v int32) *DescribeAlarmsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlarmsResponse) SetBody(v *DescribeAlarmsResponseBody) *DescribeAlarmsResponse {
	s.Body = v
	return s
}

type DescribeAlarmsForResourcesRequest struct {
	AlertState *string `json:"AlertState,omitempty" xml:"AlertState,omitempty"`
	// This parameter is required.
	Dimensions  *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EnableState *bool   `json:"EnableState,omitempty" xml:"EnableState,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MetricName  *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// This parameter is required.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Page      *string `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize  *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAlarmsForResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsForResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsForResourcesRequest) SetAlertState(v string) *DescribeAlarmsForResourcesRequest {
	s.AlertState = &v
	return s
}

func (s *DescribeAlarmsForResourcesRequest) SetDimensions(v string) *DescribeAlarmsForResourcesRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeAlarmsForResourcesRequest) SetEnableState(v bool) *DescribeAlarmsForResourcesRequest {
	s.EnableState = &v
	return s
}

func (s *DescribeAlarmsForResourcesRequest) SetGroupId(v string) *DescribeAlarmsForResourcesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAlarmsForResourcesRequest) SetMetricName(v string) *DescribeAlarmsForResourcesRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeAlarmsForResourcesRequest) SetNamespace(v string) *DescribeAlarmsForResourcesRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeAlarmsForResourcesRequest) SetPage(v string) *DescribeAlarmsForResourcesRequest {
	s.Page = &v
	return s
}

func (s *DescribeAlarmsForResourcesRequest) SetPageSize(v string) *DescribeAlarmsForResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlarmsForResourcesRequest) SetRegionId(v string) *DescribeAlarmsForResourcesRequest {
	s.RegionId = &v
	return s
}

type DescribeAlarmsForResourcesResponseBody struct {
	Code       *int32                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Datapoints *DescribeAlarmsForResourcesResponseBodyDatapoints `json:"Datapoints,omitempty" xml:"Datapoints,omitempty" type:"Struct"`
	Dimensions *string                                           `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	Message    *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Total      *string                                           `json:"Total,omitempty" xml:"Total,omitempty"`
	TraceId    *string                                           `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeAlarmsForResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsForResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsForResourcesResponseBody) SetCode(v int32) *DescribeAlarmsForResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBody) SetDatapoints(v *DescribeAlarmsForResourcesResponseBodyDatapoints) *DescribeAlarmsForResourcesResponseBody {
	s.Datapoints = v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBody) SetDimensions(v string) *DescribeAlarmsForResourcesResponseBody {
	s.Dimensions = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBody) SetMessage(v string) *DescribeAlarmsForResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBody) SetRequestId(v string) *DescribeAlarmsForResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBody) SetSuccess(v bool) *DescribeAlarmsForResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBody) SetTotal(v string) *DescribeAlarmsForResourcesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBody) SetTraceId(v string) *DescribeAlarmsForResourcesResponseBody {
	s.TraceId = &v
	return s
}

type DescribeAlarmsForResourcesResponseBodyDatapoints struct {
	Alarm []*DescribeAlarmsForResourcesResponseBodyDatapointsAlarm `json:"Alarm,omitempty" xml:"Alarm,omitempty" type:"Repeated"`
}

func (s DescribeAlarmsForResourcesResponseBodyDatapoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsForResourcesResponseBodyDatapoints) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapoints) SetAlarm(v []*DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) *DescribeAlarmsForResourcesResponseBodyDatapoints {
	s.Alarm = v
	return s
}

type DescribeAlarmsForResourcesResponseBodyDatapointsAlarm struct {
	ComparisonOperator  *string                                                           `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	ContactGroups       *string                                                           `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	DisplayName         *string                                                           `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EffectiveInterval   *string                                                           `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	Enable              *bool                                                             `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Escalations         *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	EvaluationCount     *string                                                           `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	GroupId             *string                                                           `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName           *string                                                           `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Level               *string                                                           `json:"Level,omitempty" xml:"Level,omitempty"`
	MetricName          *string                                                           `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Name                *string                                                           `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace           *string                                                           `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NoEffectiveInterval *string                                                           `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	Period              *string                                                           `json:"Period,omitempty" xml:"Period,omitempty"`
	Resources           *string                                                           `json:"Resources,omitempty" xml:"Resources,omitempty"`
	SilenceTime         *string                                                           `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	State               *string                                                           `json:"State,omitempty" xml:"State,omitempty"`
	Statistics          *string                                                           `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Subject             *string                                                           `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Threshold           *string                                                           `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Uuid                *string                                                           `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Webhook             *string                                                           `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetComparisonOperator(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetContactGroups(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.ContactGroups = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetDisplayName(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.DisplayName = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetEffectiveInterval(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.EffectiveInterval = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetEnable(v bool) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.Enable = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetEscalations(v *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalations) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.Escalations = v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetEvaluationCount(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetGroupId(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.GroupId = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetGroupName(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.GroupName = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetLevel(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.Level = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetMetricName(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.MetricName = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetName(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.Name = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetNamespace(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.Namespace = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetNoEffectiveInterval(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.NoEffectiveInterval = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetPeriod(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.Period = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetResources(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.Resources = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetSilenceTime(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.SilenceTime = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetState(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.State = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetStatistics(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.Statistics = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetSubject(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.Subject = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetThreshold(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.Threshold = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetUuid(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.Uuid = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm) SetWebhook(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarm {
	s.Webhook = &v
	return s
}

type DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalations struct {
	Critical *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalations) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalations) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalations) SetCritical(v *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsCritical) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalations {
	s.Critical = v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalations) SetInfo(v *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsInfo) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalations {
	s.Info = v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalations) SetWarn(v *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsWarn) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalations {
	s.Warn = v
	return s
}

type DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsCritical struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *string `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsCritical) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsCritical) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsCritical) SetComparisonOperator(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsCritical) SetStatistics(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsCritical) SetThreshold(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsCritical) SetTimes(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsCritical {
	s.Times = &v
	return s
}

type DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsInfo struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *string `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsInfo) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsInfo) SetComparisonOperator(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsInfo) SetStatistics(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsInfo) SetThreshold(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsInfo) SetTimes(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsInfo {
	s.Times = &v
	return s
}

type DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsWarn struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *string `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsWarn) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsWarn) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsWarn) SetComparisonOperator(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsWarn) SetStatistics(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsWarn) SetThreshold(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsWarn) SetTimes(v string) *DescribeAlarmsForResourcesResponseBodyDatapointsAlarmEscalationsWarn {
	s.Times = &v
	return s
}

type DescribeAlarmsForResourcesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlarmsForResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlarmsForResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsForResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsForResourcesResponse) SetHeaders(v map[string]*string) *DescribeAlarmsForResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlarmsForResourcesResponse) SetStatusCode(v int32) *DescribeAlarmsForResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlarmsForResourcesResponse) SetBody(v *DescribeAlarmsForResourcesResponseBody) *DescribeAlarmsForResourcesResponse {
	s.Body = v
	return s
}

type DescribeAlertHistoryListRequest struct {
	AlertName  *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Ascending  *bool   `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Namespace  *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OnlyCount  *bool   `json:"OnlyCount,omitempty" xml:"OnlyCount,omitempty"`
	Page       *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State      *string `json:"State,omitempty" xml:"State,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAlertHistoryListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListRequest) SetAlertName(v string) *DescribeAlertHistoryListRequest {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetAscending(v bool) *DescribeAlertHistoryListRequest {
	s.Ascending = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetEndTime(v string) *DescribeAlertHistoryListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetGroupId(v string) *DescribeAlertHistoryListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetMetricName(v string) *DescribeAlertHistoryListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetNamespace(v string) *DescribeAlertHistoryListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetOnlyCount(v bool) *DescribeAlertHistoryListRequest {
	s.OnlyCount = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetPage(v int32) *DescribeAlertHistoryListRequest {
	s.Page = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetPageSize(v int32) *DescribeAlertHistoryListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetRegionId(v string) *DescribeAlertHistoryListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetRuleName(v string) *DescribeAlertHistoryListRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetStartTime(v string) *DescribeAlertHistoryListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetState(v string) *DescribeAlertHistoryListRequest {
	s.State = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetStatus(v string) *DescribeAlertHistoryListRequest {
	s.Status = &v
	return s
}

type DescribeAlertHistoryListResponseBody struct {
	AlarmHistoryList *DescribeAlertHistoryListResponseBodyAlarmHistoryList `json:"AlarmHistoryList,omitempty" xml:"AlarmHistoryList,omitempty" type:"Struct"`
	Code             *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message          *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Total            *string                                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAlertHistoryListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBody) SetAlarmHistoryList(v *DescribeAlertHistoryListResponseBodyAlarmHistoryList) *DescribeAlertHistoryListResponseBody {
	s.AlarmHistoryList = v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetCode(v string) *DescribeAlertHistoryListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetMessage(v string) *DescribeAlertHistoryListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetRequestId(v string) *DescribeAlertHistoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetSuccess(v bool) *DescribeAlertHistoryListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetTotal(v string) *DescribeAlertHistoryListResponseBody {
	s.Total = &v
	return s
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryList struct {
	AlarmHistory []*DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory `json:"AlarmHistory,omitempty" xml:"AlarmHistory,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryList) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryList) SetAlarmHistory(v []*DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) *DescribeAlertHistoryListResponseBodyAlarmHistoryList {
	s.AlarmHistory = v
	return s
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory struct {
	AlertName       *string                                                                        `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	AlertTime       *int64                                                                         `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	ContactALIIMs   *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs `json:"ContactALIIMs,omitempty" xml:"ContactALIIMs,omitempty" type:"Struct"`
	ContactGroups   *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
	ContactMails    *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails  `json:"ContactMails,omitempty" xml:"ContactMails,omitempty" type:"Struct"`
	ContactSmses    *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses  `json:"ContactSmses,omitempty" xml:"ContactSmses,omitempty" type:"Struct"`
	Contacts        *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts      `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Struct"`
	Dimensions      *string                                                                        `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EvaluationCount *int32                                                                         `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Expression      *string                                                                        `json:"Expression,omitempty" xml:"Expression,omitempty"`
	GroupId         *string                                                                        `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id              *string                                                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceName    *string                                                                        `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	LastTime        *int64                                                                         `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	Level           *string                                                                        `json:"Level,omitempty" xml:"Level,omitempty"`
	MetricName      *string                                                                        `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Namespace       *string                                                                        `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PreLevel        *string                                                                        `json:"PreLevel,omitempty" xml:"PreLevel,omitempty"`
	RuleName        *string                                                                        `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	State           *string                                                                        `json:"State,omitempty" xml:"State,omitempty"`
	Status          *int32                                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId          *string                                                                        `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Value           *string                                                                        `json:"Value,omitempty" xml:"Value,omitempty"`
	Webhooks        *string                                                                        `json:"Webhooks,omitempty" xml:"Webhooks,omitempty"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetAlertName(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetAlertTime(v int64) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.AlertTime = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContactALIIMs(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactALIIMs = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContactGroups(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactGroups = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContactMails(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactMails = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContactSmses(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactSmses = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContacts(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Contacts = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetDimensions(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Dimensions = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetEvaluationCount(v int32) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetExpression(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Expression = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetGroupId(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetId(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Id = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetInstanceName(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.InstanceName = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetLastTime(v int64) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.LastTime = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetLevel(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Level = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetMetricName(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetNamespace(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetPreLevel(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.PreLevel = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetRuleName(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.RuleName = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetState(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.State = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetStatus(v int32) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Status = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetUserId(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.UserId = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetValue(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Value = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetWebhooks(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Webhooks = &v
	return s
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs struct {
	ContactALIIM []*string `json:"ContactALIIM,omitempty" xml:"ContactALIIM,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) SetContactALIIM(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs {
	s.ContactALIIM = v
	return s
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups struct {
	ContactGroup []*string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) SetContactGroup(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups {
	s.ContactGroup = v
	return s
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails struct {
	ContactMail []*string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) SetContactMail(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails {
	s.ContactMail = v
	return s
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses struct {
	ContactSms []*string `json:"ContactSms,omitempty" xml:"ContactSms,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) SetContactSms(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses {
	s.ContactSms = v
	return s
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts struct {
	Contact []*string `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) SetContact(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts {
	s.Contact = v
	return s
}

type DescribeAlertHistoryListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertHistoryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertHistoryListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponse) SetHeaders(v map[string]*string) *DescribeAlertHistoryListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertHistoryListResponse) SetStatusCode(v int32) *DescribeAlertHistoryListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertHistoryListResponse) SetBody(v *DescribeAlertHistoryListResponseBody) *DescribeAlertHistoryListResponse {
	s.Body = v
	return s
}

type DescribeContactRequest struct {
	// This parameter is required.
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeContactRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactRequest) GoString() string {
	return s.String()
}

func (s *DescribeContactRequest) SetContactName(v string) *DescribeContactRequest {
	s.ContactName = &v
	return s
}

func (s *DescribeContactRequest) SetRegionId(v string) *DescribeContactRequest {
	s.RegionId = &v
	return s
}

type DescribeContactResponseBody struct {
	Code       *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Datapoints *DescribeContactResponseBodyDatapoints `json:"Datapoints,omitempty" xml:"Datapoints,omitempty" type:"Struct"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContactResponseBody) SetCode(v int32) *DescribeContactResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeContactResponseBody) SetDatapoints(v *DescribeContactResponseBodyDatapoints) *DescribeContactResponseBody {
	s.Datapoints = v
	return s
}

func (s *DescribeContactResponseBody) SetMessage(v string) *DescribeContactResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeContactResponseBody) SetRequestId(v string) *DescribeContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContactResponseBody) SetSuccess(v bool) *DescribeContactResponseBody {
	s.Success = &v
	return s
}

type DescribeContactResponseBodyDatapoints struct {
	Channels *DescribeContactResponseBodyDatapointsChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Struct"`
	Name     *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeContactResponseBodyDatapoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactResponseBodyDatapoints) GoString() string {
	return s.String()
}

func (s *DescribeContactResponseBodyDatapoints) SetChannels(v *DescribeContactResponseBodyDatapointsChannels) *DescribeContactResponseBodyDatapoints {
	s.Channels = v
	return s
}

func (s *DescribeContactResponseBodyDatapoints) SetName(v string) *DescribeContactResponseBodyDatapoints {
	s.Name = &v
	return s
}

type DescribeContactResponseBodyDatapointsChannels struct {
	Channel []*DescribeContactResponseBodyDatapointsChannelsChannel `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Repeated"`
}

func (s DescribeContactResponseBodyDatapointsChannels) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactResponseBodyDatapointsChannels) GoString() string {
	return s.String()
}

func (s *DescribeContactResponseBodyDatapointsChannels) SetChannel(v []*DescribeContactResponseBodyDatapointsChannelsChannel) *DescribeContactResponseBodyDatapointsChannels {
	s.Channel = v
	return s
}

type DescribeContactResponseBodyDatapointsChannelsChannel struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeContactResponseBodyDatapointsChannelsChannel) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactResponseBodyDatapointsChannelsChannel) GoString() string {
	return s.String()
}

func (s *DescribeContactResponseBodyDatapointsChannelsChannel) SetType(v string) *DescribeContactResponseBodyDatapointsChannelsChannel {
	s.Type = &v
	return s
}

func (s *DescribeContactResponseBodyDatapointsChannelsChannel) SetValue(v string) *DescribeContactResponseBodyDatapointsChannelsChannel {
	s.Value = &v
	return s
}

type DescribeContactResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContactResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactResponse) GoString() string {
	return s.String()
}

func (s *DescribeContactResponse) SetHeaders(v map[string]*string) *DescribeContactResponse {
	s.Headers = v
	return s
}

func (s *DescribeContactResponse) SetStatusCode(v int32) *DescribeContactResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContactResponse) SetBody(v *DescribeContactResponseBody) *DescribeContactResponse {
	s.Body = v
	return s
}

type DescribeHybridDoubleWriteRequest struct {
	// example:
	//
	// test-source
	SourceNamespace *string `json:"SourceNamespace,omitempty" xml:"SourceNamespace,omitempty"`
	// example:
	//
	// 12706766**********
	SourceUserId    *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	TargetNamespace *string `json:"TargetNamespace,omitempty" xml:"TargetNamespace,omitempty"`
	TargetUserId    *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
}

func (s DescribeHybridDoubleWriteRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridDoubleWriteRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridDoubleWriteRequest) SetSourceNamespace(v string) *DescribeHybridDoubleWriteRequest {
	s.SourceNamespace = &v
	return s
}

func (s *DescribeHybridDoubleWriteRequest) SetSourceUserId(v string) *DescribeHybridDoubleWriteRequest {
	s.SourceUserId = &v
	return s
}

func (s *DescribeHybridDoubleWriteRequest) SetTargetNamespace(v string) *DescribeHybridDoubleWriteRequest {
	s.TargetNamespace = &v
	return s
}

func (s *DescribeHybridDoubleWriteRequest) SetTargetUserId(v string) *DescribeHybridDoubleWriteRequest {
	s.TargetUserId = &v
	return s
}

type DescribeHybridDoubleWriteResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7985D471-3FA8-4EE9-8F4B-45C19DF3D36F
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeHybridDoubleWriteResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeHybridDoubleWriteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridDoubleWriteResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridDoubleWriteResponseBody) SetCode(v string) *DescribeHybridDoubleWriteResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHybridDoubleWriteResponseBody) SetMessage(v string) *DescribeHybridDoubleWriteResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHybridDoubleWriteResponseBody) SetRequestId(v string) *DescribeHybridDoubleWriteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridDoubleWriteResponseBody) SetResult(v []*DescribeHybridDoubleWriteResponseBodyResult) *DescribeHybridDoubleWriteResponseBody {
	s.Result = v
	return s
}

func (s *DescribeHybridDoubleWriteResponseBody) SetSuccess(v string) *DescribeHybridDoubleWriteResponseBody {
	s.Success = &v
	return s
}

type DescribeHybridDoubleWriteResponseBodyResult struct {
	// example:
	//
	// test-target
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// test-source
	SourceNamespace *string `json:"SourceNamespace,omitempty" xml:"SourceNamespace,omitempty"`
	// example:
	//
	// 12706766**********
	SourceUserId *int64 `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	// example:
	//
	// 11234766**********
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeHybridDoubleWriteResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridDoubleWriteResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeHybridDoubleWriteResponseBodyResult) SetNamespace(v string) *DescribeHybridDoubleWriteResponseBodyResult {
	s.Namespace = &v
	return s
}

func (s *DescribeHybridDoubleWriteResponseBodyResult) SetSourceNamespace(v string) *DescribeHybridDoubleWriteResponseBodyResult {
	s.SourceNamespace = &v
	return s
}

func (s *DescribeHybridDoubleWriteResponseBodyResult) SetSourceUserId(v int64) *DescribeHybridDoubleWriteResponseBodyResult {
	s.SourceUserId = &v
	return s
}

func (s *DescribeHybridDoubleWriteResponseBodyResult) SetUserId(v int64) *DescribeHybridDoubleWriteResponseBodyResult {
	s.UserId = &v
	return s
}

type DescribeHybridDoubleWriteResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridDoubleWriteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridDoubleWriteResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridDoubleWriteResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridDoubleWriteResponse) SetHeaders(v map[string]*string) *DescribeHybridDoubleWriteResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridDoubleWriteResponse) SetStatusCode(v int32) *DescribeHybridDoubleWriteResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridDoubleWriteResponse) SetBody(v *DescribeHybridDoubleWriteResponseBody) *DescribeHybridDoubleWriteResponse {
	s.Body = v
	return s
}

type DescribeISPAreaCityRequest struct {
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	Isp  *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeISPAreaCityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeISPAreaCityRequest) GoString() string {
	return s.String()
}

func (s *DescribeISPAreaCityRequest) SetCity(v string) *DescribeISPAreaCityRequest {
	s.City = &v
	return s
}

func (s *DescribeISPAreaCityRequest) SetIsp(v string) *DescribeISPAreaCityRequest {
	s.Isp = &v
	return s
}

type DescribeISPAreaCityResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 84C8BA48-7FD3-46F8-AEE3-E24657C22289
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeISPAreaCityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeISPAreaCityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeISPAreaCityResponseBody) SetCode(v string) *DescribeISPAreaCityResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeISPAreaCityResponseBody) SetData(v string) *DescribeISPAreaCityResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeISPAreaCityResponseBody) SetMessage(v string) *DescribeISPAreaCityResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeISPAreaCityResponseBody) SetRequestId(v string) *DescribeISPAreaCityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeISPAreaCityResponseBody) SetSuccess(v string) *DescribeISPAreaCityResponseBody {
	s.Success = &v
	return s
}

type DescribeISPAreaCityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeISPAreaCityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeISPAreaCityResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeISPAreaCityResponse) GoString() string {
	return s.String()
}

func (s *DescribeISPAreaCityResponse) SetHeaders(v map[string]*string) *DescribeISPAreaCityResponse {
	s.Headers = v
	return s
}

func (s *DescribeISPAreaCityResponse) SetStatusCode(v int32) *DescribeISPAreaCityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeISPAreaCityResponse) SetBody(v *DescribeISPAreaCityResponseBody) *DescribeISPAreaCityResponse {
	s.Body = v
	return s
}

type DescribeMetricRuleListRequest struct {
	AlertState  *string `json:"AlertState,omitempty" xml:"AlertState,omitempty"`
	Dimensions  *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EnableState *bool   `json:"EnableState,omitempty" xml:"EnableState,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MetricName  *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Namespace   *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Page        *string `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleIds     *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
	RuleName    *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeMetricRuleListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListRequest) SetAlertState(v string) *DescribeMetricRuleListRequest {
	s.AlertState = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetDimensions(v string) *DescribeMetricRuleListRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetEnableState(v bool) *DescribeMetricRuleListRequest {
	s.EnableState = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetGroupId(v string) *DescribeMetricRuleListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetMetricName(v string) *DescribeMetricRuleListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetNamespace(v string) *DescribeMetricRuleListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetPage(v string) *DescribeMetricRuleListRequest {
	s.Page = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetPageSize(v string) *DescribeMetricRuleListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetRegionId(v string) *DescribeMetricRuleListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetRuleIds(v string) *DescribeMetricRuleListRequest {
	s.RuleIds = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetRuleName(v string) *DescribeMetricRuleListRequest {
	s.RuleName = &v
	return s
}

type DescribeMetricRuleListResponseBody struct {
	Alarms    *DescribeMetricRuleListResponseBodyAlarms `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Struct"`
	Code      *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *string                                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeMetricRuleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBody) SetAlarms(v *DescribeMetricRuleListResponseBodyAlarms) *DescribeMetricRuleListResponseBody {
	s.Alarms = v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetCode(v int32) *DescribeMetricRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetMessage(v string) *DescribeMetricRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetRequestId(v string) *DescribeMetricRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetSuccess(v bool) *DescribeMetricRuleListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetTotal(v string) *DescribeMetricRuleListResponseBody {
	s.Total = &v
	return s
}

type DescribeMetricRuleListResponseBodyAlarms struct {
	Alarm []*DescribeMetricRuleListResponseBodyAlarmsAlarm `json:"Alarm,omitempty" xml:"Alarm,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleListResponseBodyAlarms) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarms) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarms) SetAlarm(v []*DescribeMetricRuleListResponseBodyAlarmsAlarm) *DescribeMetricRuleListResponseBodyAlarms {
	s.Alarm = v
	return s
}

type DescribeMetricRuleListResponseBodyAlarmsAlarm struct {
	AlertState          *string                                                   `json:"AlertState,omitempty" xml:"AlertState,omitempty"`
	ContactGroups       *string                                                   `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	Dimensions          *string                                                   `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EffectiveInterval   *string                                                   `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	EnableState         *bool                                                     `json:"EnableState,omitempty" xml:"EnableState,omitempty"`
	Escalations         *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	GroupId             *string                                                   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName           *string                                                   `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	MailSubject         *string                                                   `json:"MailSubject,omitempty" xml:"MailSubject,omitempty"`
	MetricName          *string                                                   `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Namespace           *string                                                   `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NoEffectiveInterval *string                                                   `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	Period              *string                                                   `json:"Period,omitempty" xml:"Period,omitempty"`
	Resources           *string                                                   `json:"Resources,omitempty" xml:"Resources,omitempty"`
	RuleId              *string                                                   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName            *string                                                   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	SilenceTime         *string                                                   `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	SourceType          *string                                                   `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Webhook             *string                                                   `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarm) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarm) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetAlertState(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.AlertState = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetContactGroups(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.ContactGroups = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetDimensions(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetEffectiveInterval(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.EffectiveInterval = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetEnableState(v bool) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.EnableState = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetEscalations(v *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Escalations = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetGroupId(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.GroupId = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetGroupName(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.GroupName = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetMailSubject(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.MailSubject = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetMetricName(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetNamespace(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetNoEffectiveInterval(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.NoEffectiveInterval = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetPeriod(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Period = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetResources(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Resources = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetRuleId(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.RuleId = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetRuleName(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.RuleName = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetSilenceTime(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.SilenceTime = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetSourceType(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.SourceType = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetWebhook(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Webhook = &v
	return s
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations struct {
	Critical *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) SetCritical(v *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations {
	s.Critical = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) SetInfo(v *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations {
	s.Info = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) SetWarn(v *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations {
	s.Warn = v
	return s
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *string `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetComparisonOperator(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetStatistics(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetThreshold(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetTimes(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.Times = &v
	return s
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *string `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetComparisonOperator(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetStatistics(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetThreshold(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetTimes(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.Times = &v
	return s
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *string `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetComparisonOperator(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetStatistics(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetThreshold(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetTimes(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.Times = &v
	return s
}

type DescribeMetricRuleListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetricRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetricRuleListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponse) SetHeaders(v map[string]*string) *DescribeMetricRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricRuleListResponse) SetStatusCode(v int32) *DescribeMetricRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetricRuleListResponse) SetBody(v *DescribeMetricRuleListResponseBody) *DescribeMetricRuleListResponse {
	s.Body = v
	return s
}

type DescribeTaskDetailRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 339e0d96-2505-425f-a5c6-22e2c12f8fee
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTaskDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskDetailRequest) SetRegionId(v string) *DescribeTaskDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTaskDetailRequest) SetTaskId(v string) *DescribeTaskDetailRequest {
	s.TaskId = &v
	return s
}

type DescribeTaskDetailResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// a4e6f550-7eac-4a13-b11f-6742aa2d42d1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeTaskDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskDetailResponseBody) SetCode(v string) *DescribeTaskDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTaskDetailResponseBody) SetData(v string) *DescribeTaskDetailResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeTaskDetailResponseBody) SetMessage(v string) *DescribeTaskDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTaskDetailResponseBody) SetRequestId(v string) *DescribeTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskDetailResponseBody) SetSuccess(v string) *DescribeTaskDetailResponseBody {
	s.Success = &v
	return s
}

type DescribeTaskDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTaskDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskDetailResponse) SetHeaders(v map[string]*string) *DescribeTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskDetailResponse) SetStatusCode(v int32) *DescribeTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTaskDetailResponse) SetBody(v *DescribeTaskDetailResponseBody) *DescribeTaskDetailResponse {
	s.Body = v
	return s
}

type DescribeTasksRequest struct {
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 18b0c27f-bab3-441d-a747-9cdcaa8bbac8
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 1
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeTasksRequest) SetKeyword(v string) *DescribeTasksRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeTasksRequest) SetPage(v int32) *DescribeTasksRequest {
	s.Page = &v
	return s
}

func (s *DescribeTasksRequest) SetPageSize(v int32) *DescribeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTasksRequest) SetRegionId(v string) *DescribeTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskId(v string) *DescribeTasksRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskType(v string) *DescribeTasksRequest {
	s.TaskType = &v
	return s
}

type DescribeTasksResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 6A46B8E4-D39E-4DB5-B422-231410654E8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 14
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBody) SetCode(v string) *DescribeTasksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTasksResponseBody) SetData(v string) *DescribeTasksResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeTasksResponseBody) SetMessage(v string) *DescribeTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTasksResponseBody) SetPageNumber(v int32) *DescribeTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksResponseBody) SetPageSize(v int32) *DescribeTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTasksResponseBody) SetRequestId(v string) *DescribeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTasksResponseBody) SetSuccess(v string) *DescribeTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTasksResponseBody) SetTotalCount(v int32) *DescribeTasksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTasksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponse) SetHeaders(v map[string]*string) *DescribeTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeTasksResponse) SetStatusCode(v int32) *DescribeTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTasksResponse) SetBody(v *DescribeTasksResponseBody) *DescribeTasksResponse {
	s.Body = v
	return s
}

type DisableAlarmRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 576fbae7-2fd1-411a-ae13-6f09f4fafdde
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableAlarmRequest) GoString() string {
	return s.String()
}

func (s *DisableAlarmRequest) SetId(v string) *DisableAlarmRequest {
	s.Id = &v
	return s
}

func (s *DisableAlarmRequest) SetRegionId(v string) *DisableAlarmRequest {
	s.RegionId = &v
	return s
}

type DisableAlarmResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DEF01F10-E747-42FE-9152-85CB43B1B552
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAlarmResponseBody) SetCode(v string) *DisableAlarmResponseBody {
	s.Code = &v
	return s
}

func (s *DisableAlarmResponseBody) SetMessage(v string) *DisableAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *DisableAlarmResponseBody) SetRequestId(v string) *DisableAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAlarmResponseBody) SetSuccess(v bool) *DisableAlarmResponseBody {
	s.Success = &v
	return s
}

type DisableAlarmResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableAlarmResponse) GoString() string {
	return s.String()
}

func (s *DisableAlarmResponse) SetHeaders(v map[string]*string) *DisableAlarmResponse {
	s.Headers = v
	return s
}

func (s *DisableAlarmResponse) SetStatusCode(v int32) *DisableAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAlarmResponse) SetBody(v *DisableAlarmResponseBody) *DisableAlarmResponse {
	s.Body = v
	return s
}

type EnableAlarmRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 576fbae7-2fd1-411a-ae13-6f09f4fafdde
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableAlarmRequest) GoString() string {
	return s.String()
}

func (s *EnableAlarmRequest) SetId(v string) *EnableAlarmRequest {
	s.Id = &v
	return s
}

func (s *EnableAlarmRequest) SetRegionId(v string) *EnableAlarmRequest {
	s.RegionId = &v
	return s
}

type EnableAlarmResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1C5E0E5D-76D5-469C-9FA8-D74799B24860
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *EnableAlarmResponseBody) SetCode(v string) *EnableAlarmResponseBody {
	s.Code = &v
	return s
}

func (s *EnableAlarmResponseBody) SetMessage(v string) *EnableAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *EnableAlarmResponseBody) SetRequestId(v string) *EnableAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableAlarmResponseBody) SetSuccess(v bool) *EnableAlarmResponseBody {
	s.Success = &v
	return s
}

type EnableAlarmResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableAlarmResponse) GoString() string {
	return s.String()
}

func (s *EnableAlarmResponse) SetHeaders(v map[string]*string) *EnableAlarmResponse {
	s.Headers = v
	return s
}

func (s *EnableAlarmResponse) SetStatusCode(v int32) *EnableAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableAlarmResponse) SetBody(v *EnableAlarmResponseBody) *EnableAlarmResponse {
	s.Body = v
	return s
}

type GetContactsRequest struct {
	// This parameter is required.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetContactsRequest) GoString() string {
	return s.String()
}

func (s *GetContactsRequest) SetGroupName(v string) *GetContactsRequest {
	s.GroupName = &v
	return s
}

func (s *GetContactsRequest) SetRegionId(v string) *GetContactsRequest {
	s.RegionId = &v
	return s
}

type GetContactsResponseBody struct {
	Code           *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Datapoints     *GetContactsResponseBodyDatapoints `json:"Datapoints,omitempty" xml:"Datapoints,omitempty" type:"Struct"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetContactsResponseBody) GoString() string {
	return s.String()
}

func (s *GetContactsResponseBody) SetCode(v int32) *GetContactsResponseBody {
	s.Code = &v
	return s
}

func (s *GetContactsResponseBody) SetDatapoints(v *GetContactsResponseBodyDatapoints) *GetContactsResponseBody {
	s.Datapoints = v
	return s
}

func (s *GetContactsResponseBody) SetHttpStatusCode(v int32) *GetContactsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetContactsResponseBody) SetMessage(v string) *GetContactsResponseBody {
	s.Message = &v
	return s
}

func (s *GetContactsResponseBody) SetRequestId(v string) *GetContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContactsResponseBody) SetSuccess(v bool) *GetContactsResponseBody {
	s.Success = &v
	return s
}

type GetContactsResponseBodyDatapoints struct {
	Contacts *GetContactsResponseBodyDatapointsContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Struct"`
	Name     *string                                    `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetContactsResponseBodyDatapoints) String() string {
	return tea.Prettify(s)
}

func (s GetContactsResponseBodyDatapoints) GoString() string {
	return s.String()
}

func (s *GetContactsResponseBodyDatapoints) SetContacts(v *GetContactsResponseBodyDatapointsContacts) *GetContactsResponseBodyDatapoints {
	s.Contacts = v
	return s
}

func (s *GetContactsResponseBodyDatapoints) SetName(v string) *GetContactsResponseBodyDatapoints {
	s.Name = &v
	return s
}

type GetContactsResponseBodyDatapointsContacts struct {
	Contact []*string `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s GetContactsResponseBodyDatapointsContacts) String() string {
	return tea.Prettify(s)
}

func (s GetContactsResponseBodyDatapointsContacts) GoString() string {
	return s.String()
}

func (s *GetContactsResponseBodyDatapointsContacts) SetContact(v []*string) *GetContactsResponseBodyDatapointsContacts {
	s.Contact = v
	return s
}

type GetContactsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetContactsResponse) GoString() string {
	return s.String()
}

func (s *GetContactsResponse) SetHeaders(v map[string]*string) *GetContactsResponse {
	s.Headers = v
	return s
}

func (s *GetContactsResponse) SetStatusCode(v int32) *GetContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContactsResponse) SetBody(v *GetContactsResponseBody) *GetContactsResponse {
	s.Body = v
	return s
}

type GetLineSplitResultRequest struct {
	Line          *string `json:"Line,omitempty" xml:"Line,omitempty"`
	Prefix        *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Regex         *string `json:"Regex,omitempty" xml:"Regex,omitempty"`
	SelectContent *string `json:"SelectContent,omitempty" xml:"SelectContent,omitempty"`
	// example:
	//
	// simple|regex
	SplitType *string `json:"SplitType,omitempty" xml:"SplitType,omitempty"`
}

func (s GetLineSplitResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLineSplitResultRequest) GoString() string {
	return s.String()
}

func (s *GetLineSplitResultRequest) SetLine(v string) *GetLineSplitResultRequest {
	s.Line = &v
	return s
}

func (s *GetLineSplitResultRequest) SetPrefix(v string) *GetLineSplitResultRequest {
	s.Prefix = &v
	return s
}

func (s *GetLineSplitResultRequest) SetRegex(v string) *GetLineSplitResultRequest {
	s.Regex = &v
	return s
}

func (s *GetLineSplitResultRequest) SetSelectContent(v string) *GetLineSplitResultRequest {
	s.SelectContent = &v
	return s
}

func (s *GetLineSplitResultRequest) SetSplitType(v string) *GetLineSplitResultRequest {
	s.SplitType = &v
	return s
}

type GetLineSplitResultResponseBody struct {
	Code         *int64                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode    *int64                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resource     *GetLineSplitResultResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLineSplitResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLineSplitResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetLineSplitResultResponseBody) SetCode(v int64) *GetLineSplitResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetLineSplitResultResponseBody) SetErrorCode(v int64) *GetLineSplitResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetLineSplitResultResponseBody) SetErrorMessage(v string) *GetLineSplitResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetLineSplitResultResponseBody) SetRequestId(v string) *GetLineSplitResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLineSplitResultResponseBody) SetResource(v *GetLineSplitResultResponseBodyResource) *GetLineSplitResultResponseBody {
	s.Resource = v
	return s
}

func (s *GetLineSplitResultResponseBody) SetSuccess(v bool) *GetLineSplitResultResponseBody {
	s.Success = &v
	return s
}

type GetLineSplitResultResponseBodyResource struct {
	AdditionalRegex       *string   `json:"AdditionalRegex,omitempty" xml:"AdditionalRegex,omitempty"`
	EndSplitSymbol        *string   `json:"EndSplitSymbol,omitempty" xml:"EndSplitSymbol,omitempty"`
	Regex                 *string   `json:"Regex,omitempty" xml:"Regex,omitempty"`
	RegexSplitResult      []*string `json:"RegexSplitResult,omitempty" xml:"RegexSplitResult,omitempty" type:"Repeated"`
	StartSplitSymbol      *string   `json:"StartSplitSymbol,omitempty" xml:"StartSplitSymbol,omitempty"`
	StartSplitSymbolIndex *int64    `json:"StartSplitSymbolIndex,omitempty" xml:"StartSplitSymbolIndex,omitempty"`
}

func (s GetLineSplitResultResponseBodyResource) String() string {
	return tea.Prettify(s)
}

func (s GetLineSplitResultResponseBodyResource) GoString() string {
	return s.String()
}

func (s *GetLineSplitResultResponseBodyResource) SetAdditionalRegex(v string) *GetLineSplitResultResponseBodyResource {
	s.AdditionalRegex = &v
	return s
}

func (s *GetLineSplitResultResponseBodyResource) SetEndSplitSymbol(v string) *GetLineSplitResultResponseBodyResource {
	s.EndSplitSymbol = &v
	return s
}

func (s *GetLineSplitResultResponseBodyResource) SetRegex(v string) *GetLineSplitResultResponseBodyResource {
	s.Regex = &v
	return s
}

func (s *GetLineSplitResultResponseBodyResource) SetRegexSplitResult(v []*string) *GetLineSplitResultResponseBodyResource {
	s.RegexSplitResult = v
	return s
}

func (s *GetLineSplitResultResponseBodyResource) SetStartSplitSymbol(v string) *GetLineSplitResultResponseBodyResource {
	s.StartSplitSymbol = &v
	return s
}

func (s *GetLineSplitResultResponseBodyResource) SetStartSplitSymbolIndex(v int64) *GetLineSplitResultResponseBodyResource {
	s.StartSplitSymbolIndex = &v
	return s
}

type GetLineSplitResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLineSplitResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLineSplitResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLineSplitResultResponse) GoString() string {
	return s.String()
}

func (s *GetLineSplitResultResponse) SetHeaders(v map[string]*string) *GetLineSplitResultResponse {
	s.Headers = v
	return s
}

func (s *GetLineSplitResultResponse) SetStatusCode(v int32) *GetLineSplitResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLineSplitResultResponse) SetBody(v *GetLineSplitResultResponseBody) *GetLineSplitResultResponse {
	s.Body = v
	return s
}

type GetLogColumnTranslateResultRequest struct {
	ColumnValue     *string `json:"ColumnValue,omitempty" xml:"ColumnValue,omitempty"`
	TranslateConfig *string `json:"TranslateConfig,omitempty" xml:"TranslateConfig,omitempty"`
}

func (s GetLogColumnTranslateResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogColumnTranslateResultRequest) GoString() string {
	return s.String()
}

func (s *GetLogColumnTranslateResultRequest) SetColumnValue(v string) *GetLogColumnTranslateResultRequest {
	s.ColumnValue = &v
	return s
}

func (s *GetLogColumnTranslateResultRequest) SetTranslateConfig(v string) *GetLogColumnTranslateResultRequest {
	s.TranslateConfig = &v
	return s
}

type GetLogColumnTranslateResultResponseBody struct {
	Code         *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode    *int64  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resource     *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLogColumnTranslateResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLogColumnTranslateResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogColumnTranslateResultResponseBody) SetCode(v int64) *GetLogColumnTranslateResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetLogColumnTranslateResultResponseBody) SetErrorCode(v int64) *GetLogColumnTranslateResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetLogColumnTranslateResultResponseBody) SetErrorMessage(v string) *GetLogColumnTranslateResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetLogColumnTranslateResultResponseBody) SetMessage(v string) *GetLogColumnTranslateResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetLogColumnTranslateResultResponseBody) SetRequestId(v string) *GetLogColumnTranslateResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogColumnTranslateResultResponseBody) SetResource(v string) *GetLogColumnTranslateResultResponseBody {
	s.Resource = &v
	return s
}

func (s *GetLogColumnTranslateResultResponseBody) SetSuccess(v bool) *GetLogColumnTranslateResultResponseBody {
	s.Success = &v
	return s
}

type GetLogColumnTranslateResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogColumnTranslateResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogColumnTranslateResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogColumnTranslateResultResponse) GoString() string {
	return s.String()
}

func (s *GetLogColumnTranslateResultResponse) SetHeaders(v map[string]*string) *GetLogColumnTranslateResultResponse {
	s.Headers = v
	return s
}

func (s *GetLogColumnTranslateResultResponse) SetStatusCode(v int32) *GetLogColumnTranslateResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogColumnTranslateResultResponse) SetBody(v *GetLogColumnTranslateResultResponseBody) *GetLogColumnTranslateResultResponse {
	s.Body = v
	return s
}

type GetMonitoringTemplateRequest struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetMonitoringTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMonitoringTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetMonitoringTemplateRequest) SetId(v string) *GetMonitoringTemplateRequest {
	s.Id = &v
	return s
}

func (s *GetMonitoringTemplateRequest) SetName(v string) *GetMonitoringTemplateRequest {
	s.Name = &v
	return s
}

func (s *GetMonitoringTemplateRequest) SetRegionId(v string) *GetMonitoringTemplateRequest {
	s.RegionId = &v
	return s
}

type GetMonitoringTemplateResponseBody struct {
	ErrorCode    *int32                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resource     *GetMonitoringTemplateResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMonitoringTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMonitoringTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetMonitoringTemplateResponseBody) SetErrorCode(v int32) *GetMonitoringTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMonitoringTemplateResponseBody) SetErrorMessage(v string) *GetMonitoringTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMonitoringTemplateResponseBody) SetRequestId(v string) *GetMonitoringTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMonitoringTemplateResponseBody) SetResource(v *GetMonitoringTemplateResponseBodyResource) *GetMonitoringTemplateResponseBody {
	s.Resource = v
	return s
}

func (s *GetMonitoringTemplateResponseBody) SetSuccess(v bool) *GetMonitoringTemplateResponseBody {
	s.Success = &v
	return s
}

type GetMonitoringTemplateResponseBodyResource struct {
	AlertTemplatesJson       *string `json:"AlertTemplatesJson,omitempty" xml:"AlertTemplatesJson,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HostAvailabilityTemplate *string `json:"HostAvailabilityTemplate,omitempty" xml:"HostAvailabilityTemplate,omitempty"`
	Id                       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace                *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ProcessMonitorTemplates  *string `json:"ProcessMonitorTemplates,omitempty" xml:"ProcessMonitorTemplates,omitempty"`
	RestVersion              *string `json:"RestVersion,omitempty" xml:"RestVersion,omitempty"`
	SystemEventTemplates     *string `json:"SystemEventTemplates,omitempty" xml:"SystemEventTemplates,omitempty"`
}

func (s GetMonitoringTemplateResponseBodyResource) String() string {
	return tea.Prettify(s)
}

func (s GetMonitoringTemplateResponseBodyResource) GoString() string {
	return s.String()
}

func (s *GetMonitoringTemplateResponseBodyResource) SetAlertTemplatesJson(v string) *GetMonitoringTemplateResponseBodyResource {
	s.AlertTemplatesJson = &v
	return s
}

func (s *GetMonitoringTemplateResponseBodyResource) SetDescription(v string) *GetMonitoringTemplateResponseBodyResource {
	s.Description = &v
	return s
}

func (s *GetMonitoringTemplateResponseBodyResource) SetHostAvailabilityTemplate(v string) *GetMonitoringTemplateResponseBodyResource {
	s.HostAvailabilityTemplate = &v
	return s
}

func (s *GetMonitoringTemplateResponseBodyResource) SetId(v string) *GetMonitoringTemplateResponseBodyResource {
	s.Id = &v
	return s
}

func (s *GetMonitoringTemplateResponseBodyResource) SetName(v string) *GetMonitoringTemplateResponseBodyResource {
	s.Name = &v
	return s
}

func (s *GetMonitoringTemplateResponseBodyResource) SetNamespace(v string) *GetMonitoringTemplateResponseBodyResource {
	s.Namespace = &v
	return s
}

func (s *GetMonitoringTemplateResponseBodyResource) SetProcessMonitorTemplates(v string) *GetMonitoringTemplateResponseBodyResource {
	s.ProcessMonitorTemplates = &v
	return s
}

func (s *GetMonitoringTemplateResponseBodyResource) SetRestVersion(v string) *GetMonitoringTemplateResponseBodyResource {
	s.RestVersion = &v
	return s
}

func (s *GetMonitoringTemplateResponseBodyResource) SetSystemEventTemplates(v string) *GetMonitoringTemplateResponseBodyResource {
	s.SystemEventTemplates = &v
	return s
}

type GetMonitoringTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMonitoringTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMonitoringTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMonitoringTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetMonitoringTemplateResponse) SetHeaders(v map[string]*string) *GetMonitoringTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetMonitoringTemplateResponse) SetStatusCode(v int32) *GetMonitoringTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMonitoringTemplateResponse) SetBody(v *GetMonitoringTemplateResponseBody) *GetMonitoringTemplateResponse {
	s.Body = v
	return s
}

type GetMyGroupsRequest struct {
	BindUrl             *string `json:"BindUrl,omitempty" xml:"BindUrl,omitempty"`
	GroupId             *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName           *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SelectContactGroups *bool   `json:"SelectContactGroups,omitempty" xml:"SelectContactGroups,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetMyGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMyGroupsRequest) GoString() string {
	return s.String()
}

func (s *GetMyGroupsRequest) SetBindUrl(v string) *GetMyGroupsRequest {
	s.BindUrl = &v
	return s
}

func (s *GetMyGroupsRequest) SetGroupId(v int64) *GetMyGroupsRequest {
	s.GroupId = &v
	return s
}

func (s *GetMyGroupsRequest) SetGroupName(v string) *GetMyGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *GetMyGroupsRequest) SetInstanceId(v string) *GetMyGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetMyGroupsRequest) SetRegionId(v string) *GetMyGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *GetMyGroupsRequest) SetSelectContactGroups(v bool) *GetMyGroupsRequest {
	s.SelectContactGroups = &v
	return s
}

func (s *GetMyGroupsRequest) SetType(v string) *GetMyGroupsRequest {
	s.Type = &v
	return s
}

type GetMyGroupsResponseBody struct {
	ErrorCode    *int32                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Group        *GetMyGroupsResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	RequestId    *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMyGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMyGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMyGroupsResponseBody) SetErrorCode(v int32) *GetMyGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMyGroupsResponseBody) SetErrorMessage(v string) *GetMyGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMyGroupsResponseBody) SetGroup(v *GetMyGroupsResponseBodyGroup) *GetMyGroupsResponseBody {
	s.Group = v
	return s
}

func (s *GetMyGroupsResponseBody) SetRequestId(v string) *GetMyGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMyGroupsResponseBody) SetSuccess(v bool) *GetMyGroupsResponseBody {
	s.Success = &v
	return s
}

type GetMyGroupsResponseBodyGroup struct {
	BindUrl       *string                                    `json:"BindUrl,omitempty" xml:"BindUrl,omitempty"`
	ContactGroups *GetMyGroupsResponseBodyGroupContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
	GroupId       *int64                                     `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string                                    `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ServiceId     *int64                                     `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	Type          *string                                    `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetMyGroupsResponseBodyGroup) String() string {
	return tea.Prettify(s)
}

func (s GetMyGroupsResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *GetMyGroupsResponseBodyGroup) SetBindUrl(v string) *GetMyGroupsResponseBodyGroup {
	s.BindUrl = &v
	return s
}

func (s *GetMyGroupsResponseBodyGroup) SetContactGroups(v *GetMyGroupsResponseBodyGroupContactGroups) *GetMyGroupsResponseBodyGroup {
	s.ContactGroups = v
	return s
}

func (s *GetMyGroupsResponseBodyGroup) SetGroupId(v int64) *GetMyGroupsResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *GetMyGroupsResponseBodyGroup) SetGroupName(v string) *GetMyGroupsResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *GetMyGroupsResponseBodyGroup) SetServiceId(v int64) *GetMyGroupsResponseBodyGroup {
	s.ServiceId = &v
	return s
}

func (s *GetMyGroupsResponseBodyGroup) SetType(v string) *GetMyGroupsResponseBodyGroup {
	s.Type = &v
	return s
}

type GetMyGroupsResponseBodyGroupContactGroups struct {
	ContactGroup []*GetMyGroupsResponseBodyGroupContactGroupsContactGroup `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s GetMyGroupsResponseBodyGroupContactGroups) String() string {
	return tea.Prettify(s)
}

func (s GetMyGroupsResponseBodyGroupContactGroups) GoString() string {
	return s.String()
}

func (s *GetMyGroupsResponseBodyGroupContactGroups) SetContactGroup(v []*GetMyGroupsResponseBodyGroupContactGroupsContactGroup) *GetMyGroupsResponseBodyGroupContactGroups {
	s.ContactGroup = v
	return s
}

type GetMyGroupsResponseBodyGroupContactGroupsContactGroup struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetMyGroupsResponseBodyGroupContactGroupsContactGroup) String() string {
	return tea.Prettify(s)
}

func (s GetMyGroupsResponseBodyGroupContactGroupsContactGroup) GoString() string {
	return s.String()
}

func (s *GetMyGroupsResponseBodyGroupContactGroupsContactGroup) SetName(v string) *GetMyGroupsResponseBodyGroupContactGroupsContactGroup {
	s.Name = &v
	return s
}

type GetMyGroupsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMyGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMyGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMyGroupsResponse) GoString() string {
	return s.String()
}

func (s *GetMyGroupsResponse) SetHeaders(v map[string]*string) *GetMyGroupsResponse {
	s.Headers = v
	return s
}

func (s *GetMyGroupsResponse) SetStatusCode(v int32) *GetMyGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMyGroupsResponse) SetBody(v *GetMyGroupsResponseBody) *GetMyGroupsResponse {
	s.Body = v
	return s
}

type ListAlarmRequest struct {
	// example:
	//
	// {"instanceId":"XXX"}
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// example:
	//
	// i-2ze3w55tr2rcpejpcfap_72071739-396b-497d-849c-59a73de44bcf
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	IsEnable *bool   `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 2
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ok
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmRequest) SetDimension(v string) *ListAlarmRequest {
	s.Dimension = &v
	return s
}

func (s *ListAlarmRequest) SetId(v string) *ListAlarmRequest {
	s.Id = &v
	return s
}

func (s *ListAlarmRequest) SetIsEnable(v bool) *ListAlarmRequest {
	s.IsEnable = &v
	return s
}

func (s *ListAlarmRequest) SetName(v string) *ListAlarmRequest {
	s.Name = &v
	return s
}

func (s *ListAlarmRequest) SetNamespace(v string) *ListAlarmRequest {
	s.Namespace = &v
	return s
}

func (s *ListAlarmRequest) SetPageNumber(v int32) *ListAlarmRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlarmRequest) SetPageSize(v int32) *ListAlarmRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlarmRequest) SetRegionId(v string) *ListAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *ListAlarmRequest) SetState(v string) *ListAlarmRequest {
	s.State = &v
	return s
}

type ListAlarmResponseBody struct {
	AlarmList *ListAlarmResponseBodyAlarmList `json:"AlarmList,omitempty" xml:"AlarmList,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// EFD27F56-5799-4CE8-B625-56DF3332331C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 27
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlarmResponseBody) SetAlarmList(v *ListAlarmResponseBodyAlarmList) *ListAlarmResponseBody {
	s.AlarmList = v
	return s
}

func (s *ListAlarmResponseBody) SetCode(v string) *ListAlarmResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlarmResponseBody) SetMessage(v string) *ListAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlarmResponseBody) SetNextToken(v int32) *ListAlarmResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAlarmResponseBody) SetRequestId(v string) *ListAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlarmResponseBody) SetSuccess(v bool) *ListAlarmResponseBody {
	s.Success = &v
	return s
}

func (s *ListAlarmResponseBody) SetTotal(v int32) *ListAlarmResponseBody {
	s.Total = &v
	return s
}

type ListAlarmResponseBodyAlarmList struct {
	Alarm []*ListAlarmResponseBodyAlarmListAlarm `json:"Alarm,omitempty" xml:"Alarm,omitempty" type:"Repeated"`
}

func (s ListAlarmResponseBodyAlarmList) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmResponseBodyAlarmList) GoString() string {
	return s.String()
}

func (s *ListAlarmResponseBodyAlarmList) SetAlarm(v []*ListAlarmResponseBodyAlarmListAlarm) *ListAlarmResponseBodyAlarmList {
	s.Alarm = v
	return s
}

type ListAlarmResponseBodyAlarmListAlarm struct {
	// example:
	//
	// >
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// example:
	//
	// ["test4nudou"]
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// example:
	//
	// ["{\\"instanceId\\":\\" i-abcdefgh123456\\"}"]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 24
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// example:
	//
	// i-2ze3w55tr2rcpejpcfap_72071739-396b-497d-849c-59a73de44bcf
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// CPUUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 0
	NotifyType *int32 `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// example:
	//
	// 300
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// example:
	//
	// 0
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// ok
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// example:
	//
	// null
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s ListAlarmResponseBodyAlarmListAlarm) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmResponseBodyAlarmListAlarm) GoString() string {
	return s.String()
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetComparisonOperator(v string) *ListAlarmResponseBodyAlarmListAlarm {
	s.ComparisonOperator = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetContactGroups(v string) *ListAlarmResponseBodyAlarmListAlarm {
	s.ContactGroups = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetDimensions(v string) *ListAlarmResponseBodyAlarmListAlarm {
	s.Dimensions = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetEnable(v bool) *ListAlarmResponseBodyAlarmListAlarm {
	s.Enable = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetEndTime(v int32) *ListAlarmResponseBodyAlarmListAlarm {
	s.EndTime = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetEvaluationCount(v int32) *ListAlarmResponseBodyAlarmListAlarm {
	s.EvaluationCount = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetId(v string) *ListAlarmResponseBodyAlarmListAlarm {
	s.Id = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetMetricName(v string) *ListAlarmResponseBodyAlarmListAlarm {
	s.MetricName = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetName(v string) *ListAlarmResponseBodyAlarmListAlarm {
	s.Name = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetNamespace(v string) *ListAlarmResponseBodyAlarmListAlarm {
	s.Namespace = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetNotifyType(v int32) *ListAlarmResponseBodyAlarmListAlarm {
	s.NotifyType = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetPeriod(v int32) *ListAlarmResponseBodyAlarmListAlarm {
	s.Period = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetSilenceTime(v int32) *ListAlarmResponseBodyAlarmListAlarm {
	s.SilenceTime = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetStartTime(v int32) *ListAlarmResponseBodyAlarmListAlarm {
	s.StartTime = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetState(v string) *ListAlarmResponseBodyAlarmListAlarm {
	s.State = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetStatistics(v string) *ListAlarmResponseBodyAlarmListAlarm {
	s.Statistics = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetThreshold(v string) *ListAlarmResponseBodyAlarmListAlarm {
	s.Threshold = &v
	return s
}

func (s *ListAlarmResponseBodyAlarmListAlarm) SetWebhook(v string) *ListAlarmResponseBodyAlarmListAlarm {
	s.Webhook = &v
	return s
}

type ListAlarmResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmResponse) GoString() string {
	return s.String()
}

func (s *ListAlarmResponse) SetHeaders(v map[string]*string) *ListAlarmResponse {
	s.Headers = v
	return s
}

func (s *ListAlarmResponse) SetStatusCode(v int32) *ListAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlarmResponse) SetBody(v *ListAlarmResponseBody) *ListAlarmResponse {
	s.Body = v
	return s
}

type ListAlarmHistoryRequest struct {
	// example:
	//
	// 1
	Cursor *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	// example:
	//
	// 1548927491223
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1a775e37-dfba-430c-ab9f-7036475c8bfb_2dbe619b-0483-402e-9437-7c7a38fba7ed
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 3
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 1548913091223
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAlarmHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoryRequest) SetCursor(v string) *ListAlarmHistoryRequest {
	s.Cursor = &v
	return s
}

func (s *ListAlarmHistoryRequest) SetEndTime(v string) *ListAlarmHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *ListAlarmHistoryRequest) SetId(v string) *ListAlarmHistoryRequest {
	s.Id = &v
	return s
}

func (s *ListAlarmHistoryRequest) SetRegionId(v string) *ListAlarmHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *ListAlarmHistoryRequest) SetSize(v int32) *ListAlarmHistoryRequest {
	s.Size = &v
	return s
}

func (s *ListAlarmHistoryRequest) SetStartTime(v string) *ListAlarmHistoryRequest {
	s.StartTime = &v
	return s
}

type ListAlarmHistoryResponseBody struct {
	AlarmHistoryList *ListAlarmHistoryResponseBodyAlarmHistoryList `json:"AlarmHistoryList,omitempty" xml:"AlarmHistoryList,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Cursor *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1DBBCE29-0F69-435C-B65C-53D1011D1D72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAlarmHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoryResponseBody) SetAlarmHistoryList(v *ListAlarmHistoryResponseBodyAlarmHistoryList) *ListAlarmHistoryResponseBody {
	s.AlarmHistoryList = v
	return s
}

func (s *ListAlarmHistoryResponseBody) SetCode(v string) *ListAlarmHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlarmHistoryResponseBody) SetCursor(v string) *ListAlarmHistoryResponseBody {
	s.Cursor = &v
	return s
}

func (s *ListAlarmHistoryResponseBody) SetMessage(v string) *ListAlarmHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlarmHistoryResponseBody) SetRequestId(v string) *ListAlarmHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlarmHistoryResponseBody) SetSuccess(v bool) *ListAlarmHistoryResponseBody {
	s.Success = &v
	return s
}

type ListAlarmHistoryResponseBodyAlarmHistoryList struct {
	AlarmHistory []*ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory `json:"AlarmHistory,omitempty" xml:"AlarmHistory,omitempty" type:"Repeated"`
}

func (s ListAlarmHistoryResponseBodyAlarmHistoryList) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoryResponseBodyAlarmHistoryList) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoryResponseBodyAlarmHistoryList) SetAlarmHistory(v []*ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) *ListAlarmHistoryResponseBodyAlarmHistoryList {
	s.AlarmHistory = v
	return s
}

type ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory struct {
	// example:
	//
	// 1548926982000
	AlarmTime     *int64  `json:"AlarmTime,omitempty" xml:"AlarmTime,omitempty"`
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// example:
	//
	// {"instanceId":"XXX"}
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// example:
	//
	// 1a775e37-dfba-430c-ab9f-7036475c8bfb_2dbe619b-0483-402e-9437-7c7a38fba7ed
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test-demo
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 2851651669
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// example:
	//
	// CPUUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// ALARM
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 84401454
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetAlarmTime(v int64) *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.AlarmTime = &v
	return s
}

func (s *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetContactGroups(v string) *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactGroups = &v
	return s
}

func (s *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetDimension(v string) *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.Dimension = &v
	return s
}

func (s *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetEvaluationCount(v int32) *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.EvaluationCount = &v
	return s
}

func (s *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetId(v string) *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.Id = &v
	return s
}

func (s *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetInstanceName(v string) *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.InstanceName = &v
	return s
}

func (s *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetLastTime(v int64) *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.LastTime = &v
	return s
}

func (s *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetMetricName(v string) *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.MetricName = &v
	return s
}

func (s *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetName(v string) *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.Name = &v
	return s
}

func (s *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetNamespace(v string) *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.Namespace = &v
	return s
}

func (s *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetState(v string) *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.State = &v
	return s
}

func (s *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetStatus(v int32) *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.Status = &v
	return s
}

func (s *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory) SetValue(v string) *ListAlarmHistoryResponseBodyAlarmHistoryListAlarmHistory {
	s.Value = &v
	return s
}

type ListAlarmHistoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlarmHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlarmHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoryResponse) SetHeaders(v map[string]*string) *ListAlarmHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListAlarmHistoryResponse) SetStatusCode(v int32) *ListAlarmHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlarmHistoryResponse) SetBody(v *ListAlarmHistoryResponseBody) *ListAlarmHistoryResponse {
	s.Body = v
	return s
}

type ListContactGroupRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListContactGroupRequest) GoString() string {
	return s.String()
}

func (s *ListContactGroupRequest) SetPageNumber(v int32) *ListContactGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListContactGroupRequest) SetPageSize(v int32) *ListContactGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListContactGroupRequest) SetRegionId(v string) *ListContactGroupRequest {
	s.RegionId = &v
	return s
}

type ListContactGroupResponseBody struct {
	// example:
	//
	// 200
	Code          *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	ContactGroups *ListContactGroupResponseBodyContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// D3D03B0A-CF1A-4DAB-BF0D-D4B6ACD5677A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListContactGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListContactGroupResponseBody) SetCode(v string) *ListContactGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListContactGroupResponseBody) SetContactGroups(v *ListContactGroupResponseBodyContactGroups) *ListContactGroupResponseBody {
	s.ContactGroups = v
	return s
}

func (s *ListContactGroupResponseBody) SetMessage(v string) *ListContactGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListContactGroupResponseBody) SetNextToken(v int32) *ListContactGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListContactGroupResponseBody) SetRequestId(v string) *ListContactGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContactGroupResponseBody) SetSuccess(v bool) *ListContactGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ListContactGroupResponseBody) SetTotal(v int32) *ListContactGroupResponseBody {
	s.Total = &v
	return s
}

type ListContactGroupResponseBodyContactGroups struct {
	ContactGroup []*string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s ListContactGroupResponseBodyContactGroups) String() string {
	return tea.Prettify(s)
}

func (s ListContactGroupResponseBodyContactGroups) GoString() string {
	return s.String()
}

func (s *ListContactGroupResponseBodyContactGroups) SetContactGroup(v []*string) *ListContactGroupResponseBodyContactGroups {
	s.ContactGroup = v
	return s
}

type ListContactGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListContactGroupResponse) GoString() string {
	return s.String()
}

func (s *ListContactGroupResponse) SetHeaders(v map[string]*string) *ListContactGroupResponse {
	s.Headers = v
	return s
}

func (s *ListContactGroupResponse) SetStatusCode(v int32) *ListContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContactGroupResponse) SetBody(v *ListContactGroupResponseBody) *ListContactGroupResponse {
	s.Body = v
	return s
}

type ListEventRulesRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	NamePrefix *string `json:"NamePrefix,omitempty" xml:"NamePrefix,omitempty"`
	Page       *string `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListEventRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEventRulesRequest) GoString() string {
	return s.String()
}

func (s *ListEventRulesRequest) SetGroupId(v string) *ListEventRulesRequest {
	s.GroupId = &v
	return s
}

func (s *ListEventRulesRequest) SetNamePrefix(v string) *ListEventRulesRequest {
	s.NamePrefix = &v
	return s
}

func (s *ListEventRulesRequest) SetPage(v string) *ListEventRulesRequest {
	s.Page = &v
	return s
}

func (s *ListEventRulesRequest) SetPageSize(v string) *ListEventRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListEventRulesRequest) SetRegionId(v string) *ListEventRulesRequest {
	s.RegionId = &v
	return s
}

type ListEventRulesResponseBody struct {
	Code        *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	CurrentPage *int32                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Datapoints  *ListEventRulesResponseBodyDatapoints `json:"Datapoints,omitempty" xml:"Datapoints,omitempty" type:"Struct"`
	Message     *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken   *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Total       *int32                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEventRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventRulesResponseBody) SetCode(v string) *ListEventRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListEventRulesResponseBody) SetCurrentPage(v int32) *ListEventRulesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListEventRulesResponseBody) SetDatapoints(v *ListEventRulesResponseBodyDatapoints) *ListEventRulesResponseBody {
	s.Datapoints = v
	return s
}

func (s *ListEventRulesResponseBody) SetMessage(v string) *ListEventRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEventRulesResponseBody) SetNextToken(v string) *ListEventRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEventRulesResponseBody) SetRequestId(v string) *ListEventRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventRulesResponseBody) SetSuccess(v bool) *ListEventRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListEventRulesResponseBody) SetTotal(v int32) *ListEventRulesResponseBody {
	s.Total = &v
	return s
}

type ListEventRulesResponseBodyDatapoints struct {
	EventRule []*ListEventRulesResponseBodyDatapointsEventRule `json:"EventRule,omitempty" xml:"EventRule,omitempty" type:"Repeated"`
}

func (s ListEventRulesResponseBodyDatapoints) String() string {
	return tea.Prettify(s)
}

func (s ListEventRulesResponseBodyDatapoints) GoString() string {
	return s.String()
}

func (s *ListEventRulesResponseBodyDatapoints) SetEventRule(v []*ListEventRulesResponseBodyDatapointsEventRule) *ListEventRulesResponseBodyDatapoints {
	s.EventRule = v
	return s
}

type ListEventRulesResponseBodyDatapointsEventRule struct {
	Description  *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	EventPattern *ListEventRulesResponseBodyDatapointsEventRuleEventPattern `json:"EventPattern,omitempty" xml:"EventPattern,omitempty" type:"Struct"`
	EventType    *string                                                    `json:"EventType,omitempty" xml:"EventType,omitempty"`
	GroupId      *string                                                    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Name         *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	State        *string                                                    `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListEventRulesResponseBodyDatapointsEventRule) String() string {
	return tea.Prettify(s)
}

func (s ListEventRulesResponseBodyDatapointsEventRule) GoString() string {
	return s.String()
}

func (s *ListEventRulesResponseBodyDatapointsEventRule) SetDescription(v string) *ListEventRulesResponseBodyDatapointsEventRule {
	s.Description = &v
	return s
}

func (s *ListEventRulesResponseBodyDatapointsEventRule) SetEventPattern(v *ListEventRulesResponseBodyDatapointsEventRuleEventPattern) *ListEventRulesResponseBodyDatapointsEventRule {
	s.EventPattern = v
	return s
}

func (s *ListEventRulesResponseBodyDatapointsEventRule) SetEventType(v string) *ListEventRulesResponseBodyDatapointsEventRule {
	s.EventType = &v
	return s
}

func (s *ListEventRulesResponseBodyDatapointsEventRule) SetGroupId(v string) *ListEventRulesResponseBodyDatapointsEventRule {
	s.GroupId = &v
	return s
}

func (s *ListEventRulesResponseBodyDatapointsEventRule) SetName(v string) *ListEventRulesResponseBodyDatapointsEventRule {
	s.Name = &v
	return s
}

func (s *ListEventRulesResponseBodyDatapointsEventRule) SetState(v string) *ListEventRulesResponseBodyDatapointsEventRule {
	s.State = &v
	return s
}

type ListEventRulesResponseBodyDatapointsEventRuleEventPattern struct {
	EventPattern []*ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPattern `json:"EventPattern,omitempty" xml:"EventPattern,omitempty" type:"Repeated"`
}

func (s ListEventRulesResponseBodyDatapointsEventRuleEventPattern) String() string {
	return tea.Prettify(s)
}

func (s ListEventRulesResponseBodyDatapointsEventRuleEventPattern) GoString() string {
	return s.String()
}

func (s *ListEventRulesResponseBodyDatapointsEventRuleEventPattern) SetEventPattern(v []*ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPattern) *ListEventRulesResponseBodyDatapointsEventRuleEventPattern {
	s.EventPattern = v
	return s
}

type ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPattern struct {
	EventTypeList *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternEventTypeList `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty" type:"Struct"`
	LevelList     *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternLevelList     `json:"LevelList,omitempty" xml:"LevelList,omitempty" type:"Struct"`
	NameList      *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternNameList      `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Struct"`
	Product       *string                                                                             `json:"Product,omitempty" xml:"Product,omitempty"`
	StatusList    *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternStatusList    `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Struct"`
}

func (s ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPattern) String() string {
	return tea.Prettify(s)
}

func (s ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPattern) GoString() string {
	return s.String()
}

func (s *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPattern) SetEventTypeList(v *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternEventTypeList) *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPattern {
	s.EventTypeList = v
	return s
}

func (s *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPattern) SetLevelList(v *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternLevelList) *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPattern {
	s.LevelList = v
	return s
}

func (s *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPattern) SetNameList(v *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternNameList) *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPattern {
	s.NameList = v
	return s
}

func (s *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPattern) SetProduct(v string) *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPattern {
	s.Product = &v
	return s
}

func (s *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPattern) SetStatusList(v *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternStatusList) *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPattern {
	s.StatusList = v
	return s
}

type ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternEventTypeList struct {
	EventTypeList []*string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty" type:"Repeated"`
}

func (s ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternEventTypeList) String() string {
	return tea.Prettify(s)
}

func (s ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternEventTypeList) GoString() string {
	return s.String()
}

func (s *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternEventTypeList) SetEventTypeList(v []*string) *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternEventTypeList {
	s.EventTypeList = v
	return s
}

type ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternLevelList struct {
	LevelList []*string `json:"LevelList,omitempty" xml:"LevelList,omitempty" type:"Repeated"`
}

func (s ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternLevelList) String() string {
	return tea.Prettify(s)
}

func (s ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternLevelList) GoString() string {
	return s.String()
}

func (s *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternLevelList) SetLevelList(v []*string) *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternLevelList {
	s.LevelList = v
	return s
}

type ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternNameList struct {
	NameList []*string `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Repeated"`
}

func (s ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternNameList) String() string {
	return tea.Prettify(s)
}

func (s ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternNameList) GoString() string {
	return s.String()
}

func (s *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternNameList) SetNameList(v []*string) *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternNameList {
	s.NameList = v
	return s
}

type ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternStatusList struct {
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternStatusList) String() string {
	return tea.Prettify(s)
}

func (s ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternStatusList) GoString() string {
	return s.String()
}

func (s *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternStatusList) SetStatusList(v []*string) *ListEventRulesResponseBodyDatapointsEventRuleEventPatternEventPatternStatusList {
	s.StatusList = v
	return s
}

type ListEventRulesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEventRulesResponse) GoString() string {
	return s.String()
}

func (s *ListEventRulesResponse) SetHeaders(v map[string]*string) *ListEventRulesResponse {
	s.Headers = v
	return s
}

func (s *ListEventRulesResponse) SetStatusCode(v int32) *ListEventRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventRulesResponse) SetBody(v *ListEventRulesResponseBody) *ListEventRulesResponse {
	s.Body = v
	return s
}

type ListMyGroupInstancesRequest struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// This parameter is required.
	GroupId     *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	Keyword     *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Total       *bool   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMyGroupInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMyGroupInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListMyGroupInstancesRequest) SetCategory(v string) *ListMyGroupInstancesRequest {
	s.Category = &v
	return s
}

func (s *ListMyGroupInstancesRequest) SetGroupId(v int64) *ListMyGroupInstancesRequest {
	s.GroupId = &v
	return s
}

func (s *ListMyGroupInstancesRequest) SetInstanceIds(v string) *ListMyGroupInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *ListMyGroupInstancesRequest) SetKeyword(v string) *ListMyGroupInstancesRequest {
	s.Keyword = &v
	return s
}

func (s *ListMyGroupInstancesRequest) SetPageNumber(v int32) *ListMyGroupInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMyGroupInstancesRequest) SetPageSize(v int32) *ListMyGroupInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMyGroupInstancesRequest) SetRegionId(v string) *ListMyGroupInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListMyGroupInstancesRequest) SetTotal(v bool) *ListMyGroupInstancesRequest {
	s.Total = &v
	return s
}

type ListMyGroupInstancesResponseBody struct {
	ErrorCode    *int32                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	PageNumber   *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources    *ListMyGroupInstancesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Total        *int32                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMyGroupInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMyGroupInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMyGroupInstancesResponseBody) SetErrorCode(v int32) *ListMyGroupInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMyGroupInstancesResponseBody) SetErrorMessage(v string) *ListMyGroupInstancesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMyGroupInstancesResponseBody) SetPageNumber(v int32) *ListMyGroupInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListMyGroupInstancesResponseBody) SetPageSize(v int32) *ListMyGroupInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMyGroupInstancesResponseBody) SetRequestId(v string) *ListMyGroupInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMyGroupInstancesResponseBody) SetResources(v *ListMyGroupInstancesResponseBodyResources) *ListMyGroupInstancesResponseBody {
	s.Resources = v
	return s
}

func (s *ListMyGroupInstancesResponseBody) SetSuccess(v bool) *ListMyGroupInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListMyGroupInstancesResponseBody) SetTotal(v int32) *ListMyGroupInstancesResponseBody {
	s.Total = &v
	return s
}

type ListMyGroupInstancesResponseBodyResources struct {
	Resource []*ListMyGroupInstancesResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s ListMyGroupInstancesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListMyGroupInstancesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListMyGroupInstancesResponseBodyResources) SetResource(v []*ListMyGroupInstancesResponseBodyResourcesResource) *ListMyGroupInstancesResponseBodyResources {
	s.Resource = v
	return s
}

type ListMyGroupInstancesResponseBodyResourcesResource struct {
	AliUid       *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Category     *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListMyGroupInstancesResponseBodyResourcesResource) String() string {
	return tea.Prettify(s)
}

func (s ListMyGroupInstancesResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *ListMyGroupInstancesResponseBodyResourcesResource) SetAliUid(v int64) *ListMyGroupInstancesResponseBodyResourcesResource {
	s.AliUid = &v
	return s
}

func (s *ListMyGroupInstancesResponseBodyResourcesResource) SetCategory(v string) *ListMyGroupInstancesResponseBodyResourcesResource {
	s.Category = &v
	return s
}

func (s *ListMyGroupInstancesResponseBodyResourcesResource) SetId(v int64) *ListMyGroupInstancesResponseBodyResourcesResource {
	s.Id = &v
	return s
}

func (s *ListMyGroupInstancesResponseBodyResourcesResource) SetInstanceId(v string) *ListMyGroupInstancesResponseBodyResourcesResource {
	s.InstanceId = &v
	return s
}

func (s *ListMyGroupInstancesResponseBodyResourcesResource) SetInstanceName(v string) *ListMyGroupInstancesResponseBodyResourcesResource {
	s.InstanceName = &v
	return s
}

func (s *ListMyGroupInstancesResponseBodyResourcesResource) SetRegionId(v string) *ListMyGroupInstancesResponseBodyResourcesResource {
	s.RegionId = &v
	return s
}

type ListMyGroupInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMyGroupInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMyGroupInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMyGroupInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListMyGroupInstancesResponse) SetHeaders(v map[string]*string) *ListMyGroupInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListMyGroupInstancesResponse) SetStatusCode(v int32) *ListMyGroupInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMyGroupInstancesResponse) SetBody(v *ListMyGroupInstancesResponseBody) *ListMyGroupInstancesResponse {
	s.Body = v
	return s
}

type ListMyGroupsRequest struct {
	BindUrls            *string `json:"BindUrls,omitempty" xml:"BindUrls,omitempty"`
	GroupName           *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Keyword             *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNumber          *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SelectContactGroups *bool   `json:"SelectContactGroups,omitempty" xml:"SelectContactGroups,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListMyGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMyGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListMyGroupsRequest) SetBindUrls(v string) *ListMyGroupsRequest {
	s.BindUrls = &v
	return s
}

func (s *ListMyGroupsRequest) SetGroupName(v string) *ListMyGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *ListMyGroupsRequest) SetInstanceId(v string) *ListMyGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListMyGroupsRequest) SetKeyword(v string) *ListMyGroupsRequest {
	s.Keyword = &v
	return s
}

func (s *ListMyGroupsRequest) SetPageNumber(v int32) *ListMyGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMyGroupsRequest) SetPageSize(v int32) *ListMyGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMyGroupsRequest) SetRegionId(v string) *ListMyGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListMyGroupsRequest) SetSelectContactGroups(v bool) *ListMyGroupsRequest {
	s.SelectContactGroups = &v
	return s
}

func (s *ListMyGroupsRequest) SetType(v string) *ListMyGroupsRequest {
	s.Type = &v
	return s
}

type ListMyGroupsResponseBody struct {
	ErrorCode    *int32                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	PageNumber   *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources    *ListMyGroupsResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	Success      *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Total        *int32                             `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMyGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMyGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMyGroupsResponseBody) SetErrorCode(v int32) *ListMyGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMyGroupsResponseBody) SetErrorMessage(v string) *ListMyGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMyGroupsResponseBody) SetPageNumber(v int32) *ListMyGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListMyGroupsResponseBody) SetPageSize(v int32) *ListMyGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMyGroupsResponseBody) SetRequestId(v string) *ListMyGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMyGroupsResponseBody) SetResources(v *ListMyGroupsResponseBodyResources) *ListMyGroupsResponseBody {
	s.Resources = v
	return s
}

func (s *ListMyGroupsResponseBody) SetSuccess(v bool) *ListMyGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMyGroupsResponseBody) SetTotal(v int32) *ListMyGroupsResponseBody {
	s.Total = &v
	return s
}

type ListMyGroupsResponseBodyResources struct {
	Resource []*ListMyGroupsResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s ListMyGroupsResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListMyGroupsResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListMyGroupsResponseBodyResources) SetResource(v []*ListMyGroupsResponseBodyResourcesResource) *ListMyGroupsResponseBodyResources {
	s.Resource = v
	return s
}

type ListMyGroupsResponseBodyResourcesResource struct {
	BindUrl       *string                                                 `json:"BindUrl,omitempty" xml:"BindUrl,omitempty"`
	BindUrls      *string                                                 `json:"BindUrls,omitempty" xml:"BindUrls,omitempty"`
	ContactGroups *ListMyGroupsResponseBodyResourcesResourceContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
	GmtCreate     *int64                                                  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified   *int64                                                  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GroupId       *int64                                                  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string                                                 `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ServiceId     *string                                                 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	Type          *string                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListMyGroupsResponseBodyResourcesResource) String() string {
	return tea.Prettify(s)
}

func (s ListMyGroupsResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *ListMyGroupsResponseBodyResourcesResource) SetBindUrl(v string) *ListMyGroupsResponseBodyResourcesResource {
	s.BindUrl = &v
	return s
}

func (s *ListMyGroupsResponseBodyResourcesResource) SetBindUrls(v string) *ListMyGroupsResponseBodyResourcesResource {
	s.BindUrls = &v
	return s
}

func (s *ListMyGroupsResponseBodyResourcesResource) SetContactGroups(v *ListMyGroupsResponseBodyResourcesResourceContactGroups) *ListMyGroupsResponseBodyResourcesResource {
	s.ContactGroups = v
	return s
}

func (s *ListMyGroupsResponseBodyResourcesResource) SetGmtCreate(v int64) *ListMyGroupsResponseBodyResourcesResource {
	s.GmtCreate = &v
	return s
}

func (s *ListMyGroupsResponseBodyResourcesResource) SetGmtModified(v int64) *ListMyGroupsResponseBodyResourcesResource {
	s.GmtModified = &v
	return s
}

func (s *ListMyGroupsResponseBodyResourcesResource) SetGroupId(v int64) *ListMyGroupsResponseBodyResourcesResource {
	s.GroupId = &v
	return s
}

func (s *ListMyGroupsResponseBodyResourcesResource) SetGroupName(v string) *ListMyGroupsResponseBodyResourcesResource {
	s.GroupName = &v
	return s
}

func (s *ListMyGroupsResponseBodyResourcesResource) SetServiceId(v string) *ListMyGroupsResponseBodyResourcesResource {
	s.ServiceId = &v
	return s
}

func (s *ListMyGroupsResponseBodyResourcesResource) SetType(v string) *ListMyGroupsResponseBodyResourcesResource {
	s.Type = &v
	return s
}

type ListMyGroupsResponseBodyResourcesResourceContactGroups struct {
	ContactGroup []*ListMyGroupsResponseBodyResourcesResourceContactGroupsContactGroup `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s ListMyGroupsResponseBodyResourcesResourceContactGroups) String() string {
	return tea.Prettify(s)
}

func (s ListMyGroupsResponseBodyResourcesResourceContactGroups) GoString() string {
	return s.String()
}

func (s *ListMyGroupsResponseBodyResourcesResourceContactGroups) SetContactGroup(v []*ListMyGroupsResponseBodyResourcesResourceContactGroupsContactGroup) *ListMyGroupsResponseBodyResourcesResourceContactGroups {
	s.ContactGroup = v
	return s
}

type ListMyGroupsResponseBodyResourcesResourceContactGroupsContactGroup struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListMyGroupsResponseBodyResourcesResourceContactGroupsContactGroup) String() string {
	return tea.Prettify(s)
}

func (s ListMyGroupsResponseBodyResourcesResourceContactGroupsContactGroup) GoString() string {
	return s.String()
}

func (s *ListMyGroupsResponseBodyResourcesResourceContactGroupsContactGroup) SetName(v string) *ListMyGroupsResponseBodyResourcesResourceContactGroupsContactGroup {
	s.Name = &v
	return s
}

type ListMyGroupsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMyGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMyGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMyGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListMyGroupsResponse) SetHeaders(v map[string]*string) *ListMyGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListMyGroupsResponse) SetStatusCode(v int32) *ListMyGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMyGroupsResponse) SetBody(v *ListMyGroupsResponseBody) *ListMyGroupsResponse {
	s.Body = v
	return s
}

type ModifyTaskRequest struct {
	// example:
	//
	// http://www.aliyun.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// 5e9b-4c12-b39e-7f277ac44b11
	AlertIds *string `json:"AlertIds,omitempty" xml:"AlertIds,omitempty"`
	// example:
	//
	// [{"alarmActions":"xxx","metricName":"Availability","expression":"$Availability<96"}]
	AlertRule *string `json:"AlertRule,omitempty" xml:"AlertRule,omitempty"`
	// example:
	//
	// 1
	Interval     *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IntervalUnit *string `json:"IntervalUnit,omitempty" xml:"IntervalUnit,omitempty"`
	// example:
	//
	// [{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]
	IspCity  *string `json:"IspCity,omitempty" xml:"IspCity,omitempty"`
	Options  *string `json:"Options,omitempty" xml:"Options,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8f880e3d-d924-47ab-84d2-fa1a72e24211
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// aliyunTest
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// cms
	Caller *string `json:"caller,omitempty" xml:"caller,omitempty"`
}

func (s ModifyTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyTaskRequest) SetAddress(v string) *ModifyTaskRequest {
	s.Address = &v
	return s
}

func (s *ModifyTaskRequest) SetAlertIds(v string) *ModifyTaskRequest {
	s.AlertIds = &v
	return s
}

func (s *ModifyTaskRequest) SetAlertRule(v string) *ModifyTaskRequest {
	s.AlertRule = &v
	return s
}

func (s *ModifyTaskRequest) SetInterval(v string) *ModifyTaskRequest {
	s.Interval = &v
	return s
}

func (s *ModifyTaskRequest) SetIntervalUnit(v string) *ModifyTaskRequest {
	s.IntervalUnit = &v
	return s
}

func (s *ModifyTaskRequest) SetIspCity(v string) *ModifyTaskRequest {
	s.IspCity = &v
	return s
}

func (s *ModifyTaskRequest) SetOptions(v string) *ModifyTaskRequest {
	s.Options = &v
	return s
}

func (s *ModifyTaskRequest) SetRegionId(v string) *ModifyTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyTaskRequest) SetTaskId(v string) *ModifyTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ModifyTaskRequest) SetTaskName(v string) *ModifyTaskRequest {
	s.TaskName = &v
	return s
}

func (s *ModifyTaskRequest) SetCaller(v string) *ModifyTaskRequest {
	s.Caller = &v
	return s
}

type ModifyTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"count":1}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successfull
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// a4e6f550-7eac-4a13-b11f-6742aa2d42d1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTaskResponseBody) SetCode(v string) *ModifyTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyTaskResponseBody) SetData(v string) *ModifyTaskResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyTaskResponseBody) SetMessage(v string) *ModifyTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyTaskResponseBody) SetRequestId(v string) *ModifyTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTaskResponseBody) SetSuccess(v string) *ModifyTaskResponseBody {
	s.Success = &v
	return s
}

type ModifyTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyTaskResponse) SetHeaders(v map[string]*string) *ModifyTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyTaskResponse) SetStatusCode(v int32) *ModifyTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTaskResponse) SetBody(v *ModifyTaskResponseBody) *ModifyTaskResponse {
	s.Body = v
	return s
}

type NodeListRequest struct {
	HostName         *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceIds      *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	InstanceRegionId *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	KeyWord          *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SerialNumbers    *string `json:"SerialNumbers,omitempty" xml:"SerialNumbers,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s NodeListRequest) String() string {
	return tea.Prettify(s)
}

func (s NodeListRequest) GoString() string {
	return s.String()
}

func (s *NodeListRequest) SetHostName(v string) *NodeListRequest {
	s.HostName = &v
	return s
}

func (s *NodeListRequest) SetInstanceIds(v string) *NodeListRequest {
	s.InstanceIds = &v
	return s
}

func (s *NodeListRequest) SetInstanceRegionId(v string) *NodeListRequest {
	s.InstanceRegionId = &v
	return s
}

func (s *NodeListRequest) SetKeyWord(v string) *NodeListRequest {
	s.KeyWord = &v
	return s
}

func (s *NodeListRequest) SetPageNumber(v int32) *NodeListRequest {
	s.PageNumber = &v
	return s
}

func (s *NodeListRequest) SetPageSize(v int32) *NodeListRequest {
	s.PageSize = &v
	return s
}

func (s *NodeListRequest) SetRegionId(v string) *NodeListRequest {
	s.RegionId = &v
	return s
}

func (s *NodeListRequest) SetSerialNumbers(v string) *NodeListRequest {
	s.SerialNumbers = &v
	return s
}

func (s *NodeListRequest) SetStatus(v string) *NodeListRequest {
	s.Status = &v
	return s
}

func (s *NodeListRequest) SetUserId(v int64) *NodeListRequest {
	s.UserId = &v
	return s
}

type NodeListResponseBody struct {
	ErrorCode    *int32                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Nodes        *NodeListResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	PageNumber   *int32                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageTotal    *int32                     `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	RequestId    *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Total        *int32                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s NodeListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NodeListResponseBody) GoString() string {
	return s.String()
}

func (s *NodeListResponseBody) SetErrorCode(v int32) *NodeListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *NodeListResponseBody) SetErrorMessage(v string) *NodeListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *NodeListResponseBody) SetNodes(v *NodeListResponseBodyNodes) *NodeListResponseBody {
	s.Nodes = v
	return s
}

func (s *NodeListResponseBody) SetPageNumber(v int32) *NodeListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *NodeListResponseBody) SetPageSize(v int32) *NodeListResponseBody {
	s.PageSize = &v
	return s
}

func (s *NodeListResponseBody) SetPageTotal(v int32) *NodeListResponseBody {
	s.PageTotal = &v
	return s
}

func (s *NodeListResponseBody) SetRequestId(v string) *NodeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *NodeListResponseBody) SetSuccess(v bool) *NodeListResponseBody {
	s.Success = &v
	return s
}

func (s *NodeListResponseBody) SetTotal(v int32) *NodeListResponseBody {
	s.Total = &v
	return s
}

type NodeListResponseBodyNodes struct {
	Node []*NodeListResponseBodyNodesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Repeated"`
}

func (s NodeListResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s NodeListResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *NodeListResponseBodyNodes) SetNode(v []*NodeListResponseBodyNodesNode) *NodeListResponseBodyNodes {
	s.Node = v
	return s
}

type NodeListResponseBodyNodesNode struct {
	AliUid             *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AliyunHost         *bool   `json:"AliyunHost,omitempty" xml:"AliyunHost,omitempty"`
	EipAddress         *string `json:"EipAddress,omitempty" xml:"EipAddress,omitempty"`
	EipId              *string `json:"EipId,omitempty" xml:"EipId,omitempty"`
	HostName           *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	IpGroup            *string `json:"IpGroup,omitempty" xml:"IpGroup,omitempty"`
	NatIp              *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	NetworkType        *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OperatingSystem    *string `json:"OperatingSystem,omitempty" xml:"OperatingSystem,omitempty"`
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SerialNumber       *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	TianjimonVersion   *string `json:"TianjimonVersion,omitempty" xml:"TianjimonVersion,omitempty"`
}

func (s NodeListResponseBodyNodesNode) String() string {
	return tea.Prettify(s)
}

func (s NodeListResponseBodyNodesNode) GoString() string {
	return s.String()
}

func (s *NodeListResponseBodyNodesNode) SetAliUid(v int64) *NodeListResponseBodyNodesNode {
	s.AliUid = &v
	return s
}

func (s *NodeListResponseBodyNodesNode) SetAliyunHost(v bool) *NodeListResponseBodyNodesNode {
	s.AliyunHost = &v
	return s
}

func (s *NodeListResponseBodyNodesNode) SetEipAddress(v string) *NodeListResponseBodyNodesNode {
	s.EipAddress = &v
	return s
}

func (s *NodeListResponseBodyNodesNode) SetEipId(v string) *NodeListResponseBodyNodesNode {
	s.EipId = &v
	return s
}

func (s *NodeListResponseBodyNodesNode) SetHostName(v string) *NodeListResponseBodyNodesNode {
	s.HostName = &v
	return s
}

func (s *NodeListResponseBodyNodesNode) SetInstanceId(v string) *NodeListResponseBodyNodesNode {
	s.InstanceId = &v
	return s
}

func (s *NodeListResponseBodyNodesNode) SetInstanceTypeFamily(v string) *NodeListResponseBodyNodesNode {
	s.InstanceTypeFamily = &v
	return s
}

func (s *NodeListResponseBodyNodesNode) SetIpGroup(v string) *NodeListResponseBodyNodesNode {
	s.IpGroup = &v
	return s
}

func (s *NodeListResponseBodyNodesNode) SetNatIp(v string) *NodeListResponseBodyNodesNode {
	s.NatIp = &v
	return s
}

func (s *NodeListResponseBodyNodesNode) SetNetworkType(v string) *NodeListResponseBodyNodesNode {
	s.NetworkType = &v
	return s
}

func (s *NodeListResponseBodyNodesNode) SetOperatingSystem(v string) *NodeListResponseBodyNodesNode {
	s.OperatingSystem = &v
	return s
}

func (s *NodeListResponseBodyNodesNode) SetRegion(v string) *NodeListResponseBodyNodesNode {
	s.Region = &v
	return s
}

func (s *NodeListResponseBodyNodesNode) SetSerialNumber(v string) *NodeListResponseBodyNodesNode {
	s.SerialNumber = &v
	return s
}

func (s *NodeListResponseBodyNodesNode) SetTianjimonVersion(v string) *NodeListResponseBodyNodesNode {
	s.TianjimonVersion = &v
	return s
}

type NodeListResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NodeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NodeListResponse) String() string {
	return tea.Prettify(s)
}

func (s NodeListResponse) GoString() string {
	return s.String()
}

func (s *NodeListResponse) SetHeaders(v map[string]*string) *NodeListResponse {
	s.Headers = v
	return s
}

func (s *NodeListResponse) SetStatusCode(v int32) *NodeListResponse {
	s.StatusCode = &v
	return s
}

func (s *NodeListResponse) SetBody(v *NodeListResponseBody) *NodeListResponse {
	s.Body = v
	return s
}

type NodeProcessCreateRequest struct {
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// This parameter is required.
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	ProcessUser *string `json:"ProcessUser,omitempty" xml:"ProcessUser,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s NodeProcessCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s NodeProcessCreateRequest) GoString() string {
	return s.String()
}

func (s *NodeProcessCreateRequest) SetCommand(v string) *NodeProcessCreateRequest {
	s.Command = &v
	return s
}

func (s *NodeProcessCreateRequest) SetInstanceId(v string) *NodeProcessCreateRequest {
	s.InstanceId = &v
	return s
}

func (s *NodeProcessCreateRequest) SetName(v string) *NodeProcessCreateRequest {
	s.Name = &v
	return s
}

func (s *NodeProcessCreateRequest) SetProcessName(v string) *NodeProcessCreateRequest {
	s.ProcessName = &v
	return s
}

func (s *NodeProcessCreateRequest) SetProcessUser(v string) *NodeProcessCreateRequest {
	s.ProcessUser = &v
	return s
}

func (s *NodeProcessCreateRequest) SetRegionId(v string) *NodeProcessCreateRequest {
	s.RegionId = &v
	return s
}

type NodeProcessCreateResponseBody struct {
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s NodeProcessCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NodeProcessCreateResponseBody) GoString() string {
	return s.String()
}

func (s *NodeProcessCreateResponseBody) SetErrorCode(v int32) *NodeProcessCreateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *NodeProcessCreateResponseBody) SetErrorMessage(v string) *NodeProcessCreateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *NodeProcessCreateResponseBody) SetId(v int64) *NodeProcessCreateResponseBody {
	s.Id = &v
	return s
}

func (s *NodeProcessCreateResponseBody) SetRequestId(v string) *NodeProcessCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *NodeProcessCreateResponseBody) SetSuccess(v bool) *NodeProcessCreateResponseBody {
	s.Success = &v
	return s
}

type NodeProcessCreateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NodeProcessCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NodeProcessCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s NodeProcessCreateResponse) GoString() string {
	return s.String()
}

func (s *NodeProcessCreateResponse) SetHeaders(v map[string]*string) *NodeProcessCreateResponse {
	s.Headers = v
	return s
}

func (s *NodeProcessCreateResponse) SetStatusCode(v int32) *NodeProcessCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *NodeProcessCreateResponse) SetBody(v *NodeProcessCreateResponseBody) *NodeProcessCreateResponse {
	s.Body = v
	return s
}

type NodeProcessesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s NodeProcessesRequest) String() string {
	return tea.Prettify(s)
}

func (s NodeProcessesRequest) GoString() string {
	return s.String()
}

func (s *NodeProcessesRequest) SetInstanceId(v string) *NodeProcessesRequest {
	s.InstanceId = &v
	return s
}

func (s *NodeProcessesRequest) SetRegionId(v string) *NodeProcessesRequest {
	s.RegionId = &v
	return s
}

type NodeProcessesResponseBody struct {
	ErrorCode     *int32                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage  *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	NodeProcesses *NodeProcessesResponseBodyNodeProcesses `json:"NodeProcesses,omitempty" xml:"NodeProcesses,omitempty" type:"Struct"`
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s NodeProcessesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NodeProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *NodeProcessesResponseBody) SetErrorCode(v int32) *NodeProcessesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *NodeProcessesResponseBody) SetErrorMessage(v string) *NodeProcessesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *NodeProcessesResponseBody) SetNodeProcesses(v *NodeProcessesResponseBodyNodeProcesses) *NodeProcessesResponseBody {
	s.NodeProcesses = v
	return s
}

func (s *NodeProcessesResponseBody) SetRequestId(v string) *NodeProcessesResponseBody {
	s.RequestId = &v
	return s
}

func (s *NodeProcessesResponseBody) SetSuccess(v bool) *NodeProcessesResponseBody {
	s.Success = &v
	return s
}

type NodeProcessesResponseBodyNodeProcesses struct {
	NodeProcess []*NodeProcessesResponseBodyNodeProcessesNodeProcess `json:"NodeProcess,omitempty" xml:"NodeProcess,omitempty" type:"Repeated"`
}

func (s NodeProcessesResponseBodyNodeProcesses) String() string {
	return tea.Prettify(s)
}

func (s NodeProcessesResponseBodyNodeProcesses) GoString() string {
	return s.String()
}

func (s *NodeProcessesResponseBodyNodeProcesses) SetNodeProcess(v []*NodeProcessesResponseBodyNodeProcessesNodeProcess) *NodeProcessesResponseBodyNodeProcesses {
	s.NodeProcess = v
	return s
}

type NodeProcessesResponseBodyNodeProcessesNodeProcess struct {
	Command     *string `json:"Command,omitempty" xml:"Command,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	ProcessUser *string `json:"ProcessUser,omitempty" xml:"ProcessUser,omitempty"`
}

func (s NodeProcessesResponseBodyNodeProcessesNodeProcess) String() string {
	return tea.Prettify(s)
}

func (s NodeProcessesResponseBodyNodeProcessesNodeProcess) GoString() string {
	return s.String()
}

func (s *NodeProcessesResponseBodyNodeProcessesNodeProcess) SetCommand(v string) *NodeProcessesResponseBodyNodeProcessesNodeProcess {
	s.Command = &v
	return s
}

func (s *NodeProcessesResponseBodyNodeProcessesNodeProcess) SetId(v int64) *NodeProcessesResponseBodyNodeProcessesNodeProcess {
	s.Id = &v
	return s
}

func (s *NodeProcessesResponseBodyNodeProcessesNodeProcess) SetInstanceId(v string) *NodeProcessesResponseBodyNodeProcessesNodeProcess {
	s.InstanceId = &v
	return s
}

func (s *NodeProcessesResponseBodyNodeProcessesNodeProcess) SetName(v string) *NodeProcessesResponseBodyNodeProcessesNodeProcess {
	s.Name = &v
	return s
}

func (s *NodeProcessesResponseBodyNodeProcessesNodeProcess) SetProcessName(v string) *NodeProcessesResponseBodyNodeProcessesNodeProcess {
	s.ProcessName = &v
	return s
}

func (s *NodeProcessesResponseBodyNodeProcessesNodeProcess) SetProcessUser(v string) *NodeProcessesResponseBodyNodeProcessesNodeProcess {
	s.ProcessUser = &v
	return s
}

type NodeProcessesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NodeProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NodeProcessesResponse) String() string {
	return tea.Prettify(s)
}

func (s NodeProcessesResponse) GoString() string {
	return s.String()
}

func (s *NodeProcessesResponse) SetHeaders(v map[string]*string) *NodeProcessesResponse {
	s.Headers = v
	return s
}

func (s *NodeProcessesResponse) SetStatusCode(v int32) *NodeProcessesResponse {
	s.StatusCode = &v
	return s
}

func (s *NodeProcessesResponse) SetBody(v *NodeProcessesResponseBody) *NodeProcessesResponse {
	s.Body = v
	return s
}

type NodeStatusListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i-abcdefgh123456,i-abcdefgh123457
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s NodeStatusListRequest) String() string {
	return tea.Prettify(s)
}

func (s NodeStatusListRequest) GoString() string {
	return s.String()
}

func (s *NodeStatusListRequest) SetInstanceIds(v string) *NodeStatusListRequest {
	s.InstanceIds = &v
	return s
}

func (s *NodeStatusListRequest) SetRegionId(v string) *NodeStatusListRequest {
	s.RegionId = &v
	return s
}

type NodeStatusListResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// InstanceIds is mandatory for this action.
	ErrorMessage   *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	NodeStatusList *NodeStatusListResponseBodyNodeStatusList `json:"NodeStatusList,omitempty" xml:"NodeStatusList,omitempty" type:"Struct"`
	// example:
	//
	// 1BB8FE8E-9BBE-490F-82EC-BF70FB849D55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s NodeStatusListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NodeStatusListResponseBody) GoString() string {
	return s.String()
}

func (s *NodeStatusListResponseBody) SetErrorCode(v int32) *NodeStatusListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *NodeStatusListResponseBody) SetErrorMessage(v string) *NodeStatusListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *NodeStatusListResponseBody) SetNodeStatusList(v *NodeStatusListResponseBodyNodeStatusList) *NodeStatusListResponseBody {
	s.NodeStatusList = v
	return s
}

func (s *NodeStatusListResponseBody) SetRequestId(v string) *NodeStatusListResponseBody {
	s.RequestId = &v
	return s
}

func (s *NodeStatusListResponseBody) SetSuccess(v bool) *NodeStatusListResponseBody {
	s.Success = &v
	return s
}

type NodeStatusListResponseBodyNodeStatusList struct {
	NodeStatus []*NodeStatusListResponseBodyNodeStatusListNodeStatus `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty" type:"Repeated"`
}

func (s NodeStatusListResponseBodyNodeStatusList) String() string {
	return tea.Prettify(s)
}

func (s NodeStatusListResponseBodyNodeStatusList) GoString() string {
	return s.String()
}

func (s *NodeStatusListResponseBodyNodeStatusList) SetNodeStatus(v []*NodeStatusListResponseBodyNodeStatusListNodeStatus) *NodeStatusListResponseBodyNodeStatusList {
	s.NodeStatus = v
	return s
}

type NodeStatusListResponseBodyNodeStatusListNodeStatus struct {
	// example:
	//
	// true
	AutoInstall *bool `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	// example:
	//
	// i-abcdefgh123456
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s NodeStatusListResponseBodyNodeStatusListNodeStatus) String() string {
	return tea.Prettify(s)
}

func (s NodeStatusListResponseBodyNodeStatusListNodeStatus) GoString() string {
	return s.String()
}

func (s *NodeStatusListResponseBodyNodeStatusListNodeStatus) SetAutoInstall(v bool) *NodeStatusListResponseBodyNodeStatusListNodeStatus {
	s.AutoInstall = &v
	return s
}

func (s *NodeStatusListResponseBodyNodeStatusListNodeStatus) SetInstanceId(v string) *NodeStatusListResponseBodyNodeStatusListNodeStatus {
	s.InstanceId = &v
	return s
}

func (s *NodeStatusListResponseBodyNodeStatusListNodeStatus) SetStatus(v string) *NodeStatusListResponseBodyNodeStatusListNodeStatus {
	s.Status = &v
	return s
}

type NodeStatusListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NodeStatusListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NodeStatusListResponse) String() string {
	return tea.Prettify(s)
}

func (s NodeStatusListResponse) GoString() string {
	return s.String()
}

func (s *NodeStatusListResponse) SetHeaders(v map[string]*string) *NodeStatusListResponse {
	s.Headers = v
	return s
}

func (s *NodeStatusListResponse) SetStatusCode(v int32) *NodeStatusListResponse {
	s.StatusCode = &v
	return s
}

func (s *NodeStatusListResponse) SetBody(v *NodeStatusListResponseBody) *NodeStatusListResponse {
	s.Body = v
	return s
}

type NodeUninstallRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s NodeUninstallRequest) String() string {
	return tea.Prettify(s)
}

func (s NodeUninstallRequest) GoString() string {
	return s.String()
}

func (s *NodeUninstallRequest) SetInstanceId(v string) *NodeUninstallRequest {
	s.InstanceId = &v
	return s
}

func (s *NodeUninstallRequest) SetRegionId(v string) *NodeUninstallRequest {
	s.RegionId = &v
	return s
}

type NodeUninstallResponseBody struct {
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s NodeUninstallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NodeUninstallResponseBody) GoString() string {
	return s.String()
}

func (s *NodeUninstallResponseBody) SetErrorCode(v int32) *NodeUninstallResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *NodeUninstallResponseBody) SetErrorMessage(v string) *NodeUninstallResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *NodeUninstallResponseBody) SetRequestId(v string) *NodeUninstallResponseBody {
	s.RequestId = &v
	return s
}

func (s *NodeUninstallResponseBody) SetSuccess(v bool) *NodeUninstallResponseBody {
	s.Success = &v
	return s
}

type NodeUninstallResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NodeUninstallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NodeUninstallResponse) String() string {
	return tea.Prettify(s)
}

func (s NodeUninstallResponse) GoString() string {
	return s.String()
}

func (s *NodeUninstallResponse) SetHeaders(v map[string]*string) *NodeUninstallResponse {
	s.Headers = v
	return s
}

func (s *NodeUninstallResponse) SetStatusCode(v int32) *NodeUninstallResponse {
	s.StatusCode = &v
	return s
}

func (s *NodeUninstallResponse) SetBody(v *NodeUninstallResponseBody) *NodeUninstallResponse {
	s.Body = v
	return s
}

type PutCustomMetricRequest struct {
	MetricList *string `json:"MetricList,omitempty" xml:"MetricList,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s PutCustomMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s PutCustomMetricRequest) GoString() string {
	return s.String()
}

func (s *PutCustomMetricRequest) SetMetricList(v string) *PutCustomMetricRequest {
	s.MetricList = &v
	return s
}

func (s *PutCustomMetricRequest) SetRegionId(v string) *PutCustomMetricRequest {
	s.RegionId = &v
	return s
}

type PutCustomMetricResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutCustomMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutCustomMetricResponseBody) GoString() string {
	return s.String()
}

func (s *PutCustomMetricResponseBody) SetCode(v string) *PutCustomMetricResponseBody {
	s.Code = &v
	return s
}

func (s *PutCustomMetricResponseBody) SetData(v string) *PutCustomMetricResponseBody {
	s.Data = &v
	return s
}

func (s *PutCustomMetricResponseBody) SetMessage(v string) *PutCustomMetricResponseBody {
	s.Message = &v
	return s
}

func (s *PutCustomMetricResponseBody) SetRequestId(v string) *PutCustomMetricResponseBody {
	s.RequestId = &v
	return s
}

type PutCustomMetricResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutCustomMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutCustomMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s PutCustomMetricResponse) GoString() string {
	return s.String()
}

func (s *PutCustomMetricResponse) SetHeaders(v map[string]*string) *PutCustomMetricResponse {
	s.Headers = v
	return s
}

func (s *PutCustomMetricResponse) SetStatusCode(v int32) *PutCustomMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *PutCustomMetricResponse) SetBody(v *PutCustomMetricResponseBody) *PutCustomMetricResponse {
	s.Body = v
	return s
}

type PutEventRequest struct {
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s PutEventRequest) String() string {
	return tea.Prettify(s)
}

func (s PutEventRequest) GoString() string {
	return s.String()
}

func (s *PutEventRequest) SetEventInfo(v string) *PutEventRequest {
	s.EventInfo = &v
	return s
}

func (s *PutEventRequest) SetRegionId(v string) *PutEventRequest {
	s.RegionId = &v
	return s
}

type PutEventResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutEventResponseBody) GoString() string {
	return s.String()
}

func (s *PutEventResponseBody) SetCode(v string) *PutEventResponseBody {
	s.Code = &v
	return s
}

func (s *PutEventResponseBody) SetData(v string) *PutEventResponseBody {
	s.Data = &v
	return s
}

func (s *PutEventResponseBody) SetMessage(v string) *PutEventResponseBody {
	s.Message = &v
	return s
}

func (s *PutEventResponseBody) SetRequestId(v string) *PutEventResponseBody {
	s.RequestId = &v
	return s
}

type PutEventResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutEventResponse) String() string {
	return tea.Prettify(s)
}

func (s PutEventResponse) GoString() string {
	return s.String()
}

func (s *PutEventResponse) SetHeaders(v map[string]*string) *PutEventResponse {
	s.Headers = v
	return s
}

func (s *PutEventResponse) SetStatusCode(v int32) *PutEventResponse {
	s.StatusCode = &v
	return s
}

func (s *PutEventResponse) SetBody(v *PutEventResponseBody) *PutEventResponse {
	s.Body = v
	return s
}

type PutEventRuleRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	EventPattern []*PutEventRuleRequestEventPattern `json:"EventPattern,omitempty" xml:"EventPattern,omitempty" type:"Repeated"`
	EventType    *string                            `json:"EventType,omitempty" xml:"EventType,omitempty"`
	GroupId      *string                            `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	State    *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s PutEventRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleRequest) GoString() string {
	return s.String()
}

func (s *PutEventRuleRequest) SetDescription(v string) *PutEventRuleRequest {
	s.Description = &v
	return s
}

func (s *PutEventRuleRequest) SetEventPattern(v []*PutEventRuleRequestEventPattern) *PutEventRuleRequest {
	s.EventPattern = v
	return s
}

func (s *PutEventRuleRequest) SetEventType(v string) *PutEventRuleRequest {
	s.EventType = &v
	return s
}

func (s *PutEventRuleRequest) SetGroupId(v string) *PutEventRuleRequest {
	s.GroupId = &v
	return s
}

func (s *PutEventRuleRequest) SetName(v string) *PutEventRuleRequest {
	s.Name = &v
	return s
}

func (s *PutEventRuleRequest) SetRegionId(v string) *PutEventRuleRequest {
	s.RegionId = &v
	return s
}

func (s *PutEventRuleRequest) SetState(v string) *PutEventRuleRequest {
	s.State = &v
	return s
}

type PutEventRuleRequestEventPattern struct {
	EventTypeList []*string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty" type:"Repeated"`
	LevelList     []*string `json:"LevelList,omitempty" xml:"LevelList,omitempty" type:"Repeated"`
	NameList      []*string `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Repeated"`
	Product       *string   `json:"Product,omitempty" xml:"Product,omitempty"`
	StatusList    []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s PutEventRuleRequestEventPattern) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleRequestEventPattern) GoString() string {
	return s.String()
}

func (s *PutEventRuleRequestEventPattern) SetEventTypeList(v []*string) *PutEventRuleRequestEventPattern {
	s.EventTypeList = v
	return s
}

func (s *PutEventRuleRequestEventPattern) SetLevelList(v []*string) *PutEventRuleRequestEventPattern {
	s.LevelList = v
	return s
}

func (s *PutEventRuleRequestEventPattern) SetNameList(v []*string) *PutEventRuleRequestEventPattern {
	s.NameList = v
	return s
}

func (s *PutEventRuleRequestEventPattern) SetProduct(v string) *PutEventRuleRequestEventPattern {
	s.Product = &v
	return s
}

func (s *PutEventRuleRequestEventPattern) SetStatusList(v []*string) *PutEventRuleRequestEventPattern {
	s.StatusList = v
	return s
}

type PutEventRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutEventRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutEventRuleResponseBody) SetCode(v string) *PutEventRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PutEventRuleResponseBody) SetData(v string) *PutEventRuleResponseBody {
	s.Data = &v
	return s
}

func (s *PutEventRuleResponseBody) SetMessage(v string) *PutEventRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PutEventRuleResponseBody) SetRequestId(v string) *PutEventRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutEventRuleResponseBody) SetSuccess(v bool) *PutEventRuleResponseBody {
	s.Success = &v
	return s
}

type PutEventRuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutEventRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutEventRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleResponse) GoString() string {
	return s.String()
}

func (s *PutEventRuleResponse) SetHeaders(v map[string]*string) *PutEventRuleResponse {
	s.Headers = v
	return s
}

func (s *PutEventRuleResponse) SetStatusCode(v int32) *PutEventRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PutEventRuleResponse) SetBody(v *PutEventRuleResponseBody) *PutEventRuleResponse {
	s.Body = v
	return s
}

type PutEventTargetsRequest struct {
	ContactParameters []*PutEventTargetsRequestContactParameters `json:"ContactParameters,omitempty" xml:"ContactParameters,omitempty" type:"Repeated"`
	FcParameters      []*PutEventTargetsRequestFcParameters      `json:"FcParameters,omitempty" xml:"FcParameters,omitempty" type:"Repeated"`
	MnsParameters     []*PutEventTargetsRequestMnsParameters     `json:"MnsParameters,omitempty" xml:"MnsParameters,omitempty" type:"Repeated"`
	RegionId          *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	RuleName          *string                                    `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	SlsParameters     []*PutEventTargetsRequestSlsParameters     `json:"SlsParameters,omitempty" xml:"SlsParameters,omitempty" type:"Repeated"`
	WebhookParameters []*PutEventTargetsRequestWebhookParameters `json:"WebhookParameters,omitempty" xml:"WebhookParameters,omitempty" type:"Repeated"`
}

func (s PutEventTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsRequest) GoString() string {
	return s.String()
}

func (s *PutEventTargetsRequest) SetContactParameters(v []*PutEventTargetsRequestContactParameters) *PutEventTargetsRequest {
	s.ContactParameters = v
	return s
}

func (s *PutEventTargetsRequest) SetFcParameters(v []*PutEventTargetsRequestFcParameters) *PutEventTargetsRequest {
	s.FcParameters = v
	return s
}

func (s *PutEventTargetsRequest) SetMnsParameters(v []*PutEventTargetsRequestMnsParameters) *PutEventTargetsRequest {
	s.MnsParameters = v
	return s
}

func (s *PutEventTargetsRequest) SetRegionId(v string) *PutEventTargetsRequest {
	s.RegionId = &v
	return s
}

func (s *PutEventTargetsRequest) SetRuleName(v string) *PutEventTargetsRequest {
	s.RuleName = &v
	return s
}

func (s *PutEventTargetsRequest) SetSlsParameters(v []*PutEventTargetsRequestSlsParameters) *PutEventTargetsRequest {
	s.SlsParameters = v
	return s
}

func (s *PutEventTargetsRequest) SetWebhookParameters(v []*PutEventTargetsRequestWebhookParameters) *PutEventTargetsRequest {
	s.WebhookParameters = v
	return s
}

type PutEventTargetsRequestContactParameters struct {
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Level            *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s PutEventTargetsRequestContactParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsRequestContactParameters) GoString() string {
	return s.String()
}

func (s *PutEventTargetsRequestContactParameters) SetContactGroupName(v string) *PutEventTargetsRequestContactParameters {
	s.ContactGroupName = &v
	return s
}

func (s *PutEventTargetsRequestContactParameters) SetId(v string) *PutEventTargetsRequestContactParameters {
	s.Id = &v
	return s
}

func (s *PutEventTargetsRequestContactParameters) SetLevel(v string) *PutEventTargetsRequestContactParameters {
	s.Level = &v
	return s
}

type PutEventTargetsRequestFcParameters struct {
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ServiceName  *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s PutEventTargetsRequestFcParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsRequestFcParameters) GoString() string {
	return s.String()
}

func (s *PutEventTargetsRequestFcParameters) SetFunctionName(v string) *PutEventTargetsRequestFcParameters {
	s.FunctionName = &v
	return s
}

func (s *PutEventTargetsRequestFcParameters) SetId(v string) *PutEventTargetsRequestFcParameters {
	s.Id = &v
	return s
}

func (s *PutEventTargetsRequestFcParameters) SetRegion(v string) *PutEventTargetsRequestFcParameters {
	s.Region = &v
	return s
}

func (s *PutEventTargetsRequestFcParameters) SetServiceName(v string) *PutEventTargetsRequestFcParameters {
	s.ServiceName = &v
	return s
}

type PutEventTargetsRequestMnsParameters struct {
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Queue  *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s PutEventTargetsRequestMnsParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsRequestMnsParameters) GoString() string {
	return s.String()
}

func (s *PutEventTargetsRequestMnsParameters) SetId(v string) *PutEventTargetsRequestMnsParameters {
	s.Id = &v
	return s
}

func (s *PutEventTargetsRequestMnsParameters) SetQueue(v string) *PutEventTargetsRequestMnsParameters {
	s.Queue = &v
	return s
}

func (s *PutEventTargetsRequestMnsParameters) SetRegion(v string) *PutEventTargetsRequestMnsParameters {
	s.Region = &v
	return s
}

type PutEventTargetsRequestSlsParameters struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s PutEventTargetsRequestSlsParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsRequestSlsParameters) GoString() string {
	return s.String()
}

func (s *PutEventTargetsRequestSlsParameters) SetId(v string) *PutEventTargetsRequestSlsParameters {
	s.Id = &v
	return s
}

func (s *PutEventTargetsRequestSlsParameters) SetLogStore(v string) *PutEventTargetsRequestSlsParameters {
	s.LogStore = &v
	return s
}

func (s *PutEventTargetsRequestSlsParameters) SetProject(v string) *PutEventTargetsRequestSlsParameters {
	s.Project = &v
	return s
}

func (s *PutEventTargetsRequestSlsParameters) SetRegion(v string) *PutEventTargetsRequestSlsParameters {
	s.Region = &v
	return s
}

type PutEventTargetsRequestWebhookParameters struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Method   *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PutEventTargetsRequestWebhookParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsRequestWebhookParameters) GoString() string {
	return s.String()
}

func (s *PutEventTargetsRequestWebhookParameters) SetId(v string) *PutEventTargetsRequestWebhookParameters {
	s.Id = &v
	return s
}

func (s *PutEventTargetsRequestWebhookParameters) SetMethod(v string) *PutEventTargetsRequestWebhookParameters {
	s.Method = &v
	return s
}

func (s *PutEventTargetsRequestWebhookParameters) SetProtocol(v string) *PutEventTargetsRequestWebhookParameters {
	s.Protocol = &v
	return s
}

func (s *PutEventTargetsRequestWebhookParameters) SetUrl(v string) *PutEventTargetsRequestWebhookParameters {
	s.Url = &v
	return s
}

type PutEventTargetsResponseBody struct {
	Code                    *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	ContactParameters       *PutEventTargetsResponseBodyContactParameters       `json:"ContactParameters,omitempty" xml:"ContactParameters,omitempty" type:"Struct"`
	FailedContactParameters *PutEventTargetsResponseBodyFailedContactParameters `json:"FailedContactParameters,omitempty" xml:"FailedContactParameters,omitempty" type:"Struct"`
	FailedFcParameters      *PutEventTargetsResponseBodyFailedFcParameters      `json:"FailedFcParameters,omitempty" xml:"FailedFcParameters,omitempty" type:"Struct"`
	FailedMnsParameters     *PutEventTargetsResponseBodyFailedMnsParameters     `json:"FailedMnsParameters,omitempty" xml:"FailedMnsParameters,omitempty" type:"Struct"`
	FailedParameterCount    *string                                             `json:"FailedParameterCount,omitempty" xml:"FailedParameterCount,omitempty"`
	FailedSlsParameters     *PutEventTargetsResponseBodyFailedSlsParameters     `json:"FailedSlsParameters,omitempty" xml:"FailedSlsParameters,omitempty" type:"Struct"`
	FcParameters            *PutEventTargetsResponseBodyFcParameters            `json:"FcParameters,omitempty" xml:"FcParameters,omitempty" type:"Struct"`
	Message                 *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	MnsParameters           *PutEventTargetsResponseBodyMnsParameters           `json:"MnsParameters,omitempty" xml:"MnsParameters,omitempty" type:"Struct"`
	ParameterCount          *string                                             `json:"ParameterCount,omitempty" xml:"ParameterCount,omitempty"`
	RequestId               *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                 *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutEventTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *PutEventTargetsResponseBody) SetCode(v string) *PutEventTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *PutEventTargetsResponseBody) SetContactParameters(v *PutEventTargetsResponseBodyContactParameters) *PutEventTargetsResponseBody {
	s.ContactParameters = v
	return s
}

func (s *PutEventTargetsResponseBody) SetFailedContactParameters(v *PutEventTargetsResponseBodyFailedContactParameters) *PutEventTargetsResponseBody {
	s.FailedContactParameters = v
	return s
}

func (s *PutEventTargetsResponseBody) SetFailedFcParameters(v *PutEventTargetsResponseBodyFailedFcParameters) *PutEventTargetsResponseBody {
	s.FailedFcParameters = v
	return s
}

func (s *PutEventTargetsResponseBody) SetFailedMnsParameters(v *PutEventTargetsResponseBodyFailedMnsParameters) *PutEventTargetsResponseBody {
	s.FailedMnsParameters = v
	return s
}

func (s *PutEventTargetsResponseBody) SetFailedParameterCount(v string) *PutEventTargetsResponseBody {
	s.FailedParameterCount = &v
	return s
}

func (s *PutEventTargetsResponseBody) SetFailedSlsParameters(v *PutEventTargetsResponseBodyFailedSlsParameters) *PutEventTargetsResponseBody {
	s.FailedSlsParameters = v
	return s
}

func (s *PutEventTargetsResponseBody) SetFcParameters(v *PutEventTargetsResponseBodyFcParameters) *PutEventTargetsResponseBody {
	s.FcParameters = v
	return s
}

func (s *PutEventTargetsResponseBody) SetMessage(v string) *PutEventTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *PutEventTargetsResponseBody) SetMnsParameters(v *PutEventTargetsResponseBodyMnsParameters) *PutEventTargetsResponseBody {
	s.MnsParameters = v
	return s
}

func (s *PutEventTargetsResponseBody) SetParameterCount(v string) *PutEventTargetsResponseBody {
	s.ParameterCount = &v
	return s
}

func (s *PutEventTargetsResponseBody) SetRequestId(v string) *PutEventTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutEventTargetsResponseBody) SetSuccess(v bool) *PutEventTargetsResponseBody {
	s.Success = &v
	return s
}

type PutEventTargetsResponseBodyContactParameters struct {
	ContactParameter []*PutEventTargetsResponseBodyContactParametersContactParameter `json:"ContactParameter,omitempty" xml:"ContactParameter,omitempty" type:"Repeated"`
}

func (s PutEventTargetsResponseBodyContactParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsResponseBodyContactParameters) GoString() string {
	return s.String()
}

func (s *PutEventTargetsResponseBodyContactParameters) SetContactParameter(v []*PutEventTargetsResponseBodyContactParametersContactParameter) *PutEventTargetsResponseBodyContactParameters {
	s.ContactParameter = v
	return s
}

type PutEventTargetsResponseBodyContactParametersContactParameter struct {
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	Id               *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Level            *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s PutEventTargetsResponseBodyContactParametersContactParameter) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsResponseBodyContactParametersContactParameter) GoString() string {
	return s.String()
}

func (s *PutEventTargetsResponseBodyContactParametersContactParameter) SetContactGroupName(v string) *PutEventTargetsResponseBodyContactParametersContactParameter {
	s.ContactGroupName = &v
	return s
}

func (s *PutEventTargetsResponseBodyContactParametersContactParameter) SetId(v int32) *PutEventTargetsResponseBodyContactParametersContactParameter {
	s.Id = &v
	return s
}

func (s *PutEventTargetsResponseBodyContactParametersContactParameter) SetLevel(v string) *PutEventTargetsResponseBodyContactParametersContactParameter {
	s.Level = &v
	return s
}

type PutEventTargetsResponseBodyFailedContactParameters struct {
	ContactParameter []*PutEventTargetsResponseBodyFailedContactParametersContactParameter `json:"ContactParameter,omitempty" xml:"ContactParameter,omitempty" type:"Repeated"`
}

func (s PutEventTargetsResponseBodyFailedContactParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsResponseBodyFailedContactParameters) GoString() string {
	return s.String()
}

func (s *PutEventTargetsResponseBodyFailedContactParameters) SetContactParameter(v []*PutEventTargetsResponseBodyFailedContactParametersContactParameter) *PutEventTargetsResponseBodyFailedContactParameters {
	s.ContactParameter = v
	return s
}

type PutEventTargetsResponseBodyFailedContactParametersContactParameter struct {
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	Id               *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Level            *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s PutEventTargetsResponseBodyFailedContactParametersContactParameter) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsResponseBodyFailedContactParametersContactParameter) GoString() string {
	return s.String()
}

func (s *PutEventTargetsResponseBodyFailedContactParametersContactParameter) SetContactGroupName(v string) *PutEventTargetsResponseBodyFailedContactParametersContactParameter {
	s.ContactGroupName = &v
	return s
}

func (s *PutEventTargetsResponseBodyFailedContactParametersContactParameter) SetId(v int32) *PutEventTargetsResponseBodyFailedContactParametersContactParameter {
	s.Id = &v
	return s
}

func (s *PutEventTargetsResponseBodyFailedContactParametersContactParameter) SetLevel(v string) *PutEventTargetsResponseBodyFailedContactParametersContactParameter {
	s.Level = &v
	return s
}

type PutEventTargetsResponseBodyFailedFcParameters struct {
	FcParameter []*PutEventTargetsResponseBodyFailedFcParametersFcParameter `json:"FcParameter,omitempty" xml:"FcParameter,omitempty" type:"Repeated"`
}

func (s PutEventTargetsResponseBodyFailedFcParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsResponseBodyFailedFcParameters) GoString() string {
	return s.String()
}

func (s *PutEventTargetsResponseBodyFailedFcParameters) SetFcParameter(v []*PutEventTargetsResponseBodyFailedFcParametersFcParameter) *PutEventTargetsResponseBodyFailedFcParameters {
	s.FcParameter = v
	return s
}

type PutEventTargetsResponseBodyFailedFcParametersFcParameter struct {
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Id           *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ServiceName  *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s PutEventTargetsResponseBodyFailedFcParametersFcParameter) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsResponseBodyFailedFcParametersFcParameter) GoString() string {
	return s.String()
}

func (s *PutEventTargetsResponseBodyFailedFcParametersFcParameter) SetFunctionName(v string) *PutEventTargetsResponseBodyFailedFcParametersFcParameter {
	s.FunctionName = &v
	return s
}

func (s *PutEventTargetsResponseBodyFailedFcParametersFcParameter) SetId(v int32) *PutEventTargetsResponseBodyFailedFcParametersFcParameter {
	s.Id = &v
	return s
}

func (s *PutEventTargetsResponseBodyFailedFcParametersFcParameter) SetRegion(v string) *PutEventTargetsResponseBodyFailedFcParametersFcParameter {
	s.Region = &v
	return s
}

func (s *PutEventTargetsResponseBodyFailedFcParametersFcParameter) SetServiceName(v string) *PutEventTargetsResponseBodyFailedFcParametersFcParameter {
	s.ServiceName = &v
	return s
}

type PutEventTargetsResponseBodyFailedMnsParameters struct {
	MnsParameter []*PutEventTargetsResponseBodyFailedMnsParametersMnsParameter `json:"MnsParameter,omitempty" xml:"MnsParameter,omitempty" type:"Repeated"`
}

func (s PutEventTargetsResponseBodyFailedMnsParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsResponseBodyFailedMnsParameters) GoString() string {
	return s.String()
}

func (s *PutEventTargetsResponseBodyFailedMnsParameters) SetMnsParameter(v []*PutEventTargetsResponseBodyFailedMnsParametersMnsParameter) *PutEventTargetsResponseBodyFailedMnsParameters {
	s.MnsParameter = v
	return s
}

type PutEventTargetsResponseBodyFailedMnsParametersMnsParameter struct {
	Id     *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Queue  *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s PutEventTargetsResponseBodyFailedMnsParametersMnsParameter) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsResponseBodyFailedMnsParametersMnsParameter) GoString() string {
	return s.String()
}

func (s *PutEventTargetsResponseBodyFailedMnsParametersMnsParameter) SetId(v int32) *PutEventTargetsResponseBodyFailedMnsParametersMnsParameter {
	s.Id = &v
	return s
}

func (s *PutEventTargetsResponseBodyFailedMnsParametersMnsParameter) SetQueue(v string) *PutEventTargetsResponseBodyFailedMnsParametersMnsParameter {
	s.Queue = &v
	return s
}

func (s *PutEventTargetsResponseBodyFailedMnsParametersMnsParameter) SetRegion(v string) *PutEventTargetsResponseBodyFailedMnsParametersMnsParameter {
	s.Region = &v
	return s
}

type PutEventTargetsResponseBodyFailedSlsParameters struct {
	FailedSlsParameter []*PutEventTargetsResponseBodyFailedSlsParametersFailedSlsParameter `json:"FailedSlsParameter,omitempty" xml:"FailedSlsParameter,omitempty" type:"Repeated"`
}

func (s PutEventTargetsResponseBodyFailedSlsParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsResponseBodyFailedSlsParameters) GoString() string {
	return s.String()
}

func (s *PutEventTargetsResponseBodyFailedSlsParameters) SetFailedSlsParameter(v []*PutEventTargetsResponseBodyFailedSlsParametersFailedSlsParameter) *PutEventTargetsResponseBodyFailedSlsParameters {
	s.FailedSlsParameter = v
	return s
}

type PutEventTargetsResponseBodyFailedSlsParametersFailedSlsParameter struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s PutEventTargetsResponseBodyFailedSlsParametersFailedSlsParameter) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsResponseBodyFailedSlsParametersFailedSlsParameter) GoString() string {
	return s.String()
}

func (s *PutEventTargetsResponseBodyFailedSlsParametersFailedSlsParameter) SetId(v string) *PutEventTargetsResponseBodyFailedSlsParametersFailedSlsParameter {
	s.Id = &v
	return s
}

func (s *PutEventTargetsResponseBodyFailedSlsParametersFailedSlsParameter) SetLogStore(v string) *PutEventTargetsResponseBodyFailedSlsParametersFailedSlsParameter {
	s.LogStore = &v
	return s
}

func (s *PutEventTargetsResponseBodyFailedSlsParametersFailedSlsParameter) SetProject(v string) *PutEventTargetsResponseBodyFailedSlsParametersFailedSlsParameter {
	s.Project = &v
	return s
}

func (s *PutEventTargetsResponseBodyFailedSlsParametersFailedSlsParameter) SetRegion(v string) *PutEventTargetsResponseBodyFailedSlsParametersFailedSlsParameter {
	s.Region = &v
	return s
}

type PutEventTargetsResponseBodyFcParameters struct {
	FcParameter []*PutEventTargetsResponseBodyFcParametersFcParameter `json:"FcParameter,omitempty" xml:"FcParameter,omitempty" type:"Repeated"`
}

func (s PutEventTargetsResponseBodyFcParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsResponseBodyFcParameters) GoString() string {
	return s.String()
}

func (s *PutEventTargetsResponseBodyFcParameters) SetFcParameter(v []*PutEventTargetsResponseBodyFcParametersFcParameter) *PutEventTargetsResponseBodyFcParameters {
	s.FcParameter = v
	return s
}

type PutEventTargetsResponseBodyFcParametersFcParameter struct {
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Id           *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ServiceName  *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s PutEventTargetsResponseBodyFcParametersFcParameter) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsResponseBodyFcParametersFcParameter) GoString() string {
	return s.String()
}

func (s *PutEventTargetsResponseBodyFcParametersFcParameter) SetFunctionName(v string) *PutEventTargetsResponseBodyFcParametersFcParameter {
	s.FunctionName = &v
	return s
}

func (s *PutEventTargetsResponseBodyFcParametersFcParameter) SetId(v int32) *PutEventTargetsResponseBodyFcParametersFcParameter {
	s.Id = &v
	return s
}

func (s *PutEventTargetsResponseBodyFcParametersFcParameter) SetRegion(v string) *PutEventTargetsResponseBodyFcParametersFcParameter {
	s.Region = &v
	return s
}

func (s *PutEventTargetsResponseBodyFcParametersFcParameter) SetServiceName(v string) *PutEventTargetsResponseBodyFcParametersFcParameter {
	s.ServiceName = &v
	return s
}

type PutEventTargetsResponseBodyMnsParameters struct {
	MnsParameter []*PutEventTargetsResponseBodyMnsParametersMnsParameter `json:"MnsParameter,omitempty" xml:"MnsParameter,omitempty" type:"Repeated"`
}

func (s PutEventTargetsResponseBodyMnsParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsResponseBodyMnsParameters) GoString() string {
	return s.String()
}

func (s *PutEventTargetsResponseBodyMnsParameters) SetMnsParameter(v []*PutEventTargetsResponseBodyMnsParametersMnsParameter) *PutEventTargetsResponseBodyMnsParameters {
	s.MnsParameter = v
	return s
}

type PutEventTargetsResponseBodyMnsParametersMnsParameter struct {
	Id     *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Queue  *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s PutEventTargetsResponseBodyMnsParametersMnsParameter) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsResponseBodyMnsParametersMnsParameter) GoString() string {
	return s.String()
}

func (s *PutEventTargetsResponseBodyMnsParametersMnsParameter) SetId(v int32) *PutEventTargetsResponseBodyMnsParametersMnsParameter {
	s.Id = &v
	return s
}

func (s *PutEventTargetsResponseBodyMnsParametersMnsParameter) SetQueue(v string) *PutEventTargetsResponseBodyMnsParametersMnsParameter {
	s.Queue = &v
	return s
}

func (s *PutEventTargetsResponseBodyMnsParametersMnsParameter) SetRegion(v string) *PutEventTargetsResponseBodyMnsParametersMnsParameter {
	s.Region = &v
	return s
}

type PutEventTargetsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutEventTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutEventTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s PutEventTargetsResponse) GoString() string {
	return s.String()
}

func (s *PutEventTargetsResponse) SetHeaders(v map[string]*string) *PutEventTargetsResponse {
	s.Headers = v
	return s
}

func (s *PutEventTargetsResponse) SetStatusCode(v int32) *PutEventTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *PutEventTargetsResponse) SetBody(v *PutEventTargetsResponseBody) *PutEventTargetsResponse {
	s.Body = v
	return s
}

type PutMetricRuleTargetsRequest struct {
	Actions  *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	Targets []*PutMetricRuleTargetsRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s PutMetricRuleTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s PutMetricRuleTargetsRequest) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsRequest) SetActions(v string) *PutMetricRuleTargetsRequest {
	s.Actions = &v
	return s
}

func (s *PutMetricRuleTargetsRequest) SetRegionId(v string) *PutMetricRuleTargetsRequest {
	s.RegionId = &v
	return s
}

func (s *PutMetricRuleTargetsRequest) SetRuleName(v string) *PutMetricRuleTargetsRequest {
	s.RuleName = &v
	return s
}

func (s *PutMetricRuleTargetsRequest) SetTargets(v []*PutMetricRuleTargetsRequestTargets) *PutMetricRuleTargetsRequest {
	s.Targets = v
	return s
}

type PutMetricRuleTargetsRequestTargets struct {
	Arn   *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s PutMetricRuleTargetsRequestTargets) String() string {
	return tea.Prettify(s)
}

func (s PutMetricRuleTargetsRequestTargets) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsRequestTargets) SetArn(v string) *PutMetricRuleTargetsRequestTargets {
	s.Arn = &v
	return s
}

func (s *PutMetricRuleTargetsRequestTargets) SetId(v string) *PutMetricRuleTargetsRequestTargets {
	s.Id = &v
	return s
}

func (s *PutMetricRuleTargetsRequestTargets) SetLevel(v string) *PutMetricRuleTargetsRequestTargets {
	s.Level = &v
	return s
}

type PutMetricRuleTargetsResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	FailData  *PutMetricRuleTargetsResponseBodyFailData `json:"FailData,omitempty" xml:"FailData,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutMetricRuleTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutMetricRuleTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsResponseBody) SetCode(v string) *PutMetricRuleTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBody) SetFailData(v *PutMetricRuleTargetsResponseBodyFailData) *PutMetricRuleTargetsResponseBody {
	s.FailData = v
	return s
}

func (s *PutMetricRuleTargetsResponseBody) SetMessage(v string) *PutMetricRuleTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBody) SetRequestId(v string) *PutMetricRuleTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBody) SetSuccess(v bool) *PutMetricRuleTargetsResponseBody {
	s.Success = &v
	return s
}

type PutMetricRuleTargetsResponseBodyFailData struct {
	Targets *PutMetricRuleTargetsResponseBodyFailDataTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Struct"`
}

func (s PutMetricRuleTargetsResponseBodyFailData) String() string {
	return tea.Prettify(s)
}

func (s PutMetricRuleTargetsResponseBodyFailData) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsResponseBodyFailData) SetTargets(v *PutMetricRuleTargetsResponseBodyFailDataTargets) *PutMetricRuleTargetsResponseBodyFailData {
	s.Targets = v
	return s
}

type PutMetricRuleTargetsResponseBodyFailDataTargets struct {
	Target []*PutMetricRuleTargetsResponseBodyFailDataTargetsTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s PutMetricRuleTargetsResponseBodyFailDataTargets) String() string {
	return tea.Prettify(s)
}

func (s PutMetricRuleTargetsResponseBodyFailDataTargets) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargets) SetTarget(v []*PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) *PutMetricRuleTargetsResponseBodyFailDataTargets {
	s.Target = v
	return s
}

type PutMetricRuleTargetsResponseBodyFailDataTargetsTarget struct {
	Arn   *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) String() string {
	return tea.Prettify(s)
}

func (s PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) SetArn(v string) *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget {
	s.Arn = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) SetId(v string) *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget {
	s.Id = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) SetLevel(v string) *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget {
	s.Level = &v
	return s
}

type PutMetricRuleTargetsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutMetricRuleTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutMetricRuleTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s PutMetricRuleTargetsResponse) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsResponse) SetHeaders(v map[string]*string) *PutMetricRuleTargetsResponse {
	s.Headers = v
	return s
}

func (s *PutMetricRuleTargetsResponse) SetStatusCode(v int32) *PutMetricRuleTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *PutMetricRuleTargetsResponse) SetBody(v *PutMetricRuleTargetsResponseBody) *PutMetricRuleTargetsResponse {
	s.Body = v
	return s
}

type PutResourceMetricRuleRequest struct {
	Escalations       *PutResourceMetricRuleRequestEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	ContactGroups     *string                                  `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	EffectiveInterval *string                                  `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	EmailSubject      *string                                  `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
	Interval          *string                                  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter is required.
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// This parameter is required.
	Namespace           *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	Period              *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// This parameter is required.
	RuleId      *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName    *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	SilenceTime *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	Webhook     *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s PutResourceMetricRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRuleRequest) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequest) SetEscalations(v *PutResourceMetricRuleRequestEscalations) *PutResourceMetricRuleRequest {
	s.Escalations = v
	return s
}

func (s *PutResourceMetricRuleRequest) SetContactGroups(v string) *PutResourceMetricRuleRequest {
	s.ContactGroups = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetEffectiveInterval(v string) *PutResourceMetricRuleRequest {
	s.EffectiveInterval = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetEmailSubject(v string) *PutResourceMetricRuleRequest {
	s.EmailSubject = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetInterval(v string) *PutResourceMetricRuleRequest {
	s.Interval = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetMetricName(v string) *PutResourceMetricRuleRequest {
	s.MetricName = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetNamespace(v string) *PutResourceMetricRuleRequest {
	s.Namespace = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetNoEffectiveInterval(v string) *PutResourceMetricRuleRequest {
	s.NoEffectiveInterval = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetPeriod(v string) *PutResourceMetricRuleRequest {
	s.Period = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetResources(v string) *PutResourceMetricRuleRequest {
	s.Resources = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetRuleId(v string) *PutResourceMetricRuleRequest {
	s.RuleId = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetRuleName(v string) *PutResourceMetricRuleRequest {
	s.RuleName = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetSilenceTime(v int32) *PutResourceMetricRuleRequest {
	s.SilenceTime = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetWebhook(v string) *PutResourceMetricRuleRequest {
	s.Webhook = &v
	return s
}

type PutResourceMetricRuleRequestEscalations struct {
	Critical *PutResourceMetricRuleRequestEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *PutResourceMetricRuleRequestEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *PutResourceMetricRuleRequestEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s PutResourceMetricRuleRequestEscalations) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRuleRequestEscalations) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestEscalations) SetCritical(v *PutResourceMetricRuleRequestEscalationsCritical) *PutResourceMetricRuleRequestEscalations {
	s.Critical = v
	return s
}

func (s *PutResourceMetricRuleRequestEscalations) SetInfo(v *PutResourceMetricRuleRequestEscalationsInfo) *PutResourceMetricRuleRequestEscalations {
	s.Info = v
	return s
}

func (s *PutResourceMetricRuleRequestEscalations) SetWarn(v *PutResourceMetricRuleRequestEscalationsWarn) *PutResourceMetricRuleRequestEscalations {
	s.Warn = v
	return s
}

type PutResourceMetricRuleRequestEscalationsCritical struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRuleRequestEscalationsCritical) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRuleRequestEscalationsCritical) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) SetComparisonOperator(v string) *PutResourceMetricRuleRequestEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) SetStatistics(v string) *PutResourceMetricRuleRequestEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) SetThreshold(v string) *PutResourceMetricRuleRequestEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) SetTimes(v int32) *PutResourceMetricRuleRequestEscalationsCritical {
	s.Times = &v
	return s
}

type PutResourceMetricRuleRequestEscalationsInfo struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRuleRequestEscalationsInfo) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRuleRequestEscalationsInfo) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) SetComparisonOperator(v string) *PutResourceMetricRuleRequestEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) SetStatistics(v string) *PutResourceMetricRuleRequestEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) SetThreshold(v string) *PutResourceMetricRuleRequestEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) SetTimes(v int32) *PutResourceMetricRuleRequestEscalationsInfo {
	s.Times = &v
	return s
}

type PutResourceMetricRuleRequestEscalationsWarn struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRuleRequestEscalationsWarn) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRuleRequestEscalationsWarn) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) SetComparisonOperator(v string) *PutResourceMetricRuleRequestEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) SetStatistics(v string) *PutResourceMetricRuleRequestEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) SetThreshold(v string) *PutResourceMetricRuleRequestEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) SetTimes(v int32) *PutResourceMetricRuleRequestEscalationsWarn {
	s.Times = &v
	return s
}

type PutResourceMetricRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutResourceMetricRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleResponseBody) SetCode(v string) *PutResourceMetricRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PutResourceMetricRuleResponseBody) SetMessage(v string) *PutResourceMetricRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PutResourceMetricRuleResponseBody) SetRequestId(v string) *PutResourceMetricRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutResourceMetricRuleResponseBody) SetSuccess(v bool) *PutResourceMetricRuleResponseBody {
	s.Success = &v
	return s
}

type PutResourceMetricRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutResourceMetricRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutResourceMetricRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRuleResponse) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleResponse) SetHeaders(v map[string]*string) *PutResourceMetricRuleResponse {
	s.Headers = v
	return s
}

func (s *PutResourceMetricRuleResponse) SetStatusCode(v int32) *PutResourceMetricRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PutResourceMetricRuleResponse) SetBody(v *PutResourceMetricRuleResponseBody) *PutResourceMetricRuleResponse {
	s.Body = v
	return s
}

type QueryCustomMetricListRequest struct {
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// This parameter is required.
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Md5        *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Page       *string `json:"Page,omitempty" xml:"Page,omitempty"`
	Size       *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s QueryCustomMetricListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomMetricListRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomMetricListRequest) SetDimension(v string) *QueryCustomMetricListRequest {
	s.Dimension = &v
	return s
}

func (s *QueryCustomMetricListRequest) SetGroupId(v string) *QueryCustomMetricListRequest {
	s.GroupId = &v
	return s
}

func (s *QueryCustomMetricListRequest) SetMd5(v string) *QueryCustomMetricListRequest {
	s.Md5 = &v
	return s
}

func (s *QueryCustomMetricListRequest) SetMetricName(v string) *QueryCustomMetricListRequest {
	s.MetricName = &v
	return s
}

func (s *QueryCustomMetricListRequest) SetPage(v string) *QueryCustomMetricListRequest {
	s.Page = &v
	return s
}

func (s *QueryCustomMetricListRequest) SetSize(v string) *QueryCustomMetricListRequest {
	s.Size = &v
	return s
}

type QueryCustomMetricListResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s QueryCustomMetricListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomMetricListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomMetricListResponseBody) SetCode(v string) *QueryCustomMetricListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCustomMetricListResponseBody) SetMessage(v string) *QueryCustomMetricListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCustomMetricListResponseBody) SetRequestId(v string) *QueryCustomMetricListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCustomMetricListResponseBody) SetResult(v string) *QueryCustomMetricListResponseBody {
	s.Result = &v
	return s
}

type QueryCustomMetricListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCustomMetricListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCustomMetricListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomMetricListResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomMetricListResponse) SetHeaders(v map[string]*string) *QueryCustomMetricListResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomMetricListResponse) SetStatusCode(v int32) *QueryCustomMetricListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCustomMetricListResponse) SetBody(v *QueryCustomMetricListResponseBody) *QueryCustomMetricListResponse {
	s.Body = v
	return s
}

type QueryMetricDataRequest struct {
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Express    *string `json:"Express,omitempty" xml:"Express,omitempty"`
	Length     *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// This parameter is required.
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryMetricDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricDataRequest) GoString() string {
	return s.String()
}

func (s *QueryMetricDataRequest) SetDimensions(v string) *QueryMetricDataRequest {
	s.Dimensions = &v
	return s
}

func (s *QueryMetricDataRequest) SetEndTime(v string) *QueryMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMetricDataRequest) SetExpress(v string) *QueryMetricDataRequest {
	s.Express = &v
	return s
}

func (s *QueryMetricDataRequest) SetLength(v string) *QueryMetricDataRequest {
	s.Length = &v
	return s
}

func (s *QueryMetricDataRequest) SetMetric(v string) *QueryMetricDataRequest {
	s.Metric = &v
	return s
}

func (s *QueryMetricDataRequest) SetPeriod(v string) *QueryMetricDataRequest {
	s.Period = &v
	return s
}

func (s *QueryMetricDataRequest) SetProject(v string) *QueryMetricDataRequest {
	s.Project = &v
	return s
}

func (s *QueryMetricDataRequest) SetRegionId(v string) *QueryMetricDataRequest {
	s.RegionId = &v
	return s
}

func (s *QueryMetricDataRequest) SetResourceOwnerId(v int64) *QueryMetricDataRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryMetricDataRequest) SetStartTime(v string) *QueryMetricDataRequest {
	s.StartTime = &v
	return s
}

type QueryMetricDataResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMetricDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMetricDataResponseBody) SetCode(v string) *QueryMetricDataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMetricDataResponseBody) SetDatapoints(v string) *QueryMetricDataResponseBody {
	s.Datapoints = &v
	return s
}

func (s *QueryMetricDataResponseBody) SetMessage(v string) *QueryMetricDataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMetricDataResponseBody) SetPeriod(v string) *QueryMetricDataResponseBody {
	s.Period = &v
	return s
}

func (s *QueryMetricDataResponseBody) SetRequestId(v string) *QueryMetricDataResponseBody {
	s.RequestId = &v
	return s
}

type QueryMetricDataResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMetricDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricDataResponse) GoString() string {
	return s.String()
}

func (s *QueryMetricDataResponse) SetHeaders(v map[string]*string) *QueryMetricDataResponse {
	s.Headers = v
	return s
}

func (s *QueryMetricDataResponse) SetStatusCode(v int32) *QueryMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMetricDataResponse) SetBody(v *QueryMetricDataResponseBody) *QueryMetricDataResponse {
	s.Body = v
	return s
}

type QueryMetricLastRequest struct {
	// example:
	//
	// 1000
	Cursor *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	// example:
	//
	// [{"instanceId":"XXX"}]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// example:
	//
	// 2019-01-31 10:10:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// {"groupby":["userId","instanceId"]}
	Express *string `json:"Express,omitempty" xml:"Express,omitempty"`
	// example:
	//
	// 1000
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CPUUtilization
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// example:
	//
	// 1
	Page *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 000
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 2019-01-31 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryMetricLastRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricLastRequest) GoString() string {
	return s.String()
}

func (s *QueryMetricLastRequest) SetCursor(v string) *QueryMetricLastRequest {
	s.Cursor = &v
	return s
}

func (s *QueryMetricLastRequest) SetDimensions(v string) *QueryMetricLastRequest {
	s.Dimensions = &v
	return s
}

func (s *QueryMetricLastRequest) SetEndTime(v string) *QueryMetricLastRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMetricLastRequest) SetExpress(v string) *QueryMetricLastRequest {
	s.Express = &v
	return s
}

func (s *QueryMetricLastRequest) SetLength(v string) *QueryMetricLastRequest {
	s.Length = &v
	return s
}

func (s *QueryMetricLastRequest) SetMetric(v string) *QueryMetricLastRequest {
	s.Metric = &v
	return s
}

func (s *QueryMetricLastRequest) SetPage(v string) *QueryMetricLastRequest {
	s.Page = &v
	return s
}

func (s *QueryMetricLastRequest) SetPeriod(v string) *QueryMetricLastRequest {
	s.Period = &v
	return s
}

func (s *QueryMetricLastRequest) SetProject(v string) *QueryMetricLastRequest {
	s.Project = &v
	return s
}

func (s *QueryMetricLastRequest) SetRegionId(v string) *QueryMetricLastRequest {
	s.RegionId = &v
	return s
}

func (s *QueryMetricLastRequest) SetResourceOwnerId(v int64) *QueryMetricLastRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryMetricLastRequest) SetStartTime(v string) *QueryMetricLastRequest {
	s.StartTime = &v
	return s
}

type QueryMetricLastResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 10000
	Cursor *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	// example:
	//
	// [{"timestamp":1548900600000,"userId":"000","instanceId":"abc","Minimum":6.3,"Average":6.3,"Maximum":6.3}]
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 021472A6-25E3-4094-8D00-BA4B6A5486C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMetricLastResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricLastResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMetricLastResponseBody) SetCode(v string) *QueryMetricLastResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMetricLastResponseBody) SetCursor(v string) *QueryMetricLastResponseBody {
	s.Cursor = &v
	return s
}

func (s *QueryMetricLastResponseBody) SetDatapoints(v string) *QueryMetricLastResponseBody {
	s.Datapoints = &v
	return s
}

func (s *QueryMetricLastResponseBody) SetMessage(v string) *QueryMetricLastResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMetricLastResponseBody) SetPeriod(v string) *QueryMetricLastResponseBody {
	s.Period = &v
	return s
}

func (s *QueryMetricLastResponseBody) SetRequestId(v string) *QueryMetricLastResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMetricLastResponseBody) SetSuccess(v string) *QueryMetricLastResponseBody {
	s.Success = &v
	return s
}

type QueryMetricLastResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMetricLastResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMetricLastResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricLastResponse) GoString() string {
	return s.String()
}

func (s *QueryMetricLastResponse) SetHeaders(v map[string]*string) *QueryMetricLastResponse {
	s.Headers = v
	return s
}

func (s *QueryMetricLastResponse) SetStatusCode(v int32) *QueryMetricLastResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMetricLastResponse) SetBody(v *QueryMetricLastResponseBody) *QueryMetricLastResponse {
	s.Body = v
	return s
}

type QueryMetricListRequest struct {
	// example:
	//
	// 1000
	Cursor *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	// example:
	//
	// [{"instanceId": "i-abcdefgh123456"}]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// example:
	//
	// 2019-01-30 00:10:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// “{"groupby":["userId","instanceId"]}”
	Express *string `json:"Express,omitempty" xml:"Express,omitempty"`
	// example:
	//
	// 1000
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cpu_idle
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 000
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 2019-01-30 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryMetricListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricListRequest) GoString() string {
	return s.String()
}

func (s *QueryMetricListRequest) SetCursor(v string) *QueryMetricListRequest {
	s.Cursor = &v
	return s
}

func (s *QueryMetricListRequest) SetDimensions(v string) *QueryMetricListRequest {
	s.Dimensions = &v
	return s
}

func (s *QueryMetricListRequest) SetEndTime(v string) *QueryMetricListRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMetricListRequest) SetExpress(v string) *QueryMetricListRequest {
	s.Express = &v
	return s
}

func (s *QueryMetricListRequest) SetLength(v string) *QueryMetricListRequest {
	s.Length = &v
	return s
}

func (s *QueryMetricListRequest) SetMetric(v string) *QueryMetricListRequest {
	s.Metric = &v
	return s
}

func (s *QueryMetricListRequest) SetPeriod(v string) *QueryMetricListRequest {
	s.Period = &v
	return s
}

func (s *QueryMetricListRequest) SetProject(v string) *QueryMetricListRequest {
	s.Project = &v
	return s
}

func (s *QueryMetricListRequest) SetRegionId(v string) *QueryMetricListRequest {
	s.RegionId = &v
	return s
}

func (s *QueryMetricListRequest) SetResourceOwnerId(v int64) *QueryMetricListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryMetricListRequest) SetStartTime(v string) *QueryMetricListRequest {
	s.StartTime = &v
	return s
}

type QueryMetricListResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1000
	Cursor *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	// example:
	//
	// [{"timestamp":1548777660000,"userId":"123","instanceId":"i-abc","Minimum":9.92,"Average":9.92,"Maximum":9.92}]
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 3121AE7D-4AFF-4C25-8F1D-C8226EBB1F42
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMetricListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMetricListResponseBody) SetCode(v string) *QueryMetricListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMetricListResponseBody) SetCursor(v string) *QueryMetricListResponseBody {
	s.Cursor = &v
	return s
}

func (s *QueryMetricListResponseBody) SetDatapoints(v string) *QueryMetricListResponseBody {
	s.Datapoints = &v
	return s
}

func (s *QueryMetricListResponseBody) SetMessage(v string) *QueryMetricListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMetricListResponseBody) SetPeriod(v string) *QueryMetricListResponseBody {
	s.Period = &v
	return s
}

func (s *QueryMetricListResponseBody) SetRequestId(v string) *QueryMetricListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMetricListResponseBody) SetSuccess(v string) *QueryMetricListResponseBody {
	s.Success = &v
	return s
}

type QueryMetricListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMetricListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMetricListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricListResponse) GoString() string {
	return s.String()
}

func (s *QueryMetricListResponse) SetHeaders(v map[string]*string) *QueryMetricListResponse {
	s.Headers = v
	return s
}

func (s *QueryMetricListResponse) SetStatusCode(v int32) *QueryMetricListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMetricListResponse) SetBody(v *QueryMetricListResponseBody) *QueryMetricListResponse {
	s.Body = v
	return s
}

type QueryMetricMetaRequest struct {
	// example:
	//
	// [{\\"name\\":\\"alertUnit\\",\\"value\\":\\"%\\"}]
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// CPUUtilization
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// acs_ecs_dashboard
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s QueryMetricMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricMetaRequest) GoString() string {
	return s.String()
}

func (s *QueryMetricMetaRequest) SetLabels(v string) *QueryMetricMetaRequest {
	s.Labels = &v
	return s
}

func (s *QueryMetricMetaRequest) SetMetric(v string) *QueryMetricMetaRequest {
	s.Metric = &v
	return s
}

func (s *QueryMetricMetaRequest) SetPageNumber(v int32) *QueryMetricMetaRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryMetricMetaRequest) SetPageSize(v int32) *QueryMetricMetaRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMetricMetaRequest) SetProject(v string) *QueryMetricMetaRequest {
	s.Project = &v
	return s
}

type QueryMetricMetaResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 0CCE0AF0-053C-4B13-A583-DC9A85785D49
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources *QueryMetricMetaResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMetricMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricMetaResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMetricMetaResponseBody) SetErrorCode(v string) *QueryMetricMetaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryMetricMetaResponseBody) SetErrorMessage(v string) *QueryMetricMetaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryMetricMetaResponseBody) SetRequestId(v string) *QueryMetricMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMetricMetaResponseBody) SetResources(v *QueryMetricMetaResponseBodyResources) *QueryMetricMetaResponseBody {
	s.Resources = v
	return s
}

func (s *QueryMetricMetaResponseBody) SetSuccess(v bool) *QueryMetricMetaResponseBody {
	s.Success = &v
	return s
}

type QueryMetricMetaResponseBodyResources struct {
	Resource []*QueryMetricMetaResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s QueryMetricMetaResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricMetaResponseBodyResources) GoString() string {
	return s.String()
}

func (s *QueryMetricMetaResponseBodyResources) SetResource(v []*QueryMetricMetaResponseBodyResourcesResource) *QueryMetricMetaResponseBodyResources {
	s.Resource = v
	return s
}

type QueryMetricMetaResponseBodyResourcesResource struct {
	// example:
	//
	// ECS.CPUUtilization
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// instanceId
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// example:
	//
	// [{\\"name\\":\\"alertUnit\\",\\"value\\":\\"%\\"},{\\"name\\":\\"alertDefault\\",\\"value\\":\\"80\\"},{\\"name\\":\\"minAlertPeriod\\",\\"value\\":\\"60\\"},{\\"name\\":\\"metricCategory\\",\\"value\\":\\"instanceId\\"},{\\"name\\":\\"is_alarm\\",\\"value\\":\\"true\\"}]"
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// CPUUtilization
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// example:
	//
	// 60,300
	Periods *string `json:"Periods,omitempty" xml:"Periods,omitempty"`
	// example:
	//
	// acs_ecs_dashboard
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// example:
	//
	// Average,Minimum,Maximum
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// example:
	//
	// %
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s QueryMetricMetaResponseBodyResourcesResource) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricMetaResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *QueryMetricMetaResponseBodyResourcesResource) SetDescription(v string) *QueryMetricMetaResponseBodyResourcesResource {
	s.Description = &v
	return s
}

func (s *QueryMetricMetaResponseBodyResourcesResource) SetDimensions(v string) *QueryMetricMetaResponseBodyResourcesResource {
	s.Dimensions = &v
	return s
}

func (s *QueryMetricMetaResponseBodyResourcesResource) SetLabels(v string) *QueryMetricMetaResponseBodyResourcesResource {
	s.Labels = &v
	return s
}

func (s *QueryMetricMetaResponseBodyResourcesResource) SetMetric(v string) *QueryMetricMetaResponseBodyResourcesResource {
	s.Metric = &v
	return s
}

func (s *QueryMetricMetaResponseBodyResourcesResource) SetPeriods(v string) *QueryMetricMetaResponseBodyResourcesResource {
	s.Periods = &v
	return s
}

func (s *QueryMetricMetaResponseBodyResourcesResource) SetProject(v string) *QueryMetricMetaResponseBodyResourcesResource {
	s.Project = &v
	return s
}

func (s *QueryMetricMetaResponseBodyResourcesResource) SetStatistics(v string) *QueryMetricMetaResponseBodyResourcesResource {
	s.Statistics = &v
	return s
}

func (s *QueryMetricMetaResponseBodyResourcesResource) SetUnit(v string) *QueryMetricMetaResponseBodyResourcesResource {
	s.Unit = &v
	return s
}

type QueryMetricMetaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMetricMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMetricMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricMetaResponse) GoString() string {
	return s.String()
}

func (s *QueryMetricMetaResponse) SetHeaders(v map[string]*string) *QueryMetricMetaResponse {
	s.Headers = v
	return s
}

func (s *QueryMetricMetaResponse) SetStatusCode(v int32) *QueryMetricMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMetricMetaResponse) SetBody(v *QueryMetricMetaResponseBody) *QueryMetricMetaResponse {
	s.Body = v
	return s
}

type QueryMetricTopRequest struct {
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Express    *string `json:"Express,omitempty" xml:"Express,omitempty"`
	Length     *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// This parameter is required.
	Metric    *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	OrderDesc *string `json:"OrderDesc,omitempty" xml:"OrderDesc,omitempty"`
	Orderby   *string `json:"Orderby,omitempty" xml:"Orderby,omitempty"`
	Period    *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	Project         *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryMetricTopRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricTopRequest) GoString() string {
	return s.String()
}

func (s *QueryMetricTopRequest) SetDimensions(v string) *QueryMetricTopRequest {
	s.Dimensions = &v
	return s
}

func (s *QueryMetricTopRequest) SetEndTime(v string) *QueryMetricTopRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMetricTopRequest) SetExpress(v string) *QueryMetricTopRequest {
	s.Express = &v
	return s
}

func (s *QueryMetricTopRequest) SetLength(v string) *QueryMetricTopRequest {
	s.Length = &v
	return s
}

func (s *QueryMetricTopRequest) SetMetric(v string) *QueryMetricTopRequest {
	s.Metric = &v
	return s
}

func (s *QueryMetricTopRequest) SetOrderDesc(v string) *QueryMetricTopRequest {
	s.OrderDesc = &v
	return s
}

func (s *QueryMetricTopRequest) SetOrderby(v string) *QueryMetricTopRequest {
	s.Orderby = &v
	return s
}

func (s *QueryMetricTopRequest) SetPeriod(v string) *QueryMetricTopRequest {
	s.Period = &v
	return s
}

func (s *QueryMetricTopRequest) SetProject(v string) *QueryMetricTopRequest {
	s.Project = &v
	return s
}

func (s *QueryMetricTopRequest) SetRegionId(v string) *QueryMetricTopRequest {
	s.RegionId = &v
	return s
}

func (s *QueryMetricTopRequest) SetResourceOwnerId(v int64) *QueryMetricTopRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryMetricTopRequest) SetStartTime(v string) *QueryMetricTopRequest {
	s.StartTime = &v
	return s
}

type QueryMetricTopResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMetricTopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricTopResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMetricTopResponseBody) SetCode(v string) *QueryMetricTopResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMetricTopResponseBody) SetDatapoints(v string) *QueryMetricTopResponseBody {
	s.Datapoints = &v
	return s
}

func (s *QueryMetricTopResponseBody) SetMessage(v string) *QueryMetricTopResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMetricTopResponseBody) SetPeriod(v string) *QueryMetricTopResponseBody {
	s.Period = &v
	return s
}

func (s *QueryMetricTopResponseBody) SetRequestId(v string) *QueryMetricTopResponseBody {
	s.RequestId = &v
	return s
}

type QueryMetricTopResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMetricTopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMetricTopResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricTopResponse) GoString() string {
	return s.String()
}

func (s *QueryMetricTopResponse) SetHeaders(v map[string]*string) *QueryMetricTopResponse {
	s.Headers = v
	return s
}

func (s *QueryMetricTopResponse) SetStatusCode(v int32) *QueryMetricTopResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMetricTopResponse) SetBody(v *QueryMetricTopResponseBody) *QueryMetricTopResponse {
	s.Body = v
	return s
}

type QueryProjectMetaRequest struct {
	// example:
	//
	// [{\\"name\\":\\"product\\",\\"value\\":\\"MongoDB\\"]
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryProjectMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectMetaRequest) GoString() string {
	return s.String()
}

func (s *QueryProjectMetaRequest) SetLabels(v string) *QueryProjectMetaRequest {
	s.Labels = &v
	return s
}

func (s *QueryProjectMetaRequest) SetPageNumber(v int32) *QueryProjectMetaRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryProjectMetaRequest) SetPageSize(v int32) *QueryProjectMetaRequest {
	s.PageSize = &v
	return s
}

type QueryProjectMetaResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// D3DBF9F5-7C4D-4A67-B869-097C069C481D
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources *QueryProjectMetaResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryProjectMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectMetaResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProjectMetaResponseBody) SetErrorCode(v string) *QueryProjectMetaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryProjectMetaResponseBody) SetErrorMessage(v string) *QueryProjectMetaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryProjectMetaResponseBody) SetRequestId(v string) *QueryProjectMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryProjectMetaResponseBody) SetResources(v *QueryProjectMetaResponseBodyResources) *QueryProjectMetaResponseBody {
	s.Resources = v
	return s
}

func (s *QueryProjectMetaResponseBody) SetSuccess(v bool) *QueryProjectMetaResponseBody {
	s.Success = &v
	return s
}

type QueryProjectMetaResponseBodyResources struct {
	Resource []*QueryProjectMetaResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s QueryProjectMetaResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectMetaResponseBodyResources) GoString() string {
	return s.String()
}

func (s *QueryProjectMetaResponseBodyResources) SetResource(v []*QueryProjectMetaResponseBodyResourcesResource) *QueryProjectMetaResponseBodyResources {
	s.Resource = v
	return s
}

type QueryProjectMetaResponseBodyResourcesResource struct {
	// example:
	//
	// ApsaraDB for MongoDB
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// [{\\"name\\":\\"product\\",\\"value\\":\\"MongoDB\\"]
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// acs_mongodb
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s QueryProjectMetaResponseBodyResourcesResource) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectMetaResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *QueryProjectMetaResponseBodyResourcesResource) SetDescription(v string) *QueryProjectMetaResponseBodyResourcesResource {
	s.Description = &v
	return s
}

func (s *QueryProjectMetaResponseBodyResourcesResource) SetLabels(v string) *QueryProjectMetaResponseBodyResourcesResource {
	s.Labels = &v
	return s
}

func (s *QueryProjectMetaResponseBodyResourcesResource) SetProject(v string) *QueryProjectMetaResponseBodyResourcesResource {
	s.Project = &v
	return s
}

type QueryProjectMetaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProjectMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProjectMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectMetaResponse) GoString() string {
	return s.String()
}

func (s *QueryProjectMetaResponse) SetHeaders(v map[string]*string) *QueryProjectMetaResponse {
	s.Headers = v
	return s
}

func (s *QueryProjectMetaResponse) SetStatusCode(v int32) *QueryProjectMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProjectMetaResponse) SetBody(v *QueryProjectMetaResponseBody) *QueryProjectMetaResponse {
	s.Body = v
	return s
}

type QueryStaticsAvailabilityRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TimeRange *string `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s QueryStaticsAvailabilityRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryStaticsAvailabilityRequest) GoString() string {
	return s.String()
}

func (s *QueryStaticsAvailabilityRequest) SetRegionId(v string) *QueryStaticsAvailabilityRequest {
	s.RegionId = &v
	return s
}

func (s *QueryStaticsAvailabilityRequest) SetTaskId(v string) *QueryStaticsAvailabilityRequest {
	s.TaskId = &v
	return s
}

func (s *QueryStaticsAvailabilityRequest) SetTimeRange(v string) *QueryStaticsAvailabilityRequest {
	s.TimeRange = &v
	return s
}

type QueryStaticsAvailabilityResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryStaticsAvailabilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryStaticsAvailabilityResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStaticsAvailabilityResponseBody) SetCode(v string) *QueryStaticsAvailabilityResponseBody {
	s.Code = &v
	return s
}

func (s *QueryStaticsAvailabilityResponseBody) SetData(v string) *QueryStaticsAvailabilityResponseBody {
	s.Data = &v
	return s
}

func (s *QueryStaticsAvailabilityResponseBody) SetMessage(v string) *QueryStaticsAvailabilityResponseBody {
	s.Message = &v
	return s
}

func (s *QueryStaticsAvailabilityResponseBody) SetRequestId(v string) *QueryStaticsAvailabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStaticsAvailabilityResponseBody) SetSuccess(v string) *QueryStaticsAvailabilityResponseBody {
	s.Success = &v
	return s
}

type QueryStaticsAvailabilityResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryStaticsAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryStaticsAvailabilityResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryStaticsAvailabilityResponse) GoString() string {
	return s.String()
}

func (s *QueryStaticsAvailabilityResponse) SetHeaders(v map[string]*string) *QueryStaticsAvailabilityResponse {
	s.Headers = v
	return s
}

func (s *QueryStaticsAvailabilityResponse) SetStatusCode(v int32) *QueryStaticsAvailabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStaticsAvailabilityResponse) SetBody(v *QueryStaticsAvailabilityResponseBody) *QueryStaticsAvailabilityResponse {
	s.Body = v
	return s
}

type QueryStaticsResponseTimeRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TimeRange *string `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s QueryStaticsResponseTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryStaticsResponseTimeRequest) GoString() string {
	return s.String()
}

func (s *QueryStaticsResponseTimeRequest) SetRegionId(v string) *QueryStaticsResponseTimeRequest {
	s.RegionId = &v
	return s
}

func (s *QueryStaticsResponseTimeRequest) SetTaskId(v string) *QueryStaticsResponseTimeRequest {
	s.TaskId = &v
	return s
}

func (s *QueryStaticsResponseTimeRequest) SetTimeRange(v string) *QueryStaticsResponseTimeRequest {
	s.TimeRange = &v
	return s
}

type QueryStaticsResponseTimeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryStaticsResponseTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryStaticsResponseTimeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStaticsResponseTimeResponseBody) SetCode(v string) *QueryStaticsResponseTimeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryStaticsResponseTimeResponseBody) SetData(v string) *QueryStaticsResponseTimeResponseBody {
	s.Data = &v
	return s
}

func (s *QueryStaticsResponseTimeResponseBody) SetMessage(v string) *QueryStaticsResponseTimeResponseBody {
	s.Message = &v
	return s
}

func (s *QueryStaticsResponseTimeResponseBody) SetRequestId(v string) *QueryStaticsResponseTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStaticsResponseTimeResponseBody) SetSuccess(v string) *QueryStaticsResponseTimeResponseBody {
	s.Success = &v
	return s
}

type QueryStaticsResponseTimeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryStaticsResponseTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryStaticsResponseTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryStaticsResponseTimeResponse) GoString() string {
	return s.String()
}

func (s *QueryStaticsResponseTimeResponse) SetHeaders(v map[string]*string) *QueryStaticsResponseTimeResponse {
	s.Headers = v
	return s
}

func (s *QueryStaticsResponseTimeResponse) SetStatusCode(v int32) *QueryStaticsResponseTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStaticsResponseTimeResponse) SetBody(v *QueryStaticsResponseTimeResponseBody) *QueryStaticsResponseTimeResponse {
	s.Body = v
	return s
}

type QuerySystemEventCountRequest struct {
	QueryJson *string `json:"QueryJson,omitempty" xml:"QueryJson,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QuerySystemEventCountRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySystemEventCountRequest) GoString() string {
	return s.String()
}

func (s *QuerySystemEventCountRequest) SetQueryJson(v string) *QuerySystemEventCountRequest {
	s.QueryJson = &v
	return s
}

func (s *QuerySystemEventCountRequest) SetRegionId(v string) *QuerySystemEventCountRequest {
	s.RegionId = &v
	return s
}

type QuerySystemEventCountResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySystemEventCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySystemEventCountResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySystemEventCountResponseBody) SetCode(v string) *QuerySystemEventCountResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySystemEventCountResponseBody) SetData(v string) *QuerySystemEventCountResponseBody {
	s.Data = &v
	return s
}

func (s *QuerySystemEventCountResponseBody) SetMessage(v string) *QuerySystemEventCountResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySystemEventCountResponseBody) SetRequestId(v string) *QuerySystemEventCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySystemEventCountResponseBody) SetSuccess(v string) *QuerySystemEventCountResponseBody {
	s.Success = &v
	return s
}

type QuerySystemEventCountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySystemEventCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySystemEventCountResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySystemEventCountResponse) GoString() string {
	return s.String()
}

func (s *QuerySystemEventCountResponse) SetHeaders(v map[string]*string) *QuerySystemEventCountResponse {
	s.Headers = v
	return s
}

func (s *QuerySystemEventCountResponse) SetStatusCode(v int32) *QuerySystemEventCountResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySystemEventCountResponse) SetBody(v *QuerySystemEventCountResponseBody) *QuerySystemEventCountResponse {
	s.Body = v
	return s
}

type QuerySystemEventDemoRequest struct {
	// This parameter is required.
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// This parameter is required.
	Product  *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QuerySystemEventDemoRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySystemEventDemoRequest) GoString() string {
	return s.String()
}

func (s *QuerySystemEventDemoRequest) SetEventName(v string) *QuerySystemEventDemoRequest {
	s.EventName = &v
	return s
}

func (s *QuerySystemEventDemoRequest) SetProduct(v string) *QuerySystemEventDemoRequest {
	s.Product = &v
	return s
}

func (s *QuerySystemEventDemoRequest) SetRegionId(v string) *QuerySystemEventDemoRequest {
	s.RegionId = &v
	return s
}

type QuerySystemEventDemoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySystemEventDemoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySystemEventDemoResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySystemEventDemoResponseBody) SetCode(v string) *QuerySystemEventDemoResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySystemEventDemoResponseBody) SetData(v string) *QuerySystemEventDemoResponseBody {
	s.Data = &v
	return s
}

func (s *QuerySystemEventDemoResponseBody) SetMessage(v string) *QuerySystemEventDemoResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySystemEventDemoResponseBody) SetRequestId(v string) *QuerySystemEventDemoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySystemEventDemoResponseBody) SetSuccess(v string) *QuerySystemEventDemoResponseBody {
	s.Success = &v
	return s
}

type QuerySystemEventDemoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySystemEventDemoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySystemEventDemoResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySystemEventDemoResponse) GoString() string {
	return s.String()
}

func (s *QuerySystemEventDemoResponse) SetHeaders(v map[string]*string) *QuerySystemEventDemoResponse {
	s.Headers = v
	return s
}

func (s *QuerySystemEventDemoResponse) SetStatusCode(v int32) *QuerySystemEventDemoResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySystemEventDemoResponse) SetBody(v *QuerySystemEventDemoResponseBody) *QuerySystemEventDemoResponse {
	s.Body = v
	return s
}

type QueryTaskConfigRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryTaskConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskConfigRequest) SetRegionId(v string) *QueryTaskConfigRequest {
	s.RegionId = &v
	return s
}

type QueryTaskConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTaskConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskConfigResponseBody) SetCode(v string) *QueryTaskConfigResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTaskConfigResponseBody) SetData(v string) *QueryTaskConfigResponseBody {
	s.Data = &v
	return s
}

func (s *QueryTaskConfigResponseBody) SetMessage(v string) *QueryTaskConfigResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTaskConfigResponseBody) SetRequestId(v string) *QueryTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskConfigResponseBody) SetSuccess(v string) *QueryTaskConfigResponseBody {
	s.Success = &v
	return s
}

type QueryTaskConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTaskConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskConfigResponse) SetHeaders(v map[string]*string) *QueryTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskConfigResponse) SetStatusCode(v int32) *QueryTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskConfigResponse) SetBody(v *QueryTaskConfigResponseBody) *QueryTaskConfigResponse {
	s.Body = v
	return s
}

type QueryTaskMonitorDataRequest struct {
	Cursor    *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Length    *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	Period    *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
}

func (s QueryTaskMonitorDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskMonitorDataRequest) SetCursor(v string) *QueryTaskMonitorDataRequest {
	s.Cursor = &v
	return s
}

func (s *QueryTaskMonitorDataRequest) SetEndTime(v string) *QueryTaskMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *QueryTaskMonitorDataRequest) SetLength(v int32) *QueryTaskMonitorDataRequest {
	s.Length = &v
	return s
}

func (s *QueryTaskMonitorDataRequest) SetPeriod(v string) *QueryTaskMonitorDataRequest {
	s.Period = &v
	return s
}

func (s *QueryTaskMonitorDataRequest) SetRegionId(v string) *QueryTaskMonitorDataRequest {
	s.RegionId = &v
	return s
}

func (s *QueryTaskMonitorDataRequest) SetStartTime(v string) *QueryTaskMonitorDataRequest {
	s.StartTime = &v
	return s
}

func (s *QueryTaskMonitorDataRequest) SetTaskId(v string) *QueryTaskMonitorDataRequest {
	s.TaskId = &v
	return s
}

func (s *QueryTaskMonitorDataRequest) SetType(v string) *QueryTaskMonitorDataRequest {
	s.Type = &v
	return s
}

func (s *QueryTaskMonitorDataRequest) SetMetricName(v string) *QueryTaskMonitorDataRequest {
	s.MetricName = &v
	return s
}

type QueryTaskMonitorDataResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Cursor    *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s QueryTaskMonitorDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskMonitorDataResponseBody) SetCode(v string) *QueryTaskMonitorDataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTaskMonitorDataResponseBody) SetCursor(v string) *QueryTaskMonitorDataResponseBody {
	s.Cursor = &v
	return s
}

func (s *QueryTaskMonitorDataResponseBody) SetData(v string) *QueryTaskMonitorDataResponseBody {
	s.Data = &v
	return s
}

func (s *QueryTaskMonitorDataResponseBody) SetMessage(v string) *QueryTaskMonitorDataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTaskMonitorDataResponseBody) SetRequestId(v string) *QueryTaskMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskMonitorDataResponseBody) SetSuccess(v string) *QueryTaskMonitorDataResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTaskMonitorDataResponseBody) SetTraceId(v string) *QueryTaskMonitorDataResponseBody {
	s.TraceId = &v
	return s
}

type QueryTaskMonitorDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTaskMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTaskMonitorDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskMonitorDataResponse) SetHeaders(v map[string]*string) *QueryTaskMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskMonitorDataResponse) SetStatusCode(v int32) *QueryTaskMonitorDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskMonitorDataResponse) SetBody(v *QueryTaskMonitorDataResponseBody) *QueryTaskMonitorDataResponse {
	s.Body = v
	return s
}

type TaskConfigListRequest struct {
	GroupId    *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskName   *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s TaskConfigListRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskConfigListRequest) GoString() string {
	return s.String()
}

func (s *TaskConfigListRequest) SetGroupId(v int64) *TaskConfigListRequest {
	s.GroupId = &v
	return s
}

func (s *TaskConfigListRequest) SetId(v int64) *TaskConfigListRequest {
	s.Id = &v
	return s
}

func (s *TaskConfigListRequest) SetPageNumber(v int32) *TaskConfigListRequest {
	s.PageNumber = &v
	return s
}

func (s *TaskConfigListRequest) SetPageSize(v int32) *TaskConfigListRequest {
	s.PageSize = &v
	return s
}

func (s *TaskConfigListRequest) SetRegionId(v string) *TaskConfigListRequest {
	s.RegionId = &v
	return s
}

func (s *TaskConfigListRequest) SetTaskName(v string) *TaskConfigListRequest {
	s.TaskName = &v
	return s
}

type TaskConfigListResponseBody struct {
	ErrorCode    *int32                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	PageNumber   *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageTotal    *int32                              `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskList     *TaskConfigListResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Struct"`
	Total        *int32                              `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s TaskConfigListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TaskConfigListResponseBody) GoString() string {
	return s.String()
}

func (s *TaskConfigListResponseBody) SetErrorCode(v int32) *TaskConfigListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TaskConfigListResponseBody) SetErrorMessage(v string) *TaskConfigListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TaskConfigListResponseBody) SetPageNumber(v int32) *TaskConfigListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *TaskConfigListResponseBody) SetPageSize(v int32) *TaskConfigListResponseBody {
	s.PageSize = &v
	return s
}

func (s *TaskConfigListResponseBody) SetPageTotal(v int32) *TaskConfigListResponseBody {
	s.PageTotal = &v
	return s
}

func (s *TaskConfigListResponseBody) SetRequestId(v string) *TaskConfigListResponseBody {
	s.RequestId = &v
	return s
}

func (s *TaskConfigListResponseBody) SetSuccess(v bool) *TaskConfigListResponseBody {
	s.Success = &v
	return s
}

func (s *TaskConfigListResponseBody) SetTaskList(v *TaskConfigListResponseBodyTaskList) *TaskConfigListResponseBody {
	s.TaskList = v
	return s
}

func (s *TaskConfigListResponseBody) SetTotal(v int32) *TaskConfigListResponseBody {
	s.Total = &v
	return s
}

type TaskConfigListResponseBodyTaskList struct {
	NodeTaskConfig []*TaskConfigListResponseBodyTaskListNodeTaskConfig `json:"NodeTaskConfig,omitempty" xml:"NodeTaskConfig,omitempty" type:"Repeated"`
}

func (s TaskConfigListResponseBodyTaskList) String() string {
	return tea.Prettify(s)
}

func (s TaskConfigListResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *TaskConfigListResponseBodyTaskList) SetNodeTaskConfig(v []*TaskConfigListResponseBodyTaskListNodeTaskConfig) *TaskConfigListResponseBodyTaskList {
	s.NodeTaskConfig = v
	return s
}

type TaskConfigListResponseBodyTaskListNodeTaskConfig struct {
	AlertConfig  *string                                                       `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty"`
	Disabled     *bool                                                         `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	GroupId      *int64                                                        `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName    *string                                                       `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Id           *int64                                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceList *TaskConfigListResponseBodyTaskListNodeTaskConfigInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
	JsonData     *string                                                       `json:"JsonData,omitempty" xml:"JsonData,omitempty"`
	TaskName     *string                                                       `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskScope    *string                                                       `json:"TaskScope,omitempty" xml:"TaskScope,omitempty"`
	TaskType     *string                                                       `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s TaskConfigListResponseBodyTaskListNodeTaskConfig) String() string {
	return tea.Prettify(s)
}

func (s TaskConfigListResponseBodyTaskListNodeTaskConfig) GoString() string {
	return s.String()
}

func (s *TaskConfigListResponseBodyTaskListNodeTaskConfig) SetAlertConfig(v string) *TaskConfigListResponseBodyTaskListNodeTaskConfig {
	s.AlertConfig = &v
	return s
}

func (s *TaskConfigListResponseBodyTaskListNodeTaskConfig) SetDisabled(v bool) *TaskConfigListResponseBodyTaskListNodeTaskConfig {
	s.Disabled = &v
	return s
}

func (s *TaskConfigListResponseBodyTaskListNodeTaskConfig) SetGroupId(v int64) *TaskConfigListResponseBodyTaskListNodeTaskConfig {
	s.GroupId = &v
	return s
}

func (s *TaskConfigListResponseBodyTaskListNodeTaskConfig) SetGroupName(v string) *TaskConfigListResponseBodyTaskListNodeTaskConfig {
	s.GroupName = &v
	return s
}

func (s *TaskConfigListResponseBodyTaskListNodeTaskConfig) SetId(v int64) *TaskConfigListResponseBodyTaskListNodeTaskConfig {
	s.Id = &v
	return s
}

func (s *TaskConfigListResponseBodyTaskListNodeTaskConfig) SetInstanceList(v *TaskConfigListResponseBodyTaskListNodeTaskConfigInstanceList) *TaskConfigListResponseBodyTaskListNodeTaskConfig {
	s.InstanceList = v
	return s
}

func (s *TaskConfigListResponseBodyTaskListNodeTaskConfig) SetJsonData(v string) *TaskConfigListResponseBodyTaskListNodeTaskConfig {
	s.JsonData = &v
	return s
}

func (s *TaskConfigListResponseBodyTaskListNodeTaskConfig) SetTaskName(v string) *TaskConfigListResponseBodyTaskListNodeTaskConfig {
	s.TaskName = &v
	return s
}

func (s *TaskConfigListResponseBodyTaskListNodeTaskConfig) SetTaskScope(v string) *TaskConfigListResponseBodyTaskListNodeTaskConfig {
	s.TaskScope = &v
	return s
}

func (s *TaskConfigListResponseBodyTaskListNodeTaskConfig) SetTaskType(v string) *TaskConfigListResponseBodyTaskListNodeTaskConfig {
	s.TaskType = &v
	return s
}

type TaskConfigListResponseBodyTaskListNodeTaskConfigInstanceList struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s TaskConfigListResponseBodyTaskListNodeTaskConfigInstanceList) String() string {
	return tea.Prettify(s)
}

func (s TaskConfigListResponseBodyTaskListNodeTaskConfigInstanceList) GoString() string {
	return s.String()
}

func (s *TaskConfigListResponseBodyTaskListNodeTaskConfigInstanceList) SetString_(v []*string) *TaskConfigListResponseBodyTaskListNodeTaskConfigInstanceList {
	s.String_ = v
	return s
}

type TaskConfigListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskConfigListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskConfigListResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskConfigListResponse) GoString() string {
	return s.String()
}

func (s *TaskConfigListResponse) SetHeaders(v map[string]*string) *TaskConfigListResponse {
	s.Headers = v
	return s
}

func (s *TaskConfigListResponse) SetStatusCode(v int32) *TaskConfigListResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskConfigListResponse) SetBody(v *TaskConfigListResponseBody) *TaskConfigListResponse {
	s.Body = v
	return s
}

type UpdateAlarmRequest struct {
	// example:
	//
	// >
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	ContactGroups      *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// 24
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 576fbae7-2fd1-411a-ae13-6f09f4fafdde
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test_modify
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	NotifyType *int32 `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// example:
	//
	// 60
	Period   *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// example:
	//
	// 0
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 40
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// example:
	//
	// http://www.net.cn/example_callback.htm
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s UpdateAlarmRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlarmRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlarmRequest) SetComparisonOperator(v string) *UpdateAlarmRequest {
	s.ComparisonOperator = &v
	return s
}

func (s *UpdateAlarmRequest) SetContactGroups(v string) *UpdateAlarmRequest {
	s.ContactGroups = &v
	return s
}

func (s *UpdateAlarmRequest) SetDryRun(v bool) *UpdateAlarmRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateAlarmRequest) SetEndTime(v int32) *UpdateAlarmRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateAlarmRequest) SetEvaluationCount(v int32) *UpdateAlarmRequest {
	s.EvaluationCount = &v
	return s
}

func (s *UpdateAlarmRequest) SetId(v string) *UpdateAlarmRequest {
	s.Id = &v
	return s
}

func (s *UpdateAlarmRequest) SetName(v string) *UpdateAlarmRequest {
	s.Name = &v
	return s
}

func (s *UpdateAlarmRequest) SetNotifyType(v int32) *UpdateAlarmRequest {
	s.NotifyType = &v
	return s
}

func (s *UpdateAlarmRequest) SetPeriod(v int32) *UpdateAlarmRequest {
	s.Period = &v
	return s
}

func (s *UpdateAlarmRequest) SetRegionId(v string) *UpdateAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAlarmRequest) SetSilenceTime(v int32) *UpdateAlarmRequest {
	s.SilenceTime = &v
	return s
}

func (s *UpdateAlarmRequest) SetStartTime(v int32) *UpdateAlarmRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateAlarmRequest) SetStatistics(v string) *UpdateAlarmRequest {
	s.Statistics = &v
	return s
}

func (s *UpdateAlarmRequest) SetThreshold(v string) *UpdateAlarmRequest {
	s.Threshold = &v
	return s
}

func (s *UpdateAlarmRequest) SetWebhook(v string) *UpdateAlarmRequest {
	s.Webhook = &v
	return s
}

type UpdateAlarmResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 945B9183-95C0-44FF-B30C-9ED37D44F6DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAlarmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlarmResponseBody) SetCode(v string) *UpdateAlarmResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAlarmResponseBody) SetMessage(v string) *UpdateAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAlarmResponseBody) SetRequestId(v string) *UpdateAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlarmResponseBody) SetSuccess(v bool) *UpdateAlarmResponseBody {
	s.Success = &v
	return s
}

type UpdateAlarmResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlarmResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlarmResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlarmResponse) SetHeaders(v map[string]*string) *UpdateAlarmResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlarmResponse) SetStatusCode(v int32) *UpdateAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlarmResponse) SetBody(v *UpdateAlarmResponseBody) *UpdateAlarmResponse {
	s.Body = v
	return s
}

type UpdateMonitoringTemplateRequest struct {
	// This parameter is required.
	AlertTemplatesJson       *string `json:"AlertTemplatesJson,omitempty" xml:"AlertTemplatesJson,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HostAvailabilityTemplate *string `json:"HostAvailabilityTemplate,omitempty" xml:"HostAvailabilityTemplate,omitempty"`
	// This parameter is required.
	Id                      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProcessMonitorTemplates *string `json:"ProcessMonitorTemplates,omitempty" xml:"ProcessMonitorTemplates,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	RestVersion          *int64  `json:"RestVersion,omitempty" xml:"RestVersion,omitempty"`
	SystemEventTemplates *string `json:"SystemEventTemplates,omitempty" xml:"SystemEventTemplates,omitempty"`
}

func (s UpdateMonitoringTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMonitoringTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateMonitoringTemplateRequest) SetAlertTemplatesJson(v string) *UpdateMonitoringTemplateRequest {
	s.AlertTemplatesJson = &v
	return s
}

func (s *UpdateMonitoringTemplateRequest) SetDescription(v string) *UpdateMonitoringTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateMonitoringTemplateRequest) SetHostAvailabilityTemplate(v string) *UpdateMonitoringTemplateRequest {
	s.HostAvailabilityTemplate = &v
	return s
}

func (s *UpdateMonitoringTemplateRequest) SetId(v int64) *UpdateMonitoringTemplateRequest {
	s.Id = &v
	return s
}

func (s *UpdateMonitoringTemplateRequest) SetName(v string) *UpdateMonitoringTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateMonitoringTemplateRequest) SetProcessMonitorTemplates(v string) *UpdateMonitoringTemplateRequest {
	s.ProcessMonitorTemplates = &v
	return s
}

func (s *UpdateMonitoringTemplateRequest) SetRegionId(v string) *UpdateMonitoringTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateMonitoringTemplateRequest) SetRestVersion(v int64) *UpdateMonitoringTemplateRequest {
	s.RestVersion = &v
	return s
}

func (s *UpdateMonitoringTemplateRequest) SetSystemEventTemplates(v string) *UpdateMonitoringTemplateRequest {
	s.SystemEventTemplates = &v
	return s
}

type UpdateMonitoringTemplateResponseBody struct {
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMonitoringTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMonitoringTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMonitoringTemplateResponseBody) SetErrorCode(v int32) *UpdateMonitoringTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateMonitoringTemplateResponseBody) SetErrorMessage(v string) *UpdateMonitoringTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateMonitoringTemplateResponseBody) SetRequestId(v string) *UpdateMonitoringTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMonitoringTemplateResponseBody) SetSuccess(v bool) *UpdateMonitoringTemplateResponseBody {
	s.Success = &v
	return s
}

type UpdateMonitoringTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMonitoringTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMonitoringTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMonitoringTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateMonitoringTemplateResponse) SetHeaders(v map[string]*string) *UpdateMonitoringTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateMonitoringTemplateResponse) SetStatusCode(v int32) *UpdateMonitoringTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMonitoringTemplateResponse) SetBody(v *UpdateMonitoringTemplateResponseBody) *UpdateMonitoringTemplateResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cms"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # AccessKeyGet
//
// @param request - AccessKeyGetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccessKeyGetResponse
func (client *Client) AccessKeyGetWithOptions(request *AccessKeyGetRequest, runtime *util.RuntimeOptions) (_result *AccessKeyGetResponse, _err error) {
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
		Action:      tea.String("AccessKeyGet"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AccessKeyGetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AccessKeyGet
//
// @param request - AccessKeyGetRequest
//
// @return AccessKeyGetResponse
func (client *Client) AccessKeyGet(request *AccessKeyGetRequest) (_result *AccessKeyGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AccessKeyGetResponse{}
	_body, _err := client.AccessKeyGetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddMyGroupInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMyGroupInstancesResponse
func (client *Client) AddMyGroupInstancesWithOptions(request *AddMyGroupInstancesRequest, runtime *util.RuntimeOptions) (_result *AddMyGroupInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Instances)) {
		query["Instances"] = request.Instances
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddMyGroupInstances"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddMyGroupInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddMyGroupInstancesRequest
//
// @return AddMyGroupInstancesResponse
func (client *Client) AddMyGroupInstances(request *AddMyGroupInstancesRequest) (_result *AddMyGroupInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddMyGroupInstancesResponse{}
	_body, _err := client.AddMyGroupInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # CreateAlarm
//
// @param request - CreateAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlarmResponse
func (client *Client) CreateAlarmWithOptions(request *CreateAlarmRequest, runtime *util.RuntimeOptions) (_result *CreateAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComparisonOperator)) {
		query["ComparisonOperator"] = request.ComparisonOperator
	}

	if !tea.BoolValue(util.IsUnset(request.ContactGroups)) {
		query["ContactGroups"] = request.ContactGroups
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EvaluationCount)) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyType)) {
		query["NotifyType"] = request.NotifyType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.SilenceTime)) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Statistics)) {
		query["Statistics"] = request.Statistics
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	if !tea.BoolValue(util.IsUnset(request.Webhook)) {
		query["Webhook"] = request.Webhook
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlarm"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateAlarm
//
// @param request - CreateAlarmRequest
//
// @return CreateAlarmResponse
func (client *Client) CreateAlarm(request *CreateAlarmRequest) (_result *CreateAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAlarmResponse{}
	_body, _err := client.CreateAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建双写配置
//
// @param request - CreateHybridDoubleWriteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHybridDoubleWriteResponse
func (client *Client) CreateHybridDoubleWriteWithOptions(request *CreateHybridDoubleWriteRequest, runtime *util.RuntimeOptions) (_result *CreateHybridDoubleWriteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.SourceNamespace)) {
		query["SourceNamespace"] = request.SourceNamespace
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUserId)) {
		query["SourceUserId"] = request.SourceUserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHybridDoubleWrite"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHybridDoubleWriteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建双写配置
//
// @param request - CreateHybridDoubleWriteRequest
//
// @return CreateHybridDoubleWriteResponse
func (client *Client) CreateHybridDoubleWrite(request *CreateHybridDoubleWriteRequest) (_result *CreateHybridDoubleWriteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHybridDoubleWriteResponse{}
	_body, _err := client.CreateHybridDoubleWriteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateMonitoringTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitoringTemplateResponse
func (client *Client) CreateMonitoringTemplateWithOptions(request *CreateMonitoringTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateMonitoringTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertTemplatesJson)) {
		query["AlertTemplatesJson"] = request.AlertTemplatesJson
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.HostAvailabilityTemplate)) {
		query["HostAvailabilityTemplate"] = request.HostAvailabilityTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessMonitorTemplates)) {
		query["ProcessMonitorTemplates"] = request.ProcessMonitorTemplates
	}

	if !tea.BoolValue(util.IsUnset(request.SystemEventTemplates)) {
		query["SystemEventTemplates"] = request.SystemEventTemplates
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMonitoringTemplate"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMonitoringTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMonitoringTemplateRequest
//
// @return CreateMonitoringTemplateResponse
func (client *Client) CreateMonitoringTemplate(request *CreateMonitoringTemplateRequest) (_result *CreateMonitoringTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMonitoringTemplateResponse{}
	_body, _err := client.CreateMonitoringTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateMyGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMyGroupsResponse
func (client *Client) CreateMyGroupsWithOptions(request *CreateMyGroupsRequest, runtime *util.RuntimeOptions) (_result *CreateMyGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindUrl)) {
		query["BindUrl"] = request.BindUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ContactGroups)) {
		query["ContactGroups"] = request.ContactGroups
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		query["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMyGroups"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMyGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMyGroupsRequest
//
// @return CreateMyGroupsResponse
func (client *Client) CreateMyGroups(request *CreateMyGroupsRequest) (_result *CreateMyGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMyGroupsResponse{}
	_body, _err := client.CreateMyGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # CreateTask
//
// @param request - CreateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskResponse
func (client *Client) CreateTaskWithOptions(request *CreateTaskRequest, runtime *util.RuntimeOptions) (_result *CreateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.AlertIds)) {
		query["AlertIds"] = request.AlertIds
	}

	if !tea.BoolValue(util.IsUnset(request.AlertRule)) {
		query["AlertRule"] = request.AlertRule
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IntervalUnit)) {
		query["IntervalUnit"] = request.IntervalUnit
	}

	if !tea.BoolValue(util.IsUnset(request.IspCity)) {
		query["IspCity"] = request.IspCity
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		query["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["caller"] = request.Caller
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTask"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateTask
//
// @param request - CreateTaskRequest
//
// @return CreateTaskResponse
func (client *Client) CreateTask(request *CreateTaskRequest) (_result *CreateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTaskResponse{}
	_body, _err := client.CreateTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DeleteAlarm
//
// @param request - DeleteAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlarmResponse
func (client *Client) DeleteAlarmWithOptions(request *DeleteAlarmRequest, runtime *util.RuntimeOptions) (_result *DeleteAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlarm"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteAlarm
//
// @param request - DeleteAlarmRequest
//
// @return DeleteAlarmResponse
func (client *Client) DeleteAlarm(request *DeleteAlarmRequest) (_result *DeleteAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlarmResponse{}
	_body, _err := client.DeleteAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DeleteCustomMetric
//
// @param request - DeleteCustomMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomMetricResponse
func (client *Client) DeleteCustomMetricWithOptions(request *DeleteCustomMetricRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Md5)) {
		query["Md5"] = request.Md5
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.UUID)) {
		query["UUID"] = request.UUID
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomMetric"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteCustomMetric
//
// @param request - DeleteCustomMetricRequest
//
// @return DeleteCustomMetricResponse
func (client *Client) DeleteCustomMetric(request *DeleteCustomMetricRequest) (_result *DeleteCustomMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomMetricResponse{}
	_body, _err := client.DeleteCustomMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除双写配置
//
// @param request - DeleteHybridDoubleWriteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHybridDoubleWriteResponse
func (client *Client) DeleteHybridDoubleWriteWithOptions(request *DeleteHybridDoubleWriteRequest, runtime *util.RuntimeOptions) (_result *DeleteHybridDoubleWriteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceNamespace)) {
		query["SourceNamespace"] = request.SourceNamespace
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUserId)) {
		query["SourceUserId"] = request.SourceUserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHybridDoubleWrite"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHybridDoubleWriteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除双写配置
//
// @param request - DeleteHybridDoubleWriteRequest
//
// @return DeleteHybridDoubleWriteResponse
func (client *Client) DeleteHybridDoubleWrite(request *DeleteHybridDoubleWriteRequest) (_result *DeleteHybridDoubleWriteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHybridDoubleWriteResponse{}
	_body, _err := client.DeleteHybridDoubleWriteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DeleteMetricRuleTargets
//
// @param request - DeleteMetricRuleTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetricRuleTargetsResponse
func (client *Client) DeleteMetricRuleTargetsWithOptions(request *DeleteMetricRuleTargetsRequest, runtime *util.RuntimeOptions) (_result *DeleteMetricRuleTargetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetIds)) {
		query["TargetIds"] = request.TargetIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMetricRuleTargets"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMetricRuleTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteMetricRuleTargets
//
// @param request - DeleteMetricRuleTargetsRequest
//
// @return DeleteMetricRuleTargetsResponse
func (client *Client) DeleteMetricRuleTargets(request *DeleteMetricRuleTargetsRequest) (_result *DeleteMetricRuleTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMetricRuleTargetsResponse{}
	_body, _err := client.DeleteMetricRuleTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DeleteMetricRules
//
// @param request - DeleteMetricRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetricRulesResponse
func (client *Client) DeleteMetricRulesWithOptions(request *DeleteMetricRulesRequest, runtime *util.RuntimeOptions) (_result *DeleteMetricRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMetricRules"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMetricRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteMetricRules
//
// @param request - DeleteMetricRulesRequest
//
// @return DeleteMetricRulesResponse
func (client *Client) DeleteMetricRules(request *DeleteMetricRulesRequest) (_result *DeleteMetricRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMetricRulesResponse{}
	_body, _err := client.DeleteMetricRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// deletemygroupinstances
//
// @param request - DeleteMyGroupInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMyGroupInstancesResponse
func (client *Client) DeleteMyGroupInstancesWithOptions(request *DeleteMyGroupInstancesRequest, runtime *util.RuntimeOptions) (_result *DeleteMyGroupInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMyGroupInstances"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMyGroupInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// deletemygroupinstances
//
// @param request - DeleteMyGroupInstancesRequest
//
// @return DeleteMyGroupInstancesResponse
func (client *Client) DeleteMyGroupInstances(request *DeleteMyGroupInstancesRequest) (_result *DeleteMyGroupInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMyGroupInstancesResponse{}
	_body, _err := client.DeleteMyGroupInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteMyGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMyGroupsResponse
func (client *Client) DeleteMyGroupsWithOptions(request *DeleteMyGroupsRequest, runtime *util.RuntimeOptions) (_result *DeleteMyGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMyGroups"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMyGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteMyGroupsRequest
//
// @return DeleteMyGroupsResponse
func (client *Client) DeleteMyGroups(request *DeleteMyGroupsRequest) (_result *DeleteMyGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMyGroupsResponse{}
	_body, _err := client.DeleteMyGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTasksResponse
func (client *Client) DeleteTasksWithOptions(request *DeleteTasksRequest, runtime *util.RuntimeOptions) (_result *DeleteTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsDeleteAlarms)) {
		query["IsDeleteAlarms"] = request.IsDeleteAlarms
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIds)) {
		query["TaskIds"] = request.TaskIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTasks"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteTasksRequest
//
// @return DeleteTasksResponse
func (client *Client) DeleteTasks(request *DeleteTasksRequest) (_result *DeleteTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTasksResponse{}
	_body, _err := client.DeleteTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DescribeAlarmHistory
//
// @param request - DescribeAlarmHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlarmHistoryResponse
func (client *Client) DescribeAlarmHistoryWithOptions(request *DescribeAlarmHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeAlarmHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertName)) {
		query["AlertName"] = request.AlertName
	}

	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.OnlyCount)) {
		query["OnlyCount"] = request.OnlyCount
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlarmHistory"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlarmHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeAlarmHistory
//
// @param request - DescribeAlarmHistoryRequest
//
// @return DescribeAlarmHistoryResponse
func (client *Client) DescribeAlarmHistory(request *DescribeAlarmHistoryRequest) (_result *DescribeAlarmHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlarmHistoryResponse{}
	_body, _err := client.DescribeAlarmHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DescribeAlarms
//
// @param request - DescribeAlarmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlarmsResponse
func (client *Client) DescribeAlarmsWithOptions(request *DescribeAlarmsRequest, runtime *util.RuntimeOptions) (_result *DescribeAlarmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertState)) {
		query["AlertState"] = request.AlertState
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableState)) {
		query["EnableState"] = request.EnableState
	}

	if !tea.BoolValue(util.IsUnset(request.GroupBy)) {
		query["GroupBy"] = request.GroupBy
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.NameKeyword)) {
		query["NameKeyword"] = request.NameKeyword
	}

	if !tea.BoolValue(util.IsUnset(request.Names)) {
		query["Names"] = request.Names
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlarms"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlarmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeAlarms
//
// @param request - DescribeAlarmsRequest
//
// @return DescribeAlarmsResponse
func (client *Client) DescribeAlarms(request *DescribeAlarmsRequest) (_result *DescribeAlarmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlarmsResponse{}
	_body, _err := client.DescribeAlarmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// describealarmsforresources
//
// @param request - DescribeAlarmsForResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlarmsForResourcesResponse
func (client *Client) DescribeAlarmsForResourcesWithOptions(request *DescribeAlarmsForResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeAlarmsForResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertState)) {
		query["AlertState"] = request.AlertState
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.EnableState)) {
		query["EnableState"] = request.EnableState
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlarmsForResources"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlarmsForResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// describealarmsforresources
//
// @param request - DescribeAlarmsForResourcesRequest
//
// @return DescribeAlarmsForResourcesResponse
func (client *Client) DescribeAlarmsForResources(request *DescribeAlarmsForResourcesRequest) (_result *DescribeAlarmsForResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlarmsForResourcesResponse{}
	_body, _err := client.DescribeAlarmsForResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAlertHistoryListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertHistoryListResponse
func (client *Client) DescribeAlertHistoryListWithOptions(request *DescribeAlertHistoryListRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertHistoryListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertName)) {
		query["AlertName"] = request.AlertName
	}

	if !tea.BoolValue(util.IsUnset(request.Ascending)) {
		query["Ascending"] = request.Ascending
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.OnlyCount)) {
		query["OnlyCount"] = request.OnlyCount
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlertHistoryList"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlertHistoryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAlertHistoryListRequest
//
// @return DescribeAlertHistoryListResponse
func (client *Client) DescribeAlertHistoryList(request *DescribeAlertHistoryListRequest) (_result *DescribeAlertHistoryListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertHistoryListResponse{}
	_body, _err := client.DescribeAlertHistoryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContactResponse
func (client *Client) DescribeContactWithOptions(request *DescribeContactRequest, runtime *util.RuntimeOptions) (_result *DescribeContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContact"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeContactRequest
//
// @return DescribeContactResponse
func (client *Client) DescribeContact(request *DescribeContactRequest) (_result *DescribeContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContactResponse{}
	_body, _err := client.DescribeContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询本数据源被双写到哪里
//
// @param request - DescribeHybridDoubleWriteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridDoubleWriteResponse
func (client *Client) DescribeHybridDoubleWriteWithOptions(request *DescribeHybridDoubleWriteRequest, runtime *util.RuntimeOptions) (_result *DescribeHybridDoubleWriteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceNamespace)) {
		query["SourceNamespace"] = request.SourceNamespace
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUserId)) {
		query["SourceUserId"] = request.SourceUserId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetNamespace)) {
		query["TargetNamespace"] = request.TargetNamespace
	}

	if !tea.BoolValue(util.IsUnset(request.TargetUserId)) {
		query["TargetUserId"] = request.TargetUserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHybridDoubleWrite"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHybridDoubleWriteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询本数据源被双写到哪里
//
// @param request - DescribeHybridDoubleWriteRequest
//
// @return DescribeHybridDoubleWriteResponse
func (client *Client) DescribeHybridDoubleWrite(request *DescribeHybridDoubleWriteRequest) (_result *DescribeHybridDoubleWriteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHybridDoubleWriteResponse{}
	_body, _err := client.DescribeHybridDoubleWriteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取探针列表
//
// @param request - DescribeISPAreaCityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeISPAreaCityResponse
func (client *Client) DescribeISPAreaCityWithOptions(request *DescribeISPAreaCityRequest, runtime *util.RuntimeOptions) (_result *DescribeISPAreaCityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.Isp)) {
		query["Isp"] = request.Isp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeISPAreaCity"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeISPAreaCityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取探针列表
//
// @param request - DescribeISPAreaCityRequest
//
// @return DescribeISPAreaCityResponse
func (client *Client) DescribeISPAreaCity(request *DescribeISPAreaCityRequest) (_result *DescribeISPAreaCityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeISPAreaCityResponse{}
	_body, _err := client.DescribeISPAreaCityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeMetricRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricRuleListResponse
func (client *Client) DescribeMetricRuleListWithOptions(request *DescribeMetricRuleListRequest, runtime *util.RuntimeOptions) (_result *DescribeMetricRuleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertState)) {
		query["AlertState"] = request.AlertState
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.EnableState)) {
		query["EnableState"] = request.EnableState
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RuleIds)) {
		query["RuleIds"] = request.RuleIds
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMetricRuleList"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMetricRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeMetricRuleListRequest
//
// @return DescribeMetricRuleListResponse
func (client *Client) DescribeMetricRuleList(request *DescribeMetricRuleListRequest) (_result *DescribeMetricRuleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetricRuleListResponse{}
	_body, _err := client.DescribeMetricRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTaskDetailResponse
func (client *Client) DescribeTaskDetailWithOptions(request *DescribeTaskDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeTaskDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTaskDetail"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTaskDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeTaskDetailRequest
//
// @return DescribeTaskDetailResponse
func (client *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (_result *DescribeTaskDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTaskDetailResponse{}
	_body, _err := client.DescribeTaskDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DescribeTasks
//
// @param request - DescribeTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTasksResponse
func (client *Client) DescribeTasksWithOptions(request *DescribeTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTasks"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeTasks
//
// @param request - DescribeTasksRequest
//
// @return DescribeTasksResponse
func (client *Client) DescribeTasks(request *DescribeTasksRequest) (_result *DescribeTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTasksResponse{}
	_body, _err := client.DescribeTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DisableAlarm
//
// @param request - DisableAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAlarmResponse
func (client *Client) DisableAlarmWithOptions(request *DisableAlarmRequest, runtime *util.RuntimeOptions) (_result *DisableAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableAlarm"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DisableAlarm
//
// @param request - DisableAlarmRequest
//
// @return DisableAlarmResponse
func (client *Client) DisableAlarm(request *DisableAlarmRequest) (_result *DisableAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableAlarmResponse{}
	_body, _err := client.DisableAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # EnableAlarm
//
// @param request - EnableAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAlarmResponse
func (client *Client) EnableAlarmWithOptions(request *EnableAlarmRequest, runtime *util.RuntimeOptions) (_result *EnableAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableAlarm"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # EnableAlarm
//
// @param request - EnableAlarmRequest
//
// @return EnableAlarmResponse
func (client *Client) EnableAlarm(request *EnableAlarmRequest) (_result *EnableAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableAlarmResponse{}
	_body, _err := client.EnableAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContactsResponse
func (client *Client) GetContactsWithOptions(request *GetContactsRequest, runtime *util.RuntimeOptions) (_result *GetContactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetContacts"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetContactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetContactsRequest
//
// @return GetContactsResponse
func (client *Client) GetContacts(request *GetContactsRequest) (_result *GetContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetContactsResponse{}
	_body, _err := client.GetContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取行切分结果
//
// @param request - GetLineSplitResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLineSplitResultResponse
func (client *Client) GetLineSplitResultWithOptions(request *GetLineSplitResultRequest, runtime *util.RuntimeOptions) (_result *GetLineSplitResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Line)) {
		query["Line"] = request.Line
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["Prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.Regex)) {
		query["Regex"] = request.Regex
	}

	if !tea.BoolValue(util.IsUnset(request.SelectContent)) {
		query["SelectContent"] = request.SelectContent
	}

	if !tea.BoolValue(util.IsUnset(request.SplitType)) {
		query["SplitType"] = request.SplitType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLineSplitResult"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLineSplitResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取行切分结果
//
// @param request - GetLineSplitResultRequest
//
// @return GetLineSplitResultResponse
func (client *Client) GetLineSplitResult(request *GetLineSplitResultRequest) (_result *GetLineSplitResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLineSplitResultResponse{}
	_body, _err := client.GetLineSplitResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取日期提取结果的翻译
//
// @param request - GetLogColumnTranslateResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLogColumnTranslateResultResponse
func (client *Client) GetLogColumnTranslateResultWithOptions(request *GetLogColumnTranslateResultRequest, runtime *util.RuntimeOptions) (_result *GetLogColumnTranslateResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ColumnValue)) {
		query["ColumnValue"] = request.ColumnValue
	}

	if !tea.BoolValue(util.IsUnset(request.TranslateConfig)) {
		query["TranslateConfig"] = request.TranslateConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLogColumnTranslateResult"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLogColumnTranslateResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日期提取结果的翻译
//
// @param request - GetLogColumnTranslateResultRequest
//
// @return GetLogColumnTranslateResultResponse
func (client *Client) GetLogColumnTranslateResult(request *GetLogColumnTranslateResultRequest) (_result *GetLogColumnTranslateResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLogColumnTranslateResultResponse{}
	_body, _err := client.GetLogColumnTranslateResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetMonitoringTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMonitoringTemplateResponse
func (client *Client) GetMonitoringTemplateWithOptions(request *GetMonitoringTemplateRequest, runtime *util.RuntimeOptions) (_result *GetMonitoringTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMonitoringTemplate"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMonitoringTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetMonitoringTemplateRequest
//
// @return GetMonitoringTemplateResponse
func (client *Client) GetMonitoringTemplate(request *GetMonitoringTemplateRequest) (_result *GetMonitoringTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMonitoringTemplateResponse{}
	_body, _err := client.GetMonitoringTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetMyGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMyGroupsResponse
func (client *Client) GetMyGroupsWithOptions(request *GetMyGroupsRequest, runtime *util.RuntimeOptions) (_result *GetMyGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindUrl)) {
		query["BindUrl"] = request.BindUrl
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SelectContactGroups)) {
		query["SelectContactGroups"] = request.SelectContactGroups
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMyGroups"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMyGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetMyGroupsRequest
//
// @return GetMyGroupsResponse
func (client *Client) GetMyGroups(request *GetMyGroupsRequest) (_result *GetMyGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMyGroupsResponse{}
	_body, _err := client.GetMyGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ListAlarm
//
// @param request - ListAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlarmResponse
func (client *Client) ListAlarmWithOptions(request *ListAlarmRequest, runtime *util.RuntimeOptions) (_result *ListAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dimension)) {
		query["Dimension"] = request.Dimension
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.IsEnable)) {
		query["IsEnable"] = request.IsEnable
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlarm"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListAlarm
//
// @param request - ListAlarmRequest
//
// @return ListAlarmResponse
func (client *Client) ListAlarm(request *ListAlarmRequest) (_result *ListAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAlarmResponse{}
	_body, _err := client.ListAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ListAlarmHistory
//
// @param request - ListAlarmHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlarmHistoryResponse
func (client *Client) ListAlarmHistoryWithOptions(request *ListAlarmHistoryRequest, runtime *util.RuntimeOptions) (_result *ListAlarmHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["Cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlarmHistory"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlarmHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListAlarmHistory
//
// @param request - ListAlarmHistoryRequest
//
// @return ListAlarmHistoryResponse
func (client *Client) ListAlarmHistory(request *ListAlarmHistoryRequest) (_result *ListAlarmHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAlarmHistoryResponse{}
	_body, _err := client.ListAlarmHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListContactGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContactGroupResponse
func (client *Client) ListContactGroupWithOptions(request *ListContactGroupRequest, runtime *util.RuntimeOptions) (_result *ListContactGroupResponse, _err error) {
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
		Action:      tea.String("ListContactGroup"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListContactGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListContactGroupRequest
//
// @return ListContactGroupResponse
func (client *Client) ListContactGroup(request *ListContactGroupRequest) (_result *ListContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListContactGroupResponse{}
	_body, _err := client.ListContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListEventRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEventRulesResponse
func (client *Client) ListEventRulesWithOptions(request *ListEventRulesRequest, runtime *util.RuntimeOptions) (_result *ListEventRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.NamePrefix)) {
		query["NamePrefix"] = request.NamePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEventRules"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEventRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListEventRulesRequest
//
// @return ListEventRulesResponse
func (client *Client) ListEventRules(request *ListEventRulesRequest) (_result *ListEventRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEventRulesResponse{}
	_body, _err := client.ListEventRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListMyGroupInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMyGroupInstancesResponse
func (client *Client) ListMyGroupInstancesWithOptions(request *ListMyGroupInstancesRequest, runtime *util.RuntimeOptions) (_result *ListMyGroupInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Total)) {
		query["Total"] = request.Total
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMyGroupInstances"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMyGroupInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMyGroupInstancesRequest
//
// @return ListMyGroupInstancesResponse
func (client *Client) ListMyGroupInstances(request *ListMyGroupInstancesRequest) (_result *ListMyGroupInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMyGroupInstancesResponse{}
	_body, _err := client.ListMyGroupInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListMyGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMyGroupsResponse
func (client *Client) ListMyGroupsWithOptions(request *ListMyGroupsRequest, runtime *util.RuntimeOptions) (_result *ListMyGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindUrls)) {
		query["BindUrls"] = request.BindUrls
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SelectContactGroups)) {
		query["SelectContactGroups"] = request.SelectContactGroups
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMyGroups"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMyGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMyGroupsRequest
//
// @return ListMyGroupsResponse
func (client *Client) ListMyGroups(request *ListMyGroupsRequest) (_result *ListMyGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMyGroupsResponse{}
	_body, _err := client.ListMyGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTaskResponse
func (client *Client) ModifyTaskWithOptions(request *ModifyTaskRequest, runtime *util.RuntimeOptions) (_result *ModifyTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.AlertIds)) {
		query["AlertIds"] = request.AlertIds
	}

	if !tea.BoolValue(util.IsUnset(request.AlertRule)) {
		query["AlertRule"] = request.AlertRule
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IntervalUnit)) {
		query["IntervalUnit"] = request.IntervalUnit
	}

	if !tea.BoolValue(util.IsUnset(request.IspCity)) {
		query["IspCity"] = request.IspCity
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		query["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["caller"] = request.Caller
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTask"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyTaskRequest
//
// @return ModifyTaskResponse
func (client *Client) ModifyTask(request *ModifyTaskRequest) (_result *ModifyTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTaskResponse{}
	_body, _err := client.ModifyTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # NodeList
//
// @param request - NodeListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NodeListResponse
func (client *Client) NodeListWithOptions(request *NodeListRequest, runtime *util.RuntimeOptions) (_result *NodeListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		query["HostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceRegionId)) {
		query["InstanceRegionId"] = request.InstanceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumbers)) {
		query["SerialNumbers"] = request.SerialNumbers
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("NodeList"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &NodeListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # NodeList
//
// @param request - NodeListRequest
//
// @return NodeListResponse
func (client *Client) NodeList(request *NodeListRequest) (_result *NodeListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &NodeListResponse{}
	_body, _err := client.NodeListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # NodeProcessCreate
//
// @param request - NodeProcessCreateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NodeProcessCreateResponse
func (client *Client) NodeProcessCreateWithOptions(request *NodeProcessCreateRequest, runtime *util.RuntimeOptions) (_result *NodeProcessCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Command)) {
		query["Command"] = request.Command
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessName)) {
		query["ProcessName"] = request.ProcessName
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessUser)) {
		query["ProcessUser"] = request.ProcessUser
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("NodeProcessCreate"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &NodeProcessCreateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # NodeProcessCreate
//
// @param request - NodeProcessCreateRequest
//
// @return NodeProcessCreateResponse
func (client *Client) NodeProcessCreate(request *NodeProcessCreateRequest) (_result *NodeProcessCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &NodeProcessCreateResponse{}
	_body, _err := client.NodeProcessCreateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # NodeProcesses
//
// @param request - NodeProcessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NodeProcessesResponse
func (client *Client) NodeProcessesWithOptions(request *NodeProcessesRequest, runtime *util.RuntimeOptions) (_result *NodeProcessesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("NodeProcesses"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &NodeProcessesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # NodeProcesses
//
// @param request - NodeProcessesRequest
//
// @return NodeProcessesResponse
func (client *Client) NodeProcesses(request *NodeProcessesRequest) (_result *NodeProcessesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &NodeProcessesResponse{}
	_body, _err := client.NodeProcessesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # NodeStatusList
//
// @param request - NodeStatusListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NodeStatusListResponse
func (client *Client) NodeStatusListWithOptions(request *NodeStatusListRequest, runtime *util.RuntimeOptions) (_result *NodeStatusListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("NodeStatusList"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &NodeStatusListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # NodeStatusList
//
// @param request - NodeStatusListRequest
//
// @return NodeStatusListResponse
func (client *Client) NodeStatusList(request *NodeStatusListRequest) (_result *NodeStatusListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &NodeStatusListResponse{}
	_body, _err := client.NodeStatusListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # NodeUninstall
//
// @param request - NodeUninstallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NodeUninstallResponse
func (client *Client) NodeUninstallWithOptions(request *NodeUninstallRequest, runtime *util.RuntimeOptions) (_result *NodeUninstallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("NodeUninstall"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &NodeUninstallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # NodeUninstall
//
// @param request - NodeUninstallRequest
//
// @return NodeUninstallResponse
func (client *Client) NodeUninstall(request *NodeUninstallRequest) (_result *NodeUninstallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &NodeUninstallResponse{}
	_body, _err := client.NodeUninstallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # PutCustomMetric
//
// @param request - PutCustomMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutCustomMetricResponse
func (client *Client) PutCustomMetricWithOptions(request *PutCustomMetricRequest, runtime *util.RuntimeOptions) (_result *PutCustomMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MetricList)) {
		query["MetricList"] = request.MetricList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutCustomMetric"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutCustomMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # PutCustomMetric
//
// @param request - PutCustomMetricRequest
//
// @return PutCustomMetricResponse
func (client *Client) PutCustomMetric(request *PutCustomMetricRequest) (_result *PutCustomMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutCustomMetricResponse{}
	_body, _err := client.PutCustomMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # PutEvent
//
// @param request - PutEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutEventResponse
func (client *Client) PutEventWithOptions(request *PutEventRequest, runtime *util.RuntimeOptions) (_result *PutEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventInfo)) {
		query["EventInfo"] = request.EventInfo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutEvent"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # PutEvent
//
// @param request - PutEventRequest
//
// @return PutEventResponse
func (client *Client) PutEvent(request *PutEventRequest) (_result *PutEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutEventResponse{}
	_body, _err := client.PutEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PutEventRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutEventRuleResponse
func (client *Client) PutEventRuleWithOptions(request *PutEventRuleRequest, runtime *util.RuntimeOptions) (_result *PutEventRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EventPattern)) {
		query["EventPattern"] = request.EventPattern
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		query["EventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutEventRule"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutEventRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - PutEventRuleRequest
//
// @return PutEventRuleResponse
func (client *Client) PutEventRule(request *PutEventRuleRequest) (_result *PutEventRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutEventRuleResponse{}
	_body, _err := client.PutEventRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PutEventTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutEventTargetsResponse
func (client *Client) PutEventTargetsWithOptions(request *PutEventTargetsRequest, runtime *util.RuntimeOptions) (_result *PutEventTargetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactParameters)) {
		query["ContactParameters"] = request.ContactParameters
	}

	if !tea.BoolValue(util.IsUnset(request.FcParameters)) {
		query["FcParameters"] = request.FcParameters
	}

	if !tea.BoolValue(util.IsUnset(request.MnsParameters)) {
		query["MnsParameters"] = request.MnsParameters
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.SlsParameters)) {
		query["SlsParameters"] = request.SlsParameters
	}

	if !tea.BoolValue(util.IsUnset(request.WebhookParameters)) {
		query["WebhookParameters"] = request.WebhookParameters
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutEventTargets"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutEventTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - PutEventTargetsRequest
//
// @return PutEventTargetsResponse
func (client *Client) PutEventTargets(request *PutEventTargetsRequest) (_result *PutEventTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutEventTargetsResponse{}
	_body, _err := client.PutEventTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # PutMetricRuleTargets
//
// @param request - PutMetricRuleTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutMetricRuleTargetsResponse
func (client *Client) PutMetricRuleTargetsWithOptions(request *PutMetricRuleTargetsRequest, runtime *util.RuntimeOptions) (_result *PutMetricRuleTargetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Targets)) {
		query["Targets"] = request.Targets
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Actions)) {
		body["Actions"] = request.Actions
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutMetricRuleTargets"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutMetricRuleTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # PutMetricRuleTargets
//
// @param request - PutMetricRuleTargetsRequest
//
// @return PutMetricRuleTargetsResponse
func (client *Client) PutMetricRuleTargets(request *PutMetricRuleTargetsRequest) (_result *PutMetricRuleTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutMetricRuleTargetsResponse{}
	_body, _err := client.PutMetricRuleTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # PutResourceMetricRule
//
// @param request - PutResourceMetricRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutResourceMetricRuleResponse
func (client *Client) PutResourceMetricRuleWithOptions(request *PutResourceMetricRuleRequest, runtime *util.RuntimeOptions) (_result *PutResourceMetricRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroups)) {
		query["ContactGroups"] = request.ContactGroups
	}

	if !tea.BoolValue(util.IsUnset(request.EffectiveInterval)) {
		query["EffectiveInterval"] = request.EffectiveInterval
	}

	if !tea.BoolValue(util.IsUnset(request.EmailSubject)) {
		query["EmailSubject"] = request.EmailSubject
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NoEffectiveInterval)) {
		query["NoEffectiveInterval"] = request.NoEffectiveInterval
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		query["Resources"] = request.Resources
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.SilenceTime)) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !tea.BoolValue(util.IsUnset(request.Webhook)) {
		query["Webhook"] = request.Webhook
	}

	if !tea.BoolValue(util.IsUnset(request.Escalations)) {
		query["Escalations"] = request.Escalations
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutResourceMetricRule"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutResourceMetricRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # PutResourceMetricRule
//
// @param request - PutResourceMetricRuleRequest
//
// @return PutResourceMetricRuleResponse
func (client *Client) PutResourceMetricRule(request *PutResourceMetricRuleRequest) (_result *PutResourceMetricRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutResourceMetricRuleResponse{}
	_body, _err := client.PutResourceMetricRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryCustomMetricListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomMetricListResponse
func (client *Client) QueryCustomMetricListWithOptions(request *QueryCustomMetricListRequest, runtime *util.RuntimeOptions) (_result *QueryCustomMetricListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dimension)) {
		query["Dimension"] = request.Dimension
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Md5)) {
		query["Md5"] = request.Md5
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCustomMetricList"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCustomMetricListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryCustomMetricListRequest
//
// @return QueryCustomMetricListResponse
func (client *Client) QueryCustomMetricList(request *QueryCustomMetricListRequest) (_result *QueryCustomMetricListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCustomMetricListResponse{}
	_body, _err := client.QueryCustomMetricListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # QueryMetricData
//
// @param request - QueryMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMetricDataResponse
func (client *Client) QueryMetricDataWithOptions(request *QueryMetricDataRequest, runtime *util.RuntimeOptions) (_result *QueryMetricDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Express)) {
		query["Express"] = request.Express
	}

	if !tea.BoolValue(util.IsUnset(request.Length)) {
		query["Length"] = request.Length
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		query["Metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMetricData"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMetricDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # QueryMetricData
//
// @param request - QueryMetricDataRequest
//
// @return QueryMetricDataResponse
func (client *Client) QueryMetricData(request *QueryMetricDataRequest) (_result *QueryMetricDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMetricDataResponse{}
	_body, _err := client.QueryMetricDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryMetricLastRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMetricLastResponse
func (client *Client) QueryMetricLastWithOptions(request *QueryMetricLastRequest, runtime *util.RuntimeOptions) (_result *QueryMetricLastResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["Cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Express)) {
		query["Express"] = request.Express
	}

	if !tea.BoolValue(util.IsUnset(request.Length)) {
		query["Length"] = request.Length
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		query["Metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMetricLast"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMetricLastResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryMetricLastRequest
//
// @return QueryMetricLastResponse
func (client *Client) QueryMetricLast(request *QueryMetricLastRequest) (_result *QueryMetricLastResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMetricLastResponse{}
	_body, _err := client.QueryMetricLastWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryMetricListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMetricListResponse
func (client *Client) QueryMetricListWithOptions(request *QueryMetricListRequest, runtime *util.RuntimeOptions) (_result *QueryMetricListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["Cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Express)) {
		query["Express"] = request.Express
	}

	if !tea.BoolValue(util.IsUnset(request.Length)) {
		query["Length"] = request.Length
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		query["Metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMetricList"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMetricListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryMetricListRequest
//
// @return QueryMetricListResponse
func (client *Client) QueryMetricList(request *QueryMetricListRequest) (_result *QueryMetricListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMetricListResponse{}
	_body, _err := client.QueryMetricListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询云监控开放的监控项详情
//
// @param request - QueryMetricMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMetricMetaResponse
func (client *Client) QueryMetricMetaWithOptions(request *QueryMetricMetaRequest, runtime *util.RuntimeOptions) (_result *QueryMetricMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		query["Metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMetricMeta"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMetricMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询云监控开放的监控项详情
//
// @param request - QueryMetricMetaRequest
//
// @return QueryMetricMetaResponse
func (client *Client) QueryMetricMeta(request *QueryMetricMetaRequest) (_result *QueryMetricMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMetricMetaResponse{}
	_body, _err := client.QueryMetricMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # QueryMetricTop
//
// @param request - QueryMetricTopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMetricTopResponse
func (client *Client) QueryMetricTopWithOptions(request *QueryMetricTopRequest, runtime *util.RuntimeOptions) (_result *QueryMetricTopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Express)) {
		query["Express"] = request.Express
	}

	if !tea.BoolValue(util.IsUnset(request.Length)) {
		query["Length"] = request.Length
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		query["Metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.OrderDesc)) {
		query["OrderDesc"] = request.OrderDesc
	}

	if !tea.BoolValue(util.IsUnset(request.Orderby)) {
		query["Orderby"] = request.Orderby
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.Project)) {
		query["Project"] = request.Project
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMetricTop"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMetricTopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # QueryMetricTop
//
// @param request - QueryMetricTopRequest
//
// @return QueryMetricTopResponse
func (client *Client) QueryMetricTop(request *QueryMetricTopRequest) (_result *QueryMetricTopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMetricTopResponse{}
	_body, _err := client.QueryMetricTopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询云监控支持的时序类监控项产品列表
//
// @param request - QueryProjectMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProjectMetaResponse
func (client *Client) QueryProjectMetaWithOptions(request *QueryProjectMetaRequest, runtime *util.RuntimeOptions) (_result *QueryProjectMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

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
		Action:      tea.String("QueryProjectMeta"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryProjectMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询云监控支持的时序类监控项产品列表
//
// @param request - QueryProjectMetaRequest
//
// @return QueryProjectMetaResponse
func (client *Client) QueryProjectMeta(request *QueryProjectMetaRequest) (_result *QueryProjectMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryProjectMetaResponse{}
	_body, _err := client.QueryProjectMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryStaticsAvailabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryStaticsAvailabilityResponse
func (client *Client) QueryStaticsAvailabilityWithOptions(request *QueryStaticsAvailabilityRequest, runtime *util.RuntimeOptions) (_result *QueryStaticsAvailabilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TimeRange)) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryStaticsAvailability"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryStaticsAvailabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryStaticsAvailabilityRequest
//
// @return QueryStaticsAvailabilityResponse
func (client *Client) QueryStaticsAvailability(request *QueryStaticsAvailabilityRequest) (_result *QueryStaticsAvailabilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryStaticsAvailabilityResponse{}
	_body, _err := client.QueryStaticsAvailabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryStaticsResponseTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryStaticsResponseTimeResponse
func (client *Client) QueryStaticsResponseTimeWithOptions(request *QueryStaticsResponseTimeRequest, runtime *util.RuntimeOptions) (_result *QueryStaticsResponseTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TimeRange)) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryStaticsResponseTime"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryStaticsResponseTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryStaticsResponseTimeRequest
//
// @return QueryStaticsResponseTimeResponse
func (client *Client) QueryStaticsResponseTime(request *QueryStaticsResponseTimeRequest) (_result *QueryStaticsResponseTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryStaticsResponseTimeResponse{}
	_body, _err := client.QueryStaticsResponseTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # QuerySystemEventCount
//
// @param request - QuerySystemEventCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySystemEventCountResponse
func (client *Client) QuerySystemEventCountWithOptions(request *QuerySystemEventCountRequest, runtime *util.RuntimeOptions) (_result *QuerySystemEventCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueryJson)) {
		query["QueryJson"] = request.QueryJson
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySystemEventCount"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySystemEventCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # QuerySystemEventCount
//
// @param request - QuerySystemEventCountRequest
//
// @return QuerySystemEventCountResponse
func (client *Client) QuerySystemEventCount(request *QuerySystemEventCountRequest) (_result *QuerySystemEventCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySystemEventCountResponse{}
	_body, _err := client.QuerySystemEventCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QuerySystemEventDemoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySystemEventDemoResponse
func (client *Client) QuerySystemEventDemoWithOptions(request *QuerySystemEventDemoRequest, runtime *util.RuntimeOptions) (_result *QuerySystemEventDemoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventName)) {
		query["EventName"] = request.EventName
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySystemEventDemo"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySystemEventDemoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuerySystemEventDemoRequest
//
// @return QuerySystemEventDemoResponse
func (client *Client) QuerySystemEventDemo(request *QuerySystemEventDemoRequest) (_result *QuerySystemEventDemoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySystemEventDemoResponse{}
	_body, _err := client.QuerySystemEventDemoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskConfigResponse
func (client *Client) QueryTaskConfigWithOptions(request *QueryTaskConfigRequest, runtime *util.RuntimeOptions) (_result *QueryTaskConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QueryTaskConfig"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryTaskConfigRequest
//
// @return QueryTaskConfigResponse
func (client *Client) QueryTaskConfig(request *QueryTaskConfigRequest) (_result *QueryTaskConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTaskConfigResponse{}
	_body, _err := client.QueryTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # QueryTaskMonitorData
//
// @param request - QueryTaskMonitorDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskMonitorDataResponse
func (client *Client) QueryTaskMonitorDataWithOptions(request *QueryTaskMonitorDataRequest, runtime *util.RuntimeOptions) (_result *QueryTaskMonitorDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["Cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Length)) {
		query["Length"] = request.Length
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["metricName"] = request.MetricName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTaskMonitorData"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTaskMonitorDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # QueryTaskMonitorData
//
// @param request - QueryTaskMonitorDataRequest
//
// @return QueryTaskMonitorDataResponse
func (client *Client) QueryTaskMonitorData(request *QueryTaskMonitorDataRequest) (_result *QueryTaskMonitorDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTaskMonitorDataResponse{}
	_body, _err := client.QueryTaskMonitorDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - TaskConfigListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskConfigListResponse
func (client *Client) TaskConfigListWithOptions(request *TaskConfigListRequest, runtime *util.RuntimeOptions) (_result *TaskConfigListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TaskConfigList"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TaskConfigListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - TaskConfigListRequest
//
// @return TaskConfigListResponse
func (client *Client) TaskConfigList(request *TaskConfigListRequest) (_result *TaskConfigListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TaskConfigListResponse{}
	_body, _err := client.TaskConfigListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # UpdateAlarm
//
// @param request - UpdateAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlarmResponse
func (client *Client) UpdateAlarmWithOptions(request *UpdateAlarmRequest, runtime *util.RuntimeOptions) (_result *UpdateAlarmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComparisonOperator)) {
		query["ComparisonOperator"] = request.ComparisonOperator
	}

	if !tea.BoolValue(util.IsUnset(request.ContactGroups)) {
		query["ContactGroups"] = request.ContactGroups
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EvaluationCount)) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyType)) {
		query["NotifyType"] = request.NotifyType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.SilenceTime)) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Statistics)) {
		query["Statistics"] = request.Statistics
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		query["Threshold"] = request.Threshold
	}

	if !tea.BoolValue(util.IsUnset(request.Webhook)) {
		query["Webhook"] = request.Webhook
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlarm"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # UpdateAlarm
//
// @param request - UpdateAlarmRequest
//
// @return UpdateAlarmResponse
func (client *Client) UpdateAlarm(request *UpdateAlarmRequest) (_result *UpdateAlarmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAlarmResponse{}
	_body, _err := client.UpdateAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateMonitoringTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMonitoringTemplateResponse
func (client *Client) UpdateMonitoringTemplateWithOptions(request *UpdateMonitoringTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateMonitoringTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertTemplatesJson)) {
		query["AlertTemplatesJson"] = request.AlertTemplatesJson
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.HostAvailabilityTemplate)) {
		query["HostAvailabilityTemplate"] = request.HostAvailabilityTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessMonitorTemplates)) {
		query["ProcessMonitorTemplates"] = request.ProcessMonitorTemplates
	}

	if !tea.BoolValue(util.IsUnset(request.RestVersion)) {
		query["RestVersion"] = request.RestVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SystemEventTemplates)) {
		query["SystemEventTemplates"] = request.SystemEventTemplates
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMonitoringTemplate"),
		Version:     tea.String("2018-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMonitoringTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateMonitoringTemplateRequest
//
// @return UpdateMonitoringTemplateResponse
func (client *Client) UpdateMonitoringTemplate(request *UpdateMonitoringTemplateRequest) (_result *UpdateMonitoringTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMonitoringTemplateResponse{}
	_body, _err := client.UpdateMonitoringTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
