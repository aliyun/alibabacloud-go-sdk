// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchDeleteJobsRequest struct {
	// The ID of the application. You can obtain the application ID on the **Application Management*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job IDs. Separate multiple job IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 99341
	JobIdList []*int64 `json:"JobIdList,omitempty" xml:"JobIdList,omitempty" type:"Repeated"`
	// The ID of the namespace to which the job belongs. You can obtain the ID of the namespace on the **Namespace*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// Schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The ID of the region to which the job belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information returned.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB5A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether multiple jobs were deleted at a time. Valid values:
	//
	// 	- **true**: Multiple jobs were deleted at a time.
	//
	// 	- **false**: Multiple jobs were not deleted at a time.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *BatchDeleteJobsResponse) SetStatusCode(v int32) *BatchDeleteJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteJobsResponse) SetBody(v *BatchDeleteJobsResponseBody) *BatchDeleteJobsResponse {
	s.Body = v
	return s
}

type BatchDeleteRouteStrategyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId   *string  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	JobIdList []*int64 `json:"JobIdList,omitempty" xml:"JobIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BatchDeleteRouteStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteRouteStrategyRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteRouteStrategyRequest) SetGroupId(v string) *BatchDeleteRouteStrategyRequest {
	s.GroupId = &v
	return s
}

func (s *BatchDeleteRouteStrategyRequest) SetJobIdList(v []*int64) *BatchDeleteRouteStrategyRequest {
	s.JobIdList = v
	return s
}

func (s *BatchDeleteRouteStrategyRequest) SetNamespace(v string) *BatchDeleteRouteStrategyRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDeleteRouteStrategyRequest) SetRegionId(v string) *BatchDeleteRouteStrategyRequest {
	s.RegionId = &v
	return s
}

type BatchDeleteRouteStrategyResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// job is not existed, jobId=162837
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 704A2A61-3681-5568-92F7-2DFCC53F33D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchDeleteRouteStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteRouteStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteRouteStrategyResponseBody) SetCode(v int32) *BatchDeleteRouteStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeleteRouteStrategyResponseBody) SetMessage(v string) *BatchDeleteRouteStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDeleteRouteStrategyResponseBody) SetRequestId(v string) *BatchDeleteRouteStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteRouteStrategyResponseBody) SetSuccess(v bool) *BatchDeleteRouteStrategyResponseBody {
	s.Success = &v
	return s
}

type BatchDeleteRouteStrategyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteRouteStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteRouteStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteRouteStrategyResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteRouteStrategyResponse) SetHeaders(v map[string]*string) *BatchDeleteRouteStrategyResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteRouteStrategyResponse) SetStatusCode(v int32) *BatchDeleteRouteStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteRouteStrategyResponse) SetBody(v *BatchDeleteRouteStrategyResponseBody) *BatchDeleteRouteStrategyResponse {
	s.Body = v
	return s
}

type BatchDisableJobsRequest struct {
	// The ID of the application. You can obtain the application ID on the **Application Management*	- page in the SchedulerX console.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job IDs. Separate multiple job IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 99341
	JobIdList []*int64 `json:"JobIdList,omitempty" xml:"JobIdList,omitempty" type:"Repeated"`
	// The ID of the namespace to which the job belongs. You can obtain the ID of the namespace on the **Namespace*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The ID of the region to which the job belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that was returned.
	//
	// example:
	//
	// disable failed jobs=[99341]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB5A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDisableJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *BatchDisableJobsResponse) SetStatusCode(v int32) *BatchDisableJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDisableJobsResponse) SetBody(v *BatchDisableJobsResponseBody) *BatchDisableJobsResponse {
	s.Body = v
	return s
}

type BatchEnableJobsRequest struct {
	// The application ID. You can obtain the application ID on the **Application Management*	- page in the SchedulerX console.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job IDs. Multiple job IDs are separated with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 99341
	JobIdList []*int64 `json:"JobIdList,omitempty" xml:"JobIdList,omitempty" type:"Repeated"`
	// The ID of the namespace to which the job belongs. You can obtain the namespace ID on the **Namespace*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The ID of the region to which the job belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned additional information.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB5A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the jobs were enabled at a time. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchEnableJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *BatchEnableJobsResponse) SetStatusCode(v int32) *BatchEnableJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchEnableJobsResponse) SetBody(v *BatchEnableJobsResponseBody) *BatchEnableJobsResponse {
	s.Body = v
	return s
}

type CreateAppGroupRequest struct {
	// The AppKey for the application.
	//
	// example:
	//
	// adcExHZviLcl****
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// DocTest
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用类型。
	//
	// - 1、普通应用。
	//
	// - 2、k8s应用。
	//
	// example:
	//
	// 1
	AppType *int32 `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 是否开启日志。
	//
	// - true：开启
	//
	// - false：关闭
	//
	// example:
	//
	// true
	EnableLog *bool `json:"EnableLog,omitempty" xml:"EnableLog,omitempty"`
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The maximum number of jobs.
	//
	// example:
	//
	// 1000
	MaxJobs *int32 `json:"MaxJobs,omitempty" xml:"MaxJobs,omitempty"`
	// The configuration of the alert. The value is a JSON string. For more information about this parameter, see **Additional information about request parameters**.
	//
	// example:
	//
	// {"sendChannel":"sms,ding"}
	MonitorConfigJson *string `json:"MonitorConfigJson,omitempty" xml:"MonitorConfigJson,omitempty"`
	// The configuration of alert contacts. The value is a JSON string.
	//
	// example:
	//
	// [{"userName":"Tom","userPhone":"89756\\*\\*\\*\\*\\*\\*"},{"userName":"Bob","ding":"http://www.example.com"}]
	MonitorContactsJson *string `json:"MonitorContactsJson,omitempty" xml:"MonitorContactsJson,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// Test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// This parameter is not supported. You do not need to specify this parameter.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to schedule a busy worker.
	//
	// example:
	//
	// false
	ScheduleBusyWorkers *bool `json:"ScheduleBusyWorkers,omitempty" xml:"ScheduleBusyWorkers,omitempty"`
	// example:
	//
	// 2
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateAppGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAppGroupRequest) SetAppKey(v string) *CreateAppGroupRequest {
	s.AppKey = &v
	return s
}

func (s *CreateAppGroupRequest) SetAppName(v string) *CreateAppGroupRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppGroupRequest) SetAppType(v int32) *CreateAppGroupRequest {
	s.AppType = &v
	return s
}

func (s *CreateAppGroupRequest) SetDescription(v string) *CreateAppGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateAppGroupRequest) SetEnableLog(v bool) *CreateAppGroupRequest {
	s.EnableLog = &v
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

func (s *CreateAppGroupRequest) SetMonitorConfigJson(v string) *CreateAppGroupRequest {
	s.MonitorConfigJson = &v
	return s
}

func (s *CreateAppGroupRequest) SetMonitorContactsJson(v string) *CreateAppGroupRequest {
	s.MonitorContactsJson = &v
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

func (s *CreateAppGroupRequest) SetScheduleBusyWorkers(v bool) *CreateAppGroupRequest {
	s.ScheduleBusyWorkers = &v
	return s
}

func (s *CreateAppGroupRequest) SetVersion(v int32) *CreateAppGroupRequest {
	s.Version = &v
	return s
}

type CreateAppGroupResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the job group.
	Data *CreateAppGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// Your request is denied as lack of ssl protect.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 883AFE93-FB03-4FA9-A958-E750C6DE120C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application was created. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The job group ID.
	//
	// example:
	//
	// 6607
	AppGroupId *int64 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	// The AppKey for the application.
	//
	// example:
	//
	// adcExHZviL******
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
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

func (s *CreateAppGroupResponseBodyData) SetAppKey(v string) *CreateAppGroupResponseBodyData {
	s.AppKey = &v
	return s
}

type CreateAppGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateAppGroupResponse) SetStatusCode(v int32) *CreateAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppGroupResponse) SetBody(v *CreateAppGroupResponseBody) *CreateAppGroupResponse {
	s.Body = v
	return s
}

type CreateJobRequest struct {
	// The interval of retries after a job failure. Default value: 30. Unit: seconds.
	//
	// example:
	//
	// 30
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// If you set TimeType to 1 (cron), you can specify calendar days.
	//
	// example:
	//
	// This parameter is not supported. You do not need to specify this parameter.
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The full path of the job interface class.
	//
	// This parameter is available only when you set JobType to java. You must enter a full path.
	//
	// example:
	//
	// com.alibaba.schedulerx.test.helloworld
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The number of threads that are triggered by a single worker at a time. Default value: 5. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// 5
	ConsumerSize *int32 `json:"ConsumerSize,omitempty" xml:"ConsumerSize,omitempty"`
	// The information about the alert contact.
	ContactInfo []*CreateJobRequestContactInfo `json:"ContactInfo,omitempty" xml:"ContactInfo,omitempty" type:"Repeated"`
	// The script content. This parameter is required when you set JobType to python, shell, go, or k8s.
	//
	// example:
	//
	// echo \\"hello\\"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// If you set TimeType to 1 (cron), you can specify a time offset. Unit: seconds.
	//
	// example:
	//
	// 2400
	DataOffset *int32 `json:"DataOffset,omitempty" xml:"DataOffset,omitempty"`
	// The job description.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of task distribution threads. Default value: 5. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// 5
	DispatcherSize *int32 `json:"DispatcherSize,omitempty" xml:"DispatcherSize,omitempty"`
	// The execution mode of the job. Valid values:
	//
	// 	- **Stand-alone operation**
	//
	// 	- **Broadcast run**
	//
	// 	- **Visual MapReduce**
	//
	// 	- **MapReduce**
	//
	// 	- **Shard run**
	//
	// This parameter is required.
	//
	// example:
	//
	// standalone
	ExecuteMode *string `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	// Specifies whether to turn on Failure alarm. If the switch is turned on, an alert will be generated upon a failure. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	FailEnable *bool `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	// The number of consecutive failures. An alert will be received if the number of consecutive failures reaches the value of this parameter.
	//
	// example:
	//
	// 2
	FailTimes *int32 `json:"FailTimes,omitempty" xml:"FailTimes,omitempty"`
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job type. Valid values:
	//
	// 	- java
	//
	// 	- python
	//
	// 	- shell
	//
	// 	- go
	//
	// 	- http
	//
	// 	- xxljob
	//
	// 	- dataworks
	//
	// 	- k8s
	//
	// 	- springschedule
	//
	// This parameter is required.
	//
	// example:
	//
	// java
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The maximum number of retries after a job failure. Specify this parameter based on your business requirements. Default value: 0.
	//
	// example:
	//
	// 0
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// The maximum number of concurrent instances. Default value: 1. The default value indicates that only one instance is allowed to run at a time. When an instance is running, another instance is not triggered even if the scheduled time for running the instance is reached.
	//
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// Specifies whether to turn on No machine alarm available. If the switch is turned on, an alert will be generated when no machine is available for running the job. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	MissWorkerEnable *bool `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	// The job name.
	//
	// This parameter is required.
	//
	// example:
	//
	// helloworld
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The namespace source. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The number of tasks that can be pulled at a time. Default value: 100. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The user-defined parameters that you can obtain when the job is running.
	//
	// example:
	//
	// test
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The maximum number of tasks that can be queued. Default value: 10000. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// 10000
	QueueSize *int32 `json:"QueueSize,omitempty" xml:"QueueSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The method that is used to send alerts. Only Short Message Service (SMS) is supported. Default value: sms.
	//
	// example:
	//
	// sms
	SendChannel *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	// Specifies whether to enable the job. If this parameter is set to 0, the job is disabled. If this parameter is set to 1, the job is enabled. Default value: 1.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Specifies whether to turn on Successful notice. If the switch is turned on, a notice will be sent when a job succeeds.
	//
	// example:
	//
	// false
	SuccessNoticeEnable *bool `json:"SuccessNoticeEnable,omitempty" xml:"SuccessNoticeEnable,omitempty"`
	// The interval of retries after a task failure. Default value: 0. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// 0
	TaskAttemptInterval *int32 `json:"TaskAttemptInterval,omitempty" xml:"TaskAttemptInterval,omitempty"`
	// The number of retries after a task failure. Default value: 0. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// 0
	TaskMaxAttempt *int32 `json:"TaskMaxAttempt,omitempty" xml:"TaskMaxAttempt,omitempty"`
	// The time expression. Specify the time expression based on the value of TimeType:
	//
	// 	- If you set TimeType to **1*	- (cron), specify this parameter to a standard CRON expression.
	//
	// 	- If you set TimeType to **100*	- (api), no time expression is required.
	//
	// 	- If you set TimeType to **3*	- (fixed_rate), specify this parameter to a fixed frequency in seconds. For example, if you set this parameter to 30, the system triggers a job every 30 seconds.
	//
	// 	- If you set TimeType to **4*	- (second_delay), specify this parameter to a fixed delay after which the job is triggered. Valid values: 1 to 60. Unit: seconds.
	//
	// 	- If you set TimeType to **5*	- (one_time), specify this parameter to a specific time point at which the job is triggered. The time is in the format of yyyy-MM-dd HH:mm:ss, such as 2022-10-10 10:10:00, or a timestamp in milliseconds.
	//
	// example:
	//
	// 0 0/10 	- 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time type. Valid values:
	//
	// 	- **1**: cron
	//
	// 	- **3**: fixed_rate
	//
	// 	- **4**: second_delay
	//
	// 	- **5**: one_time
	//
	// 	- **100**: api
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The timeout threshold. Default value: 7200. Unit: seconds.
	//
	// example:
	//
	// 7200
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Specifies whether to turn on Timeout alarm. If the switch is turned on, an alert will be generated upon a timeout. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	TimeoutEnable *bool `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	// Specifies whether to turn on Timeout termination. If the switch is turned on, the job will be terminated upon a timeout. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	TimeoutKillEnable *bool `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
	// Time zone.
	//
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// If you set JobType to k8s, this parameter is required. xxljob task: {"resource":"job"} shell task: {"image":"busybox","resource":"shell"}
	//
	// example:
	//
	// {"resource":"job"}
	XAttrs *string `json:"XAttrs,omitempty" xml:"XAttrs,omitempty"`
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

func (s *CreateJobRequest) SetFailTimes(v int32) *CreateJobRequest {
	s.FailTimes = &v
	return s
}

func (s *CreateJobRequest) SetGroupId(v string) *CreateJobRequest {
	s.GroupId = &v
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

func (s *CreateJobRequest) SetStatus(v int32) *CreateJobRequest {
	s.Status = &v
	return s
}

func (s *CreateJobRequest) SetSuccessNoticeEnable(v bool) *CreateJobRequest {
	s.SuccessNoticeEnable = &v
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

func (s *CreateJobRequest) SetTimezone(v string) *CreateJobRequest {
	s.Timezone = &v
	return s
}

func (s *CreateJobRequest) SetXAttrs(v string) *CreateJobRequest {
	s.XAttrs = &v
	return s
}

type CreateJobRequestContactInfo struct {
	// The webhook URL of the DingTalk chatbot.[](https://open.dingtalk.com/document/org/application-types)
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=**********
	Ding *string `json:"Ding,omitempty" xml:"Ding,omitempty"`
	// The email address of the alert contact.
	//
	// example:
	//
	// test***@***.com
	UserMail *string `json:"UserMail,omitempty" xml:"UserMail,omitempty"`
	// The name of the alert contact.
	//
	// example:
	//
	// Tom
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The mobile number of the alert contact.
	//
	// example:
	//
	// 1381111****
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the job.
	Data *CreateJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The additional information returned.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// If you set JobType to k8s, this parameter is required. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
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
	// The job ID.
	//
	// example:
	//
	// 92583
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

type CreateNamespaceRequest struct {
	// The description of the namespace.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-env
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The unique identifier (UID) of the namespace. We recommend that you use the universally unique identifier (UUID) to generate the UID.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s CreateNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) SetDescription(v string) *CreateNamespaceRequest {
	s.Description = &v
	return s
}

func (s *CreateNamespaceRequest) SetName(v string) *CreateNamespaceRequest {
	s.Name = &v
	return s
}

func (s *CreateNamespaceRequest) SetRegionId(v string) *CreateNamespaceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNamespaceRequest) SetUid(v string) *CreateNamespaceRequest {
	s.Uid = &v
	return s
}

type CreateNamespaceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the namespace.
	Data *CreateNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// namespace=test3 is existed, noting update
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application was created. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBody) SetCode(v int32) *CreateNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetData(v *CreateNamespaceResponseBodyData) *CreateNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *CreateNamespaceResponseBody) SetMessage(v string) *CreateNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetRequestId(v string) *CreateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetSuccess(v bool) *CreateNamespaceResponseBody {
	s.Success = &v
	return s
}

type CreateNamespaceResponseBodyData struct {
	// The UID of the namespace.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	NamespaceUid *string `json:"NamespaceUid,omitempty" xml:"NamespaceUid,omitempty"`
}

func (s CreateNamespaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBodyData) SetNamespaceUid(v string) *CreateNamespaceResponseBodyData {
	s.NamespaceUid = &v
	return s
}

type CreateNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponse) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponse) SetHeaders(v map[string]*string) *CreateNamespaceResponse {
	s.Headers = v
	return s
}

func (s *CreateNamespaceResponse) SetStatusCode(v int32) *CreateNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNamespaceResponse) SetBody(v *CreateNamespaceResponseBody) *CreateNamespaceResponse {
	s.Body = v
	return s
}

type CreateRouteStrategyRequest struct {
	// The ID of the application group. You can obtain the ID on the **Application Management*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID. You can obtain the ID on the **Task Management*	- page in the SchedulerX console.
	//
	// example:
	//
	// 54978
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the routing policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-strategy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the **Namespace*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable the routing policy. Valid values:
	//
	// 	- **0**: disables the routing policy.
	//
	// 	- **1**: enables the routing policy.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The details of the routing policy. The value is a JSON string. For more information about this parameter, see **the additional information about request parameters*	- below this table.
	//
	// example:
	//
	// [{"percentage":20,"target":"[\\"version1\\"]","targetType":"label"}]
	StrategyContent *string `json:"StrategyContent,omitempty" xml:"StrategyContent,omitempty"`
	// The type of the routing policy. Valid value:
	//
	// 	- **3**: routes by proportion.
	//
	// example:
	//
	// 3
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRouteStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRouteStrategyRequest) GoString() string {
	return s.String()
}

func (s *CreateRouteStrategyRequest) SetGroupId(v string) *CreateRouteStrategyRequest {
	s.GroupId = &v
	return s
}

func (s *CreateRouteStrategyRequest) SetJobId(v int64) *CreateRouteStrategyRequest {
	s.JobId = &v
	return s
}

func (s *CreateRouteStrategyRequest) SetName(v string) *CreateRouteStrategyRequest {
	s.Name = &v
	return s
}

func (s *CreateRouteStrategyRequest) SetNamespace(v string) *CreateRouteStrategyRequest {
	s.Namespace = &v
	return s
}

func (s *CreateRouteStrategyRequest) SetRegionId(v string) *CreateRouteStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRouteStrategyRequest) SetStatus(v int32) *CreateRouteStrategyRequest {
	s.Status = &v
	return s
}

func (s *CreateRouteStrategyRequest) SetStrategyContent(v string) *CreateRouteStrategyRequest {
	s.StrategyContent = &v
	return s
}

func (s *CreateRouteStrategyRequest) SetType(v int32) *CreateRouteStrategyRequest {
	s.Type = &v
	return s
}

type CreateRouteStrategyResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateRouteStrategyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The additional information, including errors and tips.
	//
	// example:
	//
	// strategy name is null or empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRouteStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRouteStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRouteStrategyResponseBody) SetCode(v int32) *CreateRouteStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRouteStrategyResponseBody) SetData(v *CreateRouteStrategyResponseBodyData) *CreateRouteStrategyResponseBody {
	s.Data = v
	return s
}

func (s *CreateRouteStrategyResponseBody) SetMessage(v string) *CreateRouteStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRouteStrategyResponseBody) SetRequestId(v string) *CreateRouteStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRouteStrategyResponseBody) SetSuccess(v bool) *CreateRouteStrategyResponseBody {
	s.Success = &v
	return s
}

type CreateRouteStrategyResponseBodyData struct {
}

func (s CreateRouteStrategyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateRouteStrategyResponseBodyData) GoString() string {
	return s.String()
}

type CreateRouteStrategyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRouteStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRouteStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRouteStrategyResponse) GoString() string {
	return s.String()
}

func (s *CreateRouteStrategyResponse) SetHeaders(v map[string]*string) *CreateRouteStrategyResponse {
	s.Headers = v
	return s
}

func (s *CreateRouteStrategyResponse) SetStatusCode(v int32) *CreateRouteStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRouteStrategyResponse) SetBody(v *CreateRouteStrategyResponseBody) *CreateRouteStrategyResponse {
	s.Body = v
	return s
}

type CreateWorkflowRequest struct {
	// The description of the workflow.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application group ID. You can obtain the ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The maximum number of workflow instances that can be run at the same time. Default value: 1. The value 1 indicates that only one workflow instance is allowed. In this case, if the triggered workflow instance is still ongoing, no more workflow instances can be triggered even the time to schedule the next workflow arrives.
	//
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// helloworld
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time expression. You can set the time expression based on the selected method that is used to specify time.
	//
	// 	- If you set the TimeType parameter to cron, you need to enter a standard cron expression. Online verification is supported.
	//
	// 	- If you set the TimeType parameter to api, no time expression is required.
	//
	// example:
	//
	// 0 0/10 	- 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The method that is used to specify the time. Valid values:
	//
	// 	- 1: cron
	//
	// 	- 100: api
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The time zone.
	//
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s CreateWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowRequest) SetDescription(v string) *CreateWorkflowRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkflowRequest) SetGroupId(v string) *CreateWorkflowRequest {
	s.GroupId = &v
	return s
}

func (s *CreateWorkflowRequest) SetMaxConcurrency(v int32) *CreateWorkflowRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *CreateWorkflowRequest) SetName(v string) *CreateWorkflowRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkflowRequest) SetNamespace(v string) *CreateWorkflowRequest {
	s.Namespace = &v
	return s
}

func (s *CreateWorkflowRequest) SetNamespaceSource(v string) *CreateWorkflowRequest {
	s.NamespaceSource = &v
	return s
}

func (s *CreateWorkflowRequest) SetRegionId(v string) *CreateWorkflowRequest {
	s.RegionId = &v
	return s
}

func (s *CreateWorkflowRequest) SetTimeExpression(v string) *CreateWorkflowRequest {
	s.TimeExpression = &v
	return s
}

func (s *CreateWorkflowRequest) SetTimeType(v int32) *CreateWorkflowRequest {
	s.TimeType = &v
	return s
}

func (s *CreateWorkflowRequest) SetTimezone(v string) *CreateWorkflowRequest {
	s.Timezone = &v
	return s
}

type CreateWorkflowResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that was returned for the request.
	Data *CreateWorkflowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// timetype is invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the workflow was created. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWorkflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkflowResponseBody) SetCode(v string) *CreateWorkflowResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWorkflowResponseBody) SetData(v *CreateWorkflowResponseBodyData) *CreateWorkflowResponseBody {
	s.Data = v
	return s
}

func (s *CreateWorkflowResponseBody) SetMessage(v string) *CreateWorkflowResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWorkflowResponseBody) SetRequestId(v string) *CreateWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkflowResponseBody) SetSuccess(v bool) *CreateWorkflowResponseBody {
	s.Success = &v
	return s
}

type CreateWorkflowResponseBodyData struct {
	// The workflow ID.
	//
	// example:
	//
	// 92583
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s CreateWorkflowResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWorkflowResponseBodyData) SetWorkflowId(v int64) *CreateWorkflowResponseBodyData {
	s.WorkflowId = &v
	return s
}

type CreateWorkflowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkflowResponse) SetHeaders(v map[string]*string) *CreateWorkflowResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkflowResponse) SetStatusCode(v int32) *CreateWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkflowResponse) SetBody(v *CreateWorkflowResponseBody) *CreateWorkflowResponse {
	s.Body = v
	return s
}

type DeleteAppGroupRequest struct {
	// example:
	//
	// true
	DeleteJobs *bool `json:"DeleteJobs,omitempty" xml:"DeleteJobs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAppGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppGroupRequest) SetDeleteJobs(v bool) *DeleteAppGroupRequest {
	s.DeleteJobs = &v
	return s
}

func (s *DeleteAppGroupRequest) SetGroupId(v string) *DeleteAppGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteAppGroupRequest) SetNamespace(v string) *DeleteAppGroupRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteAppGroupRequest) SetRegionId(v string) *DeleteAppGroupRequest {
	s.RegionId = &v
	return s
}

type DeleteAppGroupResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// app is not existed, groupId=xxxx, namesapce=xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppGroupResponseBody) SetCode(v int32) *DeleteAppGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppGroupResponseBody) SetMessage(v string) *DeleteAppGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAppGroupResponseBody) SetRequestId(v string) *DeleteAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppGroupResponseBody) SetSuccess(v bool) *DeleteAppGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteAppGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppGroupResponse) SetHeaders(v map[string]*string) *DeleteAppGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppGroupResponse) SetStatusCode(v int32) *DeleteAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppGroupResponse) SetBody(v *DeleteAppGroupResponseBody) *DeleteAppGroupResponse {
	s.Body = v
	return s
}

type DeleteJobRequest struct {
	// The ID of the application. You can obtain the application ID on the **Application Management*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the job. You can obtain the ID on the **Task Management*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the namespace. You can obtain the ID of the namespace on the **Namespace*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information returned.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the job was deleted. Valid values:
	//
	// 	- **true**: The job was deleted.
	//
	// 	- **false**: The job was not deleted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteJobResponse) SetStatusCode(v int32) *DeleteJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobResponse) SetBody(v *DeleteJobResponseBody) *DeleteJobResponse {
	s.Body = v
	return s
}

type DeleteRouteStrategyRequest struct {
	// The application ID. You can obtain the application ID on the **Application Management*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID. You can obtain the job ID on the **Task Management*	- page in the SchedulerX console.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The namespace ID. You can obtain the ID of the namespace on the **Namespace*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteRouteStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRouteStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteStrategyRequest) SetGroupId(v string) *DeleteRouteStrategyRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteRouteStrategyRequest) SetJobId(v int64) *DeleteRouteStrategyRequest {
	s.JobId = &v
	return s
}

func (s *DeleteRouteStrategyRequest) SetNamespace(v string) *DeleteRouteStrategyRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteRouteStrategyRequest) SetRegionId(v string) *DeleteRouteStrategyRequest {
	s.RegionId = &v
	return s
}

type DeleteRouteStrategyResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// strategy is already deleted.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB5A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRouteStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRouteStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouteStrategyResponseBody) SetCode(v int32) *DeleteRouteStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRouteStrategyResponseBody) SetMessage(v string) *DeleteRouteStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRouteStrategyResponseBody) SetRequestId(v string) *DeleteRouteStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRouteStrategyResponseBody) SetSuccess(v bool) *DeleteRouteStrategyResponseBody {
	s.Success = &v
	return s
}

type DeleteRouteStrategyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRouteStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRouteStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRouteStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouteStrategyResponse) SetHeaders(v map[string]*string) *DeleteRouteStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouteStrategyResponse) SetStatusCode(v int32) *DeleteRouteStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRouteStrategyResponse) SetBody(v *DeleteRouteStrategyResponseBody) *DeleteRouteStrategyResponse {
	s.Body = v
	return s
}

type DeleteWorkflowRequest struct {
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 111
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// Your request is denied as lack of ssl protect.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the workflow was deleted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteWorkflowResponse) SetStatusCode(v int32) *DeleteWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkflowResponse) SetBody(v *DeleteWorkflowResponseBody) *DeleteWorkflowResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that was returned only if the corresponding error occurred.
	//
	// example:
	//
	// disable failed jobs=[99341]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The available regions.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The display name of the region, which varies based on the current language.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the region.
	//
	// example:
	//
	// schedulerx.cn-hangzhou.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DesignateWorkersRequest struct {
	// The type of the machines to be designated. Valid values: 1 and 2. The value 1 specifies the worker type. The value 2 specifies the label type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DesignateType *int32 `json:"DesignateType,omitempty" xml:"DesignateType,omitempty"`
	// The application group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hxm.test
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 144153
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The designated `labels`. Specify the value of the parameter in a `JSON` string.
	//
	// example:
	//
	// ["gray"]
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The unique identifier (UID) of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a06d5ea-f576-4326-842c-fb14ea043d8d
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to allow a failover.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Transferable *bool `json:"Transferable,omitempty" xml:"Transferable,omitempty"`
	// The designated machines. Specify the value of the parameter in a JSON string.
	//
	// example:
	//
	// ["127.0.0.1","127.0.0.2"]
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// job is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 765xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DesignateWorkersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DesignateWorkersResponse) SetStatusCode(v int32) *DesignateWorkersResponse {
	s.StatusCode = &v
	return s
}

func (s *DesignateWorkersResponse) SetBody(v *DesignateWorkersResponseBody) *DesignateWorkersResponse {
	s.Body = v
	return s
}

type DisableJobRequest struct {
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID. You can obtain the job ID on the Task Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// jobid: 92583 not match groupId: testSchedulerx.defaultGroup
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C8E5FB4A-6D8D-424D-9AAA-4FE06BB74FF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the job was disabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DisableJobResponse) SetStatusCode(v int32) *DisableJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableJobResponse) SetBody(v *DisableJobResponseBody) *DisableJobResponse {
	s.Body = v
	return s
}

type DisableWorkflowRequest struct {
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 111
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// Your request is denied as lack of ssl protect.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the workflow was disabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DisableWorkflowResponse) SetStatusCode(v int32) *DisableWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableWorkflowResponse) SetBody(v *DisableWorkflowResponseBody) *DisableWorkflowResponse {
	s.Body = v
	return s
}

type EnableJobRequest struct {
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID. You can obtain the job ID on the Task Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// jobid: 92583 not match groupId: testSchedulerx.defaultGroup
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB5A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *EnableJobResponse) SetStatusCode(v int32) *EnableJobResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableJobResponse) SetBody(v *EnableJobResponseBody) *EnableJobResponse {
	s.Body = v
	return s
}

type EnableWorkflowRequest struct {
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 111
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// Your request is denied as lack of ssl protect.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the workflow was enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *EnableWorkflowResponse) SetStatusCode(v int32) *EnableWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableWorkflowResponse) SetBody(v *EnableWorkflowResponseBody) *EnableWorkflowResponse {
	s.Body = v
	return s
}

type ExecuteJobRequest struct {
	// Specifies whether to check the job status. Valid values: -**true**: The job can be run only if the job is enabled. -**false**: The job can be run even if the job is disabled.
	//
	// example:
	//
	// true
	CheckJobStatus *bool `json:"CheckJobStatus,omitempty" xml:"CheckJobStatus,omitempty"`
	// The type of the designated machine. Valid values: -**1**: worker. -**2**: label.
	//
	// example:
	//
	// 1
	DesignateType *int32 `json:"DesignateType,omitempty" xml:"DesignateType,omitempty"`
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The parameters that are passed to trigger the job to run. The input value can be a random string. The parameters that are passed are obtained by calling the `context.getInstanceParameters()` class in the `processor` code. The parameters are different from custom parameters for creating jobs.
	//
	// example:
	//
	// test
	InstanceParameters *string `json:"InstanceParameters,omitempty" xml:"InstanceParameters,omitempty"`
	// The job ID. You can obtain the job ID on the Task Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The label of the worker.
	//
	// example:
	//
	// gray
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The worker address of the application. To query the worker address, call the GetWokerList operation.
	//
	// example:
	//
	// xxxxxxx@127.0.0.1:222
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the job instance that is returned if the request is successful.
	Data *ExecuteJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// groupid not exist groupId: testSchedulerx.defaultGroup namespace: adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The job instance ID.
	//
	// example:
	//
	// 11111111
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ExecuteJobResponse) SetStatusCode(v int32) *ExecuteJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteJobResponse) SetBody(v *ExecuteJobResponseBody) *ExecuteJobResponse {
	s.Body = v
	return s
}

type ExecuteWorkflowRequest struct {
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The dynamic parameter of the workflow instance. The value of the parameter can be up to 1,000 bytes in length.
	//
	// example:
	//
	// test
	InstanceParameters *string `json:"InstanceParameters,omitempty" xml:"InstanceParameters,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 111
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
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

func (s *ExecuteWorkflowRequest) SetRegionId(v string) *ExecuteWorkflowRequest {
	s.RegionId = &v
	return s
}

func (s *ExecuteWorkflowRequest) SetWorkflowId(v int64) *ExecuteWorkflowRequest {
	s.WorkflowId = &v
	return s
}

type ExecuteWorkflowResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// If the request is successful, the ID of the workflow instance is returned.
	Data *ExecuteWorkflowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// Cannot find product according to your domain.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The workflow instance ID.
	//
	// example:
	//
	// 111111
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ExecuteWorkflowResponse) SetStatusCode(v int32) *ExecuteWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteWorkflowResponse) SetBody(v *ExecuteWorkflowResponseBody) *ExecuteWorkflowResponse {
	s.Body = v
	return s
}

type GetAppGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAppGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppGroupRequest) GoString() string {
	return s.String()
}

func (s *GetAppGroupRequest) SetGroupId(v string) *GetAppGroupRequest {
	s.GroupId = &v
	return s
}

func (s *GetAppGroupRequest) SetNamespace(v string) *GetAppGroupRequest {
	s.Namespace = &v
	return s
}

func (s *GetAppGroupRequest) SetRegionId(v string) *GetAppGroupRequest {
	s.RegionId = &v
	return s
}

type GetAppGroupResponseBody struct {
	// example:
	//
	// 200
	Code *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAppGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// app is not existed, groupId=xxxx, namesapce=xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppGroupResponseBody) SetCode(v int32) *GetAppGroupResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppGroupResponseBody) SetData(v *GetAppGroupResponseBodyData) *GetAppGroupResponseBody {
	s.Data = v
	return s
}

func (s *GetAppGroupResponseBody) SetMessage(v string) *GetAppGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppGroupResponseBody) SetRequestId(v string) *GetAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppGroupResponseBody) SetSuccess(v bool) *GetAppGroupResponseBody {
	s.Success = &v
	return s
}

type GetAppGroupResponseBodyData struct {
	// example:
	//
	// QI4lWMZ+xk1rNB67jFUhaw==
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// DocTest
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	CurJobs *int32 `json:"CurJobs,omitempty" xml:"CurJobs,omitempty"`
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 1000
	MaxJobs *int32 `json:"MaxJobs,omitempty" xml:"MaxJobs,omitempty"`
	// example:
	//
	// {"sendChannel":"sms,mail,ding"}
	MonitorConfigJson *string `json:"MonitorConfigJson,omitempty" xml:"MonitorConfigJson,omitempty"`
}

func (s GetAppGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAppGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppGroupResponseBodyData) SetAppKey(v string) *GetAppGroupResponseBodyData {
	s.AppKey = &v
	return s
}

func (s *GetAppGroupResponseBodyData) SetAppName(v string) *GetAppGroupResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetAppGroupResponseBodyData) SetCurJobs(v int32) *GetAppGroupResponseBodyData {
	s.CurJobs = &v
	return s
}

func (s *GetAppGroupResponseBodyData) SetDescription(v string) *GetAppGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAppGroupResponseBodyData) SetGroupId(v string) *GetAppGroupResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *GetAppGroupResponseBodyData) SetMaxJobs(v int32) *GetAppGroupResponseBodyData {
	s.MaxJobs = &v
	return s
}

func (s *GetAppGroupResponseBodyData) SetMonitorConfigJson(v string) *GetAppGroupResponseBodyData {
	s.MonitorConfigJson = &v
	return s
}

type GetAppGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppGroupResponse) GoString() string {
	return s.String()
}

func (s *GetAppGroupResponse) SetHeaders(v map[string]*string) *GetAppGroupResponse {
	s.Headers = v
	return s
}

func (s *GetAppGroupResponse) SetStatusCode(v int32) *GetAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppGroupResponse) SetBody(v *GetAppGroupResponseBody) *GetAppGroupResponse {
	s.Body = v
	return s
}

type GetJobInfoRequest struct {
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID. You can obtain the job ID on the Task Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job name.
	//
	// example:
	//
	// simpleJob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The namespace source. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *GetJobInfoRequest) SetJobName(v string) *GetJobInfoRequest {
	s.JobName = &v
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the job.
	Data *GetJobInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned only if an error occurs.
	//
	// example:
	//
	// jobid: 92583 not match groupId: testSchedulerx.defaultGroup
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the job details were obtained. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The configurations of the job.
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
	// The interval at which the system retried to run the job after a job failure. Default value: 30. Unit: seconds.
	//
	// example:
	//
	// 30
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// The full path of the job interface class. This parameter is returned only for jobs whose job type is Java.
	//
	// example:
	//
	// com.alibaba.test.helloword
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The script of a script job.
	//
	// example:
	//
	// echo "clear" > /home/admin/edas-container/logs/catalina.out
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description of the job.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution mode of the job. Valid values:
	//
	// 	- **Stand-alone operation**
	//
	// 	- **Broadcast run**
	//
	// 	- **Visual MapReduce**
	//
	// 	- **MapReduce**
	//
	// 	- ****
	//
	// 	- **Shard run**
	//
	// example:
	//
	// standalone
	ExecuteMode *string `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	// The full path used to upload files to Object Storage Service (OSS).
	//
	// If you use a JAR package, you can upload the JAR package to this OSS path.
	//
	// example:
	//
	// https://test.oss-cn-hangzhou.aliyuncs.com/schedulerX/test.jar
	JarUrl *string `json:"JarUrl,omitempty" xml:"JarUrl,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 538039
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The monitoring information of the job.
	JobMonitorInfo *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo `json:"JobMonitorInfo,omitempty" xml:"JobMonitorInfo,omitempty" type:"Struct"`
	// The job type.
	//
	// example:
	//
	// java
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The advanced configurations of the job.
	MapTaskXAttrs *GetJobInfoResponseBodyDataJobConfigInfoMapTaskXAttrs `json:"MapTaskXAttrs,omitempty" xml:"MapTaskXAttrs,omitempty" type:"Struct"`
	// The maximum number of retries after a job failure. This parameter was specified based on your business requirements. Default value: 0.
	//
	// example:
	//
	// 0
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// The maximum number of concurrent instances. Default value: 1. The default value indicates that if the last triggered instance is running, the next instance is not triggered even if the scheduled point in time for running the next instance is reached.
	//
	// example:
	//
	// 1
	MaxConcurrency *string `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The job name.
	//
	// example:
	//
	// helloworld
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user-defined parameters that you can obtain when the job is running.
	//
	// example:
	//
	// test
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// Indicates whether the job was enabled. Valid values:
	//
	// 	- **1**: The job was enabled and could be triggered.
	//
	// 	- **0**: The job was disabled and could not be triggered.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time configurations.
	TimeConfig *GetJobInfoResponseBodyDataJobConfigInfoTimeConfig `json:"TimeConfig,omitempty" xml:"TimeConfig,omitempty" type:"Struct"`
	// The extended fields.
	//
	// example:
	//
	// {"pageSize":5,"queueSize":10,"consumerSize":5,"dispatcherSize":5,"taskMaxAttempt":0,"taskAttemptInterval":0,"globalConsumerSize":1000,"taskDispatchMode":"push"}
	XAttrs *string `json:"XAttrs,omitempty" xml:"XAttrs,omitempty"`
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

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetJobId(v int64) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.JobId = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetJobMonitorInfo(v *GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.JobMonitorInfo = v
	return s
}

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetJobType(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.JobType = &v
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

func (s *GetJobInfoResponseBodyDataJobConfigInfo) SetXAttrs(v string) *GetJobInfoResponseBodyDataJobConfigInfo {
	s.XAttrs = &v
	return s
}

type GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfo struct {
	// The alert contact Information.
	ContactInfo []*GetJobInfoResponseBodyDataJobConfigInfoJobMonitorInfoContactInfo `json:"ContactInfo,omitempty" xml:"ContactInfo,omitempty" type:"Repeated"`
	// The configurations of the alerting features and the alert thresholds.
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
	// The webhook URL of the DingTalk chatbot.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=XXXXXX
	Ding *string `json:"Ding,omitempty" xml:"Ding,omitempty"`
	// The email address of the alert contact.
	//
	// example:
	//
	// user@demo.com
	UserMail *string `json:"UserMail,omitempty" xml:"UserMail,omitempty"`
	// The name of the alert contact.
	//
	// example:
	//
	// userA
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The mobile phone number of the alert contact.
	//
	// example:
	//
	// 1381111****
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
	// Indicates whether the Failure alarm switch was turned on. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	FailEnable *bool `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	// Indicates whether the No machine alarm available switch was turned on.
	//
	// example:
	//
	// true
	MissWorkerEnable *bool `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	// The method used to send alerts. Only Short Message Service (SMS) is supported.
	//
	// example:
	//
	// sms
	SendChannel *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	// The timeout threshold. Default value: 7200. Unit: seconds.
	//
	// example:
	//
	// 12300
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Indicates whether the Timeout alarm switch was turned on. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	TimeoutEnable *bool `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	// Indicates whether the Timeout termination switch was turned on. The switch is turned off by default.
	//
	// example:
	//
	// true
	TimeoutKillEnable *bool `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
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
	// The number of threads that were triggered by a single worker at a time. Default value: 5.
	//
	// example:
	//
	// 5
	ConsumerSize *int32 `json:"ConsumerSize,omitempty" xml:"ConsumerSize,omitempty"`
	// The number of task distribution threads. Default value: 5.
	//
	// example:
	//
	// 5
	DispatcherSize *int32 `json:"DispatcherSize,omitempty" xml:"DispatcherSize,omitempty"`
	// The number of tasks that were pulled by a parallel job at a time. Default value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The maximum number of tasks that can be queued. Default value: 10000.
	//
	// example:
	//
	// 10000
	QueueSize *int32 `json:"QueueSize,omitempty" xml:"QueueSize,omitempty"`
	// The interval at which the system retried to run the task after a task failure.
	//
	// example:
	//
	// 0
	TaskAttemptInterval *int32 `json:"TaskAttemptInterval,omitempty" xml:"TaskAttemptInterval,omitempty"`
	// The number of retries after a task failure.
	//
	// example:
	//
	// 0
	TaskMaxAttempt *int32 `json:"TaskMaxAttempt,omitempty" xml:"TaskMaxAttempt,omitempty"`
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
	// Custom calendar days specified if TimeType is set to **1*	- (cron).
	//
	// example:
	//
	// Business days
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The time offset specified if TimeType is set to **1*	- (cron). Unit: seconds.
	//
	// example:
	//
	// 0
	DataOffset *int32 `json:"DataOffset,omitempty" xml:"DataOffset,omitempty"`
	// The time expression specified based on the value of TimeType:
	//
	// 	- If TimeType is set to **100*	- (api), no time expression is required.
	//
	// 	- If TimeType is set to **3*	- (fix_rate), this parameter value indicates the specific and fixed frequency. For example, if the value is 30, the system triggers a job every 30 seconds.
	//
	// 	- If TimeType is set to **1*	- (cron), this parameter value indicates the standard CRON expression used to specify the time when to schedule the job.
	//
	// 	- If TimeType is set to **4*	- (second_delay), this parameter value indicates the fixed delay after which the job is triggered. Valid values: 1 to 60. Unit: seconds.
	//
	// example:
	//
	// 0 0/10 	- 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time type. Valid values:
	//
	// 	- **1**: cron
	//
	// 	- **3**: fix_rate
	//
	// 	- **4**: second_delay
	//
	// 	- **5**: one_time
	//
	// 	- **100**: api
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetJobInfoResponse) SetStatusCode(v int32) *GetJobInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobInfoResponse) SetBody(v *GetJobInfoResponseBody) *GetJobInfoResponse {
	s.Body = v
	return s
}

type GetJobInstanceRequest struct {
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11111111
	JobInstanceId *int64 `json:"JobInstanceId,omitempty" xml:"JobInstanceId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *GetJobInstanceRequest) SetRegionId(v string) *GetJobInstanceRequest {
	s.RegionId = &v
	return s
}

type GetJobInstanceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the job instance.
	Data *GetJobInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// jobid: 92583 not match groupId: testSchedulerx.defaultGroup
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The details of the job instance.
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
	// The data timestamp of the job instance.
	//
	// example:
	//
	// 2020-07-27 11:52:10
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The end time of the job execution.
	//
	// example:
	//
	// 2020-07-27 11:52:10
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The user who executes the job.
	//
	// example:
	//
	// A
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// The job instance ID.
	//
	// example:
	//
	// 11111111
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 92583
	JobId      *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobName    *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The progress of the job instance.
	//
	// example:
	//
	// complete
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The execution results of the job instance.
	//
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The time when the job was scheduled to run.
	//
	// example:
	//
	// 2020-07-27 11:52:10
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The start time of the job execution.
	//
	// example:
	//
	// 2020-07-27 11:52:10
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the job instance. Valid values:
	//
	// 	- **1**: The job instance is waiting for execution.
	//
	// 	- **3**: The job instance is running.
	//
	// 	- **4**: The job instance is successful.
	//
	// 	- **5**: The job instance failed.
	//
	// 	- **9**: The job instance is rejected.
	//
	// Enumeration class: com.alibaba.schedulerx.common.domain.InstanceStatus
	//
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The method that is used to specify the time when to schedule the job instance. Valid values:
	//
	// 	- **1**: cron
	//
	// 	- **3**: fix_rate
	//
	// 	- **4**: second_delay
	//
	// 	- **100**: api
	//
	// Enumeration class: com.alibaba.schedulerx.common.domain.TimeType
	//
	// example:
	//
	// 1
	TimeType *int32  `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	TraceId  *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// The trigger type of the job instance. Valid values:
	//
	// 	- **1**: The job instance was triggered at the scheduled time.
	//
	// 	- **2**: The job instance was triggered due to data update.
	//
	// 	- **3**: The job instance was triggered by an API call.
	//
	// 	- **4**: The job instance was triggered because it is manually rerun.
	//
	// 	- **5**: The job instance was triggered because the system automatically reruns the job instance upon a system exception, such as a database exception.
	//
	// Enumeration class: com.alibaba.schedulerx.common.domain.TriggerType
	//
	// example:
	//
	// 3
	TriggerType *int32 `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The endpoint of the triggered client. The value is in the IP address:Port number format.
	//
	// example:
	//
	// 192.168.0.0:16
	WorkAddr *string `json:"WorkAddr,omitempty" xml:"WorkAddr,omitempty"`
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

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetJobName(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.JobName = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetParameters(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.Parameters = &v
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

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetTraceId(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.TraceId = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetJobInstanceResponse) SetStatusCode(v int32) *GetJobInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobInstanceResponse) SetBody(v *GetJobInstanceResponseBody) *GetJobInstanceResponse {
	s.Body = v
	return s
}

type GetJobInstanceListRequest struct {
	// The end of the time range to query. Specify the time as a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1684202400000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The application group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the **Namespace*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time as a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1684116000000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The state of the job instance. Valid values:
	//
	// 1: The job instance is waiting for execution. 3: The job instance is running. 4: The job instance is successful. 5: The job instance fails. 9: The job instance is rejected. Enumeration class: com.alibaba.schedulerx.common.domain.InstanceStatus
	//
	// example:
	//
	// 5
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetJobInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobInstanceListRequest) GoString() string {
	return s.String()
}

func (s *GetJobInstanceListRequest) SetEndTimestamp(v int64) *GetJobInstanceListRequest {
	s.EndTimestamp = &v
	return s
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

func (s *GetJobInstanceListRequest) SetPageNum(v int32) *GetJobInstanceListRequest {
	s.PageNum = &v
	return s
}

func (s *GetJobInstanceListRequest) SetPageSize(v int32) *GetJobInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *GetJobInstanceListRequest) SetRegionId(v string) *GetJobInstanceListRequest {
	s.RegionId = &v
	return s
}

func (s *GetJobInstanceListRequest) SetStartTimestamp(v int64) *GetJobInstanceListRequest {
	s.StartTimestamp = &v
	return s
}

func (s *GetJobInstanceListRequest) SetStatus(v int32) *GetJobInstanceListRequest {
	s.Status = &v
	return s
}

type GetJobInstanceListResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the job instances.
	Data *GetJobInstanceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// jobid: 92583 not match groupId: testSchedulerx.defaultGroup
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The details of the job instance.
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
	// The data timestamp of the job instance.
	//
	// example:
	//
	// 2020-07-27 11:52:10
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The end time of the job execution.
	//
	// example:
	//
	// 2020-07-27 11:52:10
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The user who executes the job.
	//
	// example:
	//
	// A
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// The job instance ID.
	//
	// example:
	//
	// 11111111
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The progress of the job instance.
	//
	// example:
	//
	// complete
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The execution results of the job instance.
	//
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The time when the job was scheduled to run.
	//
	// example:
	//
	// 2020-07-27 11:52:10
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The start time of the job execution.
	//
	// example:
	//
	// 2020-07-27 11:52:10
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the job instance. Valid values:
	//
	// 	- **1**: The job instance is waiting for execution.
	//
	// 	- **3**: The job instance is running.
	//
	// 	- **4**: The job instance is successful.
	//
	// 	- **5**: The job instance failed.
	//
	// 	- **9**: The job instance is rejected.
	//
	// Enumeration class: com.alibaba.schedulerx.common.domain.InstanceStatus
	//
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The method that is used to specify the time when to schedule the job instance. Valid values:
	//
	// 	- **1**: cron
	//
	// 	- **3**: fix_rate
	//
	// 	- **4**: second_delay
	//
	// 	- **100**: api
	//
	// Enumeration class: com.alibaba.schedulerx.common.domain.TimeType
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The trigger type of the job instance. Valid values:
	//
	// 	- **1**: The job instance was triggered at the scheduled time.
	//
	// 	- **2**: The job instance was triggered due to data updates.
	//
	// 	- **3**: The job instance was triggered by an API call.
	//
	// 	- **4**: The job instance was triggered because it is manually rerun.
	//
	// 	- **5**: The job instance was triggered because the system automatically reruns the job instance upon a system exception, such as a database exception.
	//
	// Enumeration class: com.alibaba.schedulerx.common.domain.TriggerType
	//
	// example:
	//
	// 3
	TriggerType *int32 `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The endpoint of the triggered client. The value is in the IP address:Port number format.
	//
	// example:
	//
	// 192.168.0.0:16
	WorkAddr *string `json:"WorkAddr,omitempty" xml:"WorkAddr,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetJobInstanceListResponse) SetStatusCode(v int32) *GetJobInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobInstanceListResponse) SetBody(v *GetJobInstanceListResponseBody) *GetJobInstanceListResponse {
	s.Body = v
	return s
}

type GetLogRequest struct {
	// The time when the job stops running. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1675739484000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The application group ID. You can obtain the application group ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 123
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job instance ID.
	//
	// example:
	//
	// 123456
	JobInstanceId *string `json:"JobInstanceId,omitempty" xml:"JobInstanceId,omitempty"`
	// The keyword.
	//
	// example:
	//
	// ERROR
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of rows to return. The maximum number is 200.
	//
	// example:
	//
	// 50
	Line *int32 `json:"Line,omitempty" xml:"Line,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The number of offset rows. This parameter can be used for a paged query.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to reverse the order. By default, the order is reversed.
	//
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// The time when the job starts to run. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1675739364000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s GetLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogRequest) GoString() string {
	return s.String()
}

func (s *GetLogRequest) SetEndTimestamp(v int64) *GetLogRequest {
	s.EndTimestamp = &v
	return s
}

func (s *GetLogRequest) SetGroupId(v string) *GetLogRequest {
	s.GroupId = &v
	return s
}

func (s *GetLogRequest) SetJobId(v string) *GetLogRequest {
	s.JobId = &v
	return s
}

func (s *GetLogRequest) SetJobInstanceId(v string) *GetLogRequest {
	s.JobInstanceId = &v
	return s
}

func (s *GetLogRequest) SetKeyword(v string) *GetLogRequest {
	s.Keyword = &v
	return s
}

func (s *GetLogRequest) SetLine(v int32) *GetLogRequest {
	s.Line = &v
	return s
}

func (s *GetLogRequest) SetNamespace(v string) *GetLogRequest {
	s.Namespace = &v
	return s
}

func (s *GetLogRequest) SetNamespaceSource(v string) *GetLogRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GetLogRequest) SetOffset(v int32) *GetLogRequest {
	s.Offset = &v
	return s
}

func (s *GetLogRequest) SetRegionId(v string) *GetLogRequest {
	s.RegionId = &v
	return s
}

func (s *GetLogRequest) SetReverse(v bool) *GetLogRequest {
	s.Reverse = &v
	return s
}

func (s *GetLogRequest) SetStartTimestamp(v int64) *GetLogRequest {
	s.StartTimestamp = &v
	return s
}

type GetLogResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// jobid=xxx is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
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

func (s *GetLogResponseBody) SetData(v *GetLogResponseBodyData) *GetLogResponseBody {
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

type GetLogResponseBodyData struct {
	// The logs. The value is an array of strings.
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
}

func (s GetLogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLogResponseBodyData) SetLogs(v []*string) *GetLogResponseBodyData {
	s.Logs = v
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

type GetOverviewRequest struct {
	// example:
	//
	// 1684166400
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	MetricType *int32 `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// query
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1684166400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOverviewRequest) GoString() string {
	return s.String()
}

func (s *GetOverviewRequest) SetEndTime(v int64) *GetOverviewRequest {
	s.EndTime = &v
	return s
}

func (s *GetOverviewRequest) SetGroupId(v string) *GetOverviewRequest {
	s.GroupId = &v
	return s
}

func (s *GetOverviewRequest) SetMetricType(v int32) *GetOverviewRequest {
	s.MetricType = &v
	return s
}

func (s *GetOverviewRequest) SetNamespace(v string) *GetOverviewRequest {
	s.Namespace = &v
	return s
}

func (s *GetOverviewRequest) SetNamespaceSource(v string) *GetOverviewRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GetOverviewRequest) SetOperate(v string) *GetOverviewRequest {
	s.Operate = &v
	return s
}

func (s *GetOverviewRequest) SetRegionId(v string) *GetOverviewRequest {
	s.RegionId = &v
	return s
}

func (s *GetOverviewRequest) SetStartTime(v int64) *GetOverviewRequest {
	s.StartTime = &v
	return s
}

type GetOverviewResponseBody struct {
	// example:
	//
	// 200
	Code *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// No access permission for the namespace [***]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetOverviewResponseBody) SetCode(v int32) *GetOverviewResponseBody {
	s.Code = &v
	return s
}

func (s *GetOverviewResponseBody) SetData(v string) *GetOverviewResponseBody {
	s.Data = &v
	return s
}

func (s *GetOverviewResponseBody) SetMessage(v string) *GetOverviewResponseBody {
	s.Message = &v
	return s
}

func (s *GetOverviewResponseBody) SetRequestId(v string) *GetOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOverviewResponseBody) SetSuccess(v bool) *GetOverviewResponseBody {
	s.Success = &v
	return s
}

type GetOverviewResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOverviewResponse) GoString() string {
	return s.String()
}

func (s *GetOverviewResponse) SetHeaders(v map[string]*string) *GetOverviewResponse {
	s.Headers = v
	return s
}

func (s *GetOverviewResponse) SetStatusCode(v int32) *GetOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOverviewResponse) SetBody(v *GetOverviewResponseBody) *GetOverviewResponse {
	s.Body = v
	return s
}

type GetWorkFlowRequest struct {
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// hxm.test
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a06d5ea-f576-4326-842c-fb14ea043d8d
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace.
	//
	// example:
	//
	// source
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region information.
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
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
	// Error codes
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the workflow.
	Data *GetWorkFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error message
	//
	// example:
	//
	// workflow is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 45678xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the API call.
	//
	// example:
	//
	// true
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
	// The basic information of the workflow.
	WorkFlowInfo *GetWorkFlowResponseBodyDataWorkFlowInfo `json:"WorkFlowInfo,omitempty" xml:"WorkFlowInfo,omitempty" type:"Struct"`
	// The node information of the workflow.
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
	// The description of the workflow.
	//
	// example:
	//
	// my first workflow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the workflow.
	//
	// example:
	//
	// workflow_111
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the workflow.
	//
	// example:
	//
	// Successful
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time expression of the workflow.
	//
	// example:
	//
	// 0 0 2 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time type of the workflow.
	//
	// example:
	//
	// cron
	TimeType *string `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The ID of the workflow.
	//
	// example:
	//
	// 1234xxx
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
	// The workflow edges.
	Edges []*GetWorkFlowResponseBodyDataWorkFlowNodeInfoEdges `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Repeated"`
	// The list of workflow nodes.
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
	// The ID of the source job.
	//
	// example:
	//
	// 100
	Source *int64 `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the object job.
	//
	// example:
	//
	// 200
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
	// The ID of the job.
	//
	// example:
	//
	// 123456xxx
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// job_111
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The status of the job.
	//
	// example:
	//
	// 1
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetWorkFlowResponse) SetStatusCode(v int32) *GetWorkFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkFlowResponse) SetBody(v *GetWorkFlowResponseBody) *GetWorkFlowResponse {
	s.Body = v
	return s
}

type GetWorkerListRequest struct {
	// The ID of the permission group.
	//
	// This parameter is required.
	//
	// example:
	//
	// usercenter
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the namespace. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The job information.
	Data *GetWorkerListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// Cannot find product according to your domain.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- **true**: The call is successful.
	//
	// 	- **false**: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The worker information.
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
	// The IP address of the worker.
	//
	// example:
	//
	// 30.225.16.49
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The label of the worker.
	//
	// example:
	//
	// gray
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The port number of the worker.
	//
	// example:
	//
	// 60831
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The startup method of the worker.
	//
	// example:
	//
	// springboot
	Starter *string `json:"Starter,omitempty" xml:"Starter,omitempty"`
	// The version of the worker.
	//
	// example:
	//
	// 1.3.4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The address of the worker. The address is in the format of ${worker_id}@${worker_ip}:${worker_port}.
	//
	// example:
	//
	// 030225016049_11734_25917@30.225.16.49:60831
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkerListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetWorkerListResponse) SetStatusCode(v int32) *GetWorkerListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkerListResponse) SetBody(v *GetWorkerListResponseBody) *GetWorkerListResponse {
	s.Body = v
	return s
}

type GetWorkflowInstanceRequest struct {
	// The application group ID. You can obtain the ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The workflow instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	WfInstanceId *int64 `json:"WfInstanceId,omitempty" xml:"WfInstanceId,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetWorkflowInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceRequest) SetGroupId(v string) *GetWorkflowInstanceRequest {
	s.GroupId = &v
	return s
}

func (s *GetWorkflowInstanceRequest) SetNamespace(v string) *GetWorkflowInstanceRequest {
	s.Namespace = &v
	return s
}

func (s *GetWorkflowInstanceRequest) SetNamespaceSource(v string) *GetWorkflowInstanceRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GetWorkflowInstanceRequest) SetRegionId(v string) *GetWorkflowInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetWorkflowInstanceRequest) SetWfInstanceId(v int64) *GetWorkflowInstanceRequest {
	s.WfInstanceId = &v
	return s
}

func (s *GetWorkflowInstanceRequest) SetWorkflowId(v int64) *GetWorkflowInstanceRequest {
	s.WorkflowId = &v
	return s
}

type GetWorkflowInstanceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the workflow instance.
	Data *GetWorkflowInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// workflowId=xxx is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkflowInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBody) SetCode(v int32) *GetWorkflowInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetData(v *GetWorkflowInstanceResponseBodyData) *GetWorkflowInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetMessage(v string) *GetWorkflowInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetRequestId(v string) *GetWorkflowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetSuccess(v bool) *GetWorkflowInstanceResponseBody {
	s.Success = &v
	return s
}

type GetWorkflowInstanceResponseBodyData struct {
	// The directed acyclic graph (DAG) of the workflow instance, including job instances and dependencies.
	WfInstanceDag *GetWorkflowInstanceResponseBodyDataWfInstanceDag `json:"WfInstanceDag,omitempty" xml:"WfInstanceDag,omitempty" type:"Struct"`
	// The details of the workflow instance, including the state of the workflow instance, the time when the workflow instance started to run, the time when the workflow instance stopped running, the state of each job instance, and the dependencies between job instances.
	WfInstanceInfo *GetWorkflowInstanceResponseBodyDataWfInstanceInfo `json:"WfInstanceInfo,omitempty" xml:"WfInstanceInfo,omitempty" type:"Struct"`
}

func (s GetWorkflowInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBodyData) SetWfInstanceDag(v *GetWorkflowInstanceResponseBodyDataWfInstanceDag) *GetWorkflowInstanceResponseBodyData {
	s.WfInstanceDag = v
	return s
}

func (s *GetWorkflowInstanceResponseBodyData) SetWfInstanceInfo(v *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) *GetWorkflowInstanceResponseBodyData {
	s.WfInstanceInfo = v
	return s
}

type GetWorkflowInstanceResponseBodyDataWfInstanceDag struct {
	// The dependencies between job instances.
	Edges []*GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Repeated"`
	// The job instances.
	Nodes []*GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s GetWorkflowInstanceResponseBodyDataWfInstanceDag) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowInstanceResponseBodyDataWfInstanceDag) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDag) SetEdges(v []*GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges) *GetWorkflowInstanceResponseBodyDataWfInstanceDag {
	s.Edges = v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDag) SetNodes(v []*GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) *GetWorkflowInstanceResponseBodyDataWfInstanceDag {
	s.Nodes = v
	return s
}

type GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges struct {
	// The upstream job instance of the current job instance. The value 0 indicates that the upstream job instance is the root node.
	//
	// example:
	//
	// 24058798
	Source *int64 `json:"Source,omitempty" xml:"Source,omitempty"`
	// The downstream job instance of the current job instance.
	//
	// example:
	//
	// 24058796
	Target *int64 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges) SetSource(v int64) *GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges {
	s.Source = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges) SetTarget(v int64) *GetWorkflowInstanceResponseBodyDataWfInstanceDagEdges {
	s.Target = &v
	return s
}

type GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes struct {
	// The number of retries when the job instance failed.
	//
	// example:
	//
	// 0
	Attempt *int32 `json:"Attempt,omitempty" xml:"Attempt,omitempty"`
	// The data timestamp of the job instance.
	//
	// example:
	//
	// 2023-01-03 18:00:00
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The time when the job instance stopped running.
	//
	// example:
	//
	// 2023-01-03 18:00:21
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 1472
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job instance ID.
	//
	// example:
	//
	// 24058796
	JobInstanceId *int64  `json:"JobInstanceId,omitempty" xml:"JobInstanceId,omitempty"`
	JobName       *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The state of the job instance.
	//
	// example:
	//
	// code=200
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The time when the job instance was scheduled to run.
	//
	// example:
	//
	// 2023-01-03 18:00:03
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The time when the job instance started to run.
	//
	// example:
	//
	// 2023-01-03 18:00:03
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	// The server on which the job instance was run.
	//
	// example:
	//
	// 10.163.0.101:34027
	WorkAddr *string `json:"WorkAddr,omitempty" xml:"WorkAddr,omitempty"`
}

func (s GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetAttempt(v int32) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.Attempt = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetDataTime(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.DataTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetEndTime(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.EndTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetJobId(v int64) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.JobId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetJobInstanceId(v int64) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.JobInstanceId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetJobName(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.JobName = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetResult(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.Result = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetScheduleTime(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.ScheduleTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetStartTime(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.StartTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetStatus(v int32) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.Status = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes) SetWorkAddr(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceDagNodes {
	s.WorkAddr = &v
	return s
}

type GetWorkflowInstanceResponseBodyDataWfInstanceInfo struct {
	// The data timestamp of the workflow instance.
	//
	// example:
	//
	// 2023-01-03 18:00:00
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The time when the workflow instance stopped running.
	//
	// example:
	//
	// 2023-01-03 18:00:21
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the workflow instance was scheduled to run.
	//
	// example:
	//
	// 2023-01-03 18:00:00
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The time when the workflow instance started to run.
	//
	// example:
	//
	// 2023-01-03 18:00:01
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the workflow instance. Valid values:
	//
	// 	- 1: pending
	//
	// 	- 2: preparing
	//
	// 	- 3: running
	//
	// 	- 4: successful
	//
	// 	- 5: failed
	//
	// example:
	//
	// 5
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetWorkflowInstanceResponseBodyDataWfInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowInstanceResponseBodyDataWfInstanceInfo) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) SetDataTime(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceInfo {
	s.DataTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) SetEndTime(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceInfo {
	s.EndTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) SetScheduleTime(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceInfo {
	s.ScheduleTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) SetStartTime(v string) *GetWorkflowInstanceResponseBodyDataWfInstanceInfo {
	s.StartTime = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyDataWfInstanceInfo) SetStatus(v int32) *GetWorkflowInstanceResponseBodyDataWfInstanceInfo {
	s.Status = &v
	return s
}

type GetWorkflowInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkflowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkflowInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponse) SetHeaders(v map[string]*string) *GetWorkflowInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetWorkflowInstanceResponse) SetStatusCode(v int32) *GetWorkflowInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkflowInstanceResponse) SetBody(v *GetWorkflowInstanceResponseBody) *GetWorkflowInstanceResponse {
	s.Body = v
	return s
}

type GrantPermissionRequest struct {
	// Specifies whether to grant the permissions with the GRANT option. Valid values: -**true*	- -**false**
	//
	// example:
	//
	// false
	GrantOption *bool `json:"GrantOption,omitempty" xml:"GrantOption,omitempty"`
	// The application group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffcdf01
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 277641081920123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// lilei
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	// The HTTP status code.
	//
	// example:
	//
	// 400
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// Your request is denied as lack of ssl protect.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GrantPermissionResponse) SetStatusCode(v int32) *GrantPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantPermissionResponse) SetBody(v *GrantPermissionResponseBody) *GrantPermissionResponse {
	s.Body = v
	return s
}

type ListGroupsRequest struct {
	// The name of the application group.
	//
	// example:
	//
	// k8s-test
	AppGroupName *string `json:"AppGroupName,omitempty" xml:"AppGroupName,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the **Namespace*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1a72ecb1-b4cc-400a-a71b-20cdec9b****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) SetAppGroupName(v string) *ListGroupsRequest {
	s.AppGroupName = &v
	return s
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the applications.
	Data *ListGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB58A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The applications and their details.
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
	// The application group ID.
	//
	// example:
	//
	// 1
	AppGroupId *int64 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	// The AppKey for the application.
	//
	// example:
	//
	// a3G77O6NZxq/lyo1NC****==
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// DocTest
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application ID.
	//
	// example:
	//
	// DocTest.Group
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Version *int32  `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListGroupsResponseBodyDataAppGroups) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBodyDataAppGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyDataAppGroups) SetAppGroupId(v int64) *ListGroupsResponseBodyDataAppGroups {
	s.AppGroupId = &v
	return s
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

func (s *ListGroupsResponseBodyDataAppGroups) SetVersion(v int32) *ListGroupsResponseBodyDataAppGroups {
	s.Version = &v
	return s
}

type ListGroupsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListGroupsResponse) SetStatusCode(v int32) *ListGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsResponse) SetBody(v *ListGroupsResponseBody) *ListGroupsResponse {
	s.Body = v
	return s
}

type ListJobsRequest struct {
	// The ID of the application. You can obtain the application ID on the **Application Management*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// DocTest.Group
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// helloword
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The ID of the namespace. You can obtain the namespace ID on the **Namespace*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1a72ecb1-b4cc-400a-a71b-20cdec9b****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable the job. Valid values:
	//
	// 	- **0**: disables the job.
	//
	// 	- **1**: enables the job.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the jobs.
	Data *ListJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned if an error occurs.
	//
	// example:
	//
	// namespace can not find namespace: 1a72ecb1-b4cc-400a-a71b-20cdec9b****, namespaceSource:null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB58B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- **true**: The call is successful.
	//
	// 	- **false**: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The jobs and their details.
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
	// The interval at which the system retries to run the job after a job failure. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 30
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// The full path of the job interface class. This parameter is returned only for a Java job.
	//
	// example:
	//
	// com.alibaba.schedulerx.test.helloworld
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The script of the job. This parameter is returned only for a Python, Shell, or Go job.
	//
	// example:
	//
	// echo \\"hello\\"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description of the job.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution mode of the job. Valid values:
	//
	// 	- **standalone**: The job runs in standalone mode.
	//
	// 	- **broadcast**: The job runs in broadcast mode.
	//
	// 	- **parallel**: The job runs in parallel computing mode.
	//
	// 	- **grid**: The job runs in memory grid mode.
	//
	// 	- **batch**: The job runs in grid computing mode.
	//
	// 	- **shard**: The job runs in multipart mode.
	//
	// example:
	//
	// standalone
	ExecuteMode *string `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	// The full path to which a JAR package is uploaded in Object Storage Service (OSS).
	//
	// example:
	//
	// https:doc***.oss-cn-hangzhou.aliyuncs.com/sc-****-D-0.0.2-SNAPSHOT.jar
	JarUrl *string `json:"JarUrl,omitempty" xml:"JarUrl,omitempty"`
	// The ID of the job.
	//
	// example:
	//
	// 99341
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The monitoring information of the job.
	JobMonitorInfo *ListJobsResponseBodyDataJobsJobMonitorInfo `json:"JobMonitorInfo,omitempty" xml:"JobMonitorInfo,omitempty" type:"Struct"`
	// The type of the job.
	//
	// example:
	//
	// java
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The advanced configurations of the job. The parameters are returned only if the value of the ExecuteMode parameter is parallel, grid, or batch.
	MapTaskXAttrs *ListJobsResponseBodyDataJobsMapTaskXAttrs `json:"MapTaskXAttrs,omitempty" xml:"MapTaskXAttrs,omitempty" type:"Struct"`
	// The maximum number of retries after a job failure. This parameter is specified based on your business requirements. Default value: 0.
	//
	// example:
	//
	// 0
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// The maximum number of instances that can concurrently run for the job. Default value: 1. A value of 1 indicates that if the last triggered instance is running, the next instance is not triggered even if the scheduled point in time for running the instance is reached.
	//
	// example:
	//
	// 1
	MaxConcurrency *string `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// helloworld
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user-defined parameters. These parameters can be obtained when the job is running.
	//
	// example:
	//
	// test
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// Indicates whether the job is enabled. Valid values:
	//
	// 	- **1**: The job is enabled and can be triggered.
	//
	// 	- **0**: The job is disabled and cannot be triggered.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time configurations.
	TimeConfig *ListJobsResponseBodyDataJobsTimeConfig `json:"TimeConfig,omitempty" xml:"TimeConfig,omitempty" type:"Struct"`
	// The extended fields.
	//
	// example:
	//
	// {"pageSize":5,"queueSize":10,"consumerSize":5,"dispatcherSize":5,"taskMaxAttempt":0,"taskAttemptInterval":0,"globalConsumerSize":1000,"taskDispatchMode":"push"}
	XAttrs *string `json:"XAttrs,omitempty" xml:"XAttrs,omitempty"`
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

func (s *ListJobsResponseBodyDataJobs) SetJobType(v string) *ListJobsResponseBodyDataJobs {
	s.JobType = &v
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

func (s *ListJobsResponseBodyDataJobs) SetXAttrs(v string) *ListJobsResponseBodyDataJobs {
	s.XAttrs = &v
	return s
}

type ListJobsResponseBodyDataJobsJobMonitorInfo struct {
	// The contact information.
	ContactInfo []*ListJobsResponseBodyDataJobsJobMonitorInfoContactInfo `json:"ContactInfo,omitempty" xml:"ContactInfo,omitempty" type:"Repeated"`
	// The configurations of the alerting feature and the alert threshold.
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
	// The webhook URL of the DingTalk chatbot.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=**********
	Ding *string `json:"Ding,omitempty" xml:"Ding,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// user@mail.com
	UserMail *string `json:"UserMail,omitempty" xml:"UserMail,omitempty"`
	// The username.
	//
	// example:
	//
	// userA
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The mobile number of the user.
	//
	// example:
	//
	// 1381111****
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
	// Indicates whether the feature of generating an alert upon a failure is enabled. Valid values:
	//
	// 	- **true**: The feature is enabled.
	//
	// 	- **false**: The feature is disabled.
	//
	// example:
	//
	// true
	FailEnable *bool `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	// Indicates whether the feature of generating an alert when no machine is available for running the job is enabled.
	//
	// example:
	//
	// true
	MissWorkerEnable *bool `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	// The method that is used to send an alert notification. Only Short Message Service (SMS) is supported.
	//
	// example:
	//
	// sms
	SendChannel *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	// The timeout threshold. Unit: seconds. Default value: 7200.
	//
	// example:
	//
	// 12300
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Indicates whether the feature of generating an alert upon a timeout is enabled. Valid values:
	//
	// 	- **true**: The feature is enabled.
	//
	// 	- **false**: The feature is disabled.
	//
	// example:
	//
	// true
	TimeoutEnable *bool `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	// Indicates whether the feature of stopping job triggering upon a timeout is enabled. By default, the feature is disabled.
	//
	// 	- **true**: The feature is enabled.
	//
	// 	- **false**: The feature is disabled.
	//
	// example:
	//
	// false
	TimeoutKillEnable *bool `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
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
	// The number of threads that are triggered by a standalone job at a time. Default value: 5.
	//
	// example:
	//
	// 5
	ConsumerSize *int32 `json:"ConsumerSize,omitempty" xml:"ConsumerSize,omitempty"`
	// The number of task distribution threads. Default value: 5.
	//
	// example:
	//
	// 5
	DispatcherSize *int32 `json:"DispatcherSize,omitempty" xml:"DispatcherSize,omitempty"`
	// The number of tasks that are pulled by a parallel job at a time. Default value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The maximum number of task queues that can be cached. Default value: 10000.
	//
	// example:
	//
	// 10000
	QueueSize *int32 `json:"QueueSize,omitempty" xml:"QueueSize,omitempty"`
	// The interval at which the system retries to run the task after a task failure.
	//
	// example:
	//
	// 0
	TaskAttemptInterval *int32 `json:"TaskAttemptInterval,omitempty" xml:"TaskAttemptInterval,omitempty"`
	// The number of retries after a task failure.
	//
	// example:
	//
	// 0
	TaskMaxAttempt *int32 `json:"TaskMaxAttempt,omitempty" xml:"TaskMaxAttempt,omitempty"`
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
	// If the TimeType parameter is set to cron, you can specify custom calendar days.
	//
	// example:
	//
	// Business days
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The time offset if the TimeType parameter is set to cron. Unit: seconds.
	//
	// example:
	//
	// 0
	DataOffset *int32 `json:"DataOffset,omitempty" xml:"DataOffset,omitempty"`
	// The time expression. Valid values:
	//
	// 	- **api**: indicates that no time expression is used to specify the time when to schedule the job.
	//
	// 	- **fix_rate**: indicates that the job is triggered at a fixed frequency. For example, a value of 30 indicates that the job is triggered every 30 seconds.
	//
	// 	- **cron**: indicates that a standard CRON expression is used to specify the time when to schedule the job.
	//
	// 	- **second_delay**: indicates that the job is triggered after a fixed delay. Valid values: 1 to 60. Unit: seconds.
	//
	// example:
	//
	// 0 0/10 	- 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The method that is used to specify the time when to schedule the job. Valid values:
	//
	// 	- **1**: cron
	//
	// 	- **3**: fix_rate
	//
	// 	- **4**: second_delay
	//
	// 	- **100**: api
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
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

type ListNamespacesRequest struct {
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// schedulerx-dev
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListNamespacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListNamespacesRequest) SetNamespace(v string) *ListNamespacesRequest {
	s.Namespace = &v
	return s
}

func (s *ListNamespacesRequest) SetNamespaceName(v string) *ListNamespacesRequest {
	s.NamespaceName = &v
	return s
}

func (s *ListNamespacesRequest) SetRegionId(v string) *ListNamespacesRequest {
	s.RegionId = &v
	return s
}

type ListNamespacesResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the namespaces.
	Data *ListNamespacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB58C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The namespaces and their details.
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
	// The description of the namespace.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// doc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// 1a72ecb1-b4cc-400a-a71b-20cdec9b****
	UId *string `json:"UId,omitempty" xml:"UId,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListNamespacesResponse) SetStatusCode(v int32) *ListNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNamespacesResponse) SetBody(v *ListNamespacesResponseBody) *ListNamespacesResponse {
	s.Body = v
	return s
}

type ListWorkflowInstanceRequest struct {
	// The application group ID. You can obtain the ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ListWorkflowInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstanceRequest) SetGroupId(v string) *ListWorkflowInstanceRequest {
	s.GroupId = &v
	return s
}

func (s *ListWorkflowInstanceRequest) SetNamespace(v string) *ListWorkflowInstanceRequest {
	s.Namespace = &v
	return s
}

func (s *ListWorkflowInstanceRequest) SetNamespaceSource(v string) *ListWorkflowInstanceRequest {
	s.NamespaceSource = &v
	return s
}

func (s *ListWorkflowInstanceRequest) SetRegionId(v string) *ListWorkflowInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ListWorkflowInstanceRequest) SetWorkflowId(v string) *ListWorkflowInstanceRequest {
	s.WorkflowId = &v
	return s
}

type ListWorkflowInstanceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about workflow instances.
	Data *ListWorkflowInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// workflowId=xxx is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWorkflowInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstanceResponseBody) SetCode(v int32) *ListWorkflowInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkflowInstanceResponseBody) SetData(v *ListWorkflowInstanceResponseBodyData) *ListWorkflowInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkflowInstanceResponseBody) SetMessage(v string) *ListWorkflowInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkflowInstanceResponseBody) SetRequestId(v string) *ListWorkflowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowInstanceResponseBody) SetSuccess(v bool) *ListWorkflowInstanceResponseBody {
	s.Success = &v
	return s
}

type ListWorkflowInstanceResponseBodyData struct {
	// The workflow instances.
	WfInstanceInfos []*ListWorkflowInstanceResponseBodyDataWfInstanceInfos `json:"WfInstanceInfos,omitempty" xml:"WfInstanceInfos,omitempty" type:"Repeated"`
}

func (s ListWorkflowInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstanceResponseBodyData) SetWfInstanceInfos(v []*ListWorkflowInstanceResponseBodyDataWfInstanceInfos) *ListWorkflowInstanceResponseBodyData {
	s.WfInstanceInfos = v
	return s
}

type ListWorkflowInstanceResponseBodyDataWfInstanceInfos struct {
	// The data timestamp of the workflow instance.
	//
	// example:
	//
	// 2023-01-03 18:00:00
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The time when the workflow instance stopped running.
	//
	// example:
	//
	// 2023-01-03 18:00:21
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the workflow instance was scheduled to run.
	//
	// example:
	//
	// 2023-01-03 18:00:00
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The time when the workflow instance started to run.
	//
	// example:
	//
	// 2023-01-03 18:00:01
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the workflow instance. Valid values:
	//
	// 	- 1: pending
	//
	// 	- 2: preparing
	//
	// 	- 3: running
	//
	// 	- 4: successful
	//
	// 	- 5: failed
	//
	// example:
	//
	// 5
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The workflow instance ID.
	//
	// example:
	//
	// 123456
	WfInstanceId *int64 `json:"WfInstanceId,omitempty" xml:"WfInstanceId,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// 123
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ListWorkflowInstanceResponseBodyDataWfInstanceInfos) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowInstanceResponseBodyDataWfInstanceInfos) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) SetDataTime(v string) *ListWorkflowInstanceResponseBodyDataWfInstanceInfos {
	s.DataTime = &v
	return s
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) SetEndTime(v string) *ListWorkflowInstanceResponseBodyDataWfInstanceInfos {
	s.EndTime = &v
	return s
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) SetScheduleTime(v string) *ListWorkflowInstanceResponseBodyDataWfInstanceInfos {
	s.ScheduleTime = &v
	return s
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) SetStartTime(v string) *ListWorkflowInstanceResponseBodyDataWfInstanceInfos {
	s.StartTime = &v
	return s
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) SetStatus(v int32) *ListWorkflowInstanceResponseBodyDataWfInstanceInfos {
	s.Status = &v
	return s
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) SetWfInstanceId(v int64) *ListWorkflowInstanceResponseBodyDataWfInstanceInfos {
	s.WfInstanceId = &v
	return s
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) SetWorkflowId(v int64) *ListWorkflowInstanceResponseBodyDataWfInstanceInfos {
	s.WorkflowId = &v
	return s
}

type ListWorkflowInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstanceResponse) SetHeaders(v map[string]*string) *ListWorkflowInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowInstanceResponse) SetStatusCode(v int32) *ListWorkflowInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowInstanceResponse) SetBody(v *ListWorkflowInstanceResponseBody) *ListWorkflowInstanceResponse {
	s.Body = v
	return s
}

type RerunJobRequest struct {
	// The data timestamp of the job. Specify a string in the HH:mm:ss format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10:00:00
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The time when the job stops running. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1645718400000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The application group ID. You can obtain the application group ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time when the job starts to rerun. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1645459200000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s RerunJobRequest) String() string {
	return tea.Prettify(s)
}

func (s RerunJobRequest) GoString() string {
	return s.String()
}

func (s *RerunJobRequest) SetDataTime(v string) *RerunJobRequest {
	s.DataTime = &v
	return s
}

func (s *RerunJobRequest) SetEndDate(v int64) *RerunJobRequest {
	s.EndDate = &v
	return s
}

func (s *RerunJobRequest) SetGroupId(v string) *RerunJobRequest {
	s.GroupId = &v
	return s
}

func (s *RerunJobRequest) SetJobId(v int64) *RerunJobRequest {
	s.JobId = &v
	return s
}

func (s *RerunJobRequest) SetNamespace(v string) *RerunJobRequest {
	s.Namespace = &v
	return s
}

func (s *RerunJobRequest) SetNamespaceSource(v string) *RerunJobRequest {
	s.NamespaceSource = &v
	return s
}

func (s *RerunJobRequest) SetRegionId(v string) *RerunJobRequest {
	s.RegionId = &v
	return s
}

func (s *RerunJobRequest) SetStartDate(v int64) *RerunJobRequest {
	s.StartDate = &v
	return s
}

type RerunJobResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// jobId=xxx is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RerunJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RerunJobResponseBody) GoString() string {
	return s.String()
}

func (s *RerunJobResponseBody) SetCode(v int32) *RerunJobResponseBody {
	s.Code = &v
	return s
}

func (s *RerunJobResponseBody) SetMessage(v string) *RerunJobResponseBody {
	s.Message = &v
	return s
}

func (s *RerunJobResponseBody) SetRequestId(v string) *RerunJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *RerunJobResponseBody) SetSuccess(v bool) *RerunJobResponseBody {
	s.Success = &v
	return s
}

type RerunJobResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RerunJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RerunJobResponse) String() string {
	return tea.Prettify(s)
}

func (s RerunJobResponse) GoString() string {
	return s.String()
}

func (s *RerunJobResponse) SetHeaders(v map[string]*string) *RerunJobResponse {
	s.Headers = v
	return s
}

func (s *RerunJobResponse) SetStatusCode(v int32) *RerunJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RerunJobResponse) SetBody(v *RerunJobResponseBody) *RerunJobResponse {
	s.Body = v
	return s
}

type RetryJobInstanceRequest struct {
	// The application group ID. You can obtain the application group ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	JobInstanceId *int64 `json:"JobInstanceId,omitempty" xml:"JobInstanceId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RetryJobInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RetryJobInstanceRequest) GoString() string {
	return s.String()
}

func (s *RetryJobInstanceRequest) SetGroupId(v string) *RetryJobInstanceRequest {
	s.GroupId = &v
	return s
}

func (s *RetryJobInstanceRequest) SetJobId(v int64) *RetryJobInstanceRequest {
	s.JobId = &v
	return s
}

func (s *RetryJobInstanceRequest) SetJobInstanceId(v int64) *RetryJobInstanceRequest {
	s.JobInstanceId = &v
	return s
}

func (s *RetryJobInstanceRequest) SetNamespace(v string) *RetryJobInstanceRequest {
	s.Namespace = &v
	return s
}

func (s *RetryJobInstanceRequest) SetNamespaceSource(v string) *RetryJobInstanceRequest {
	s.NamespaceSource = &v
	return s
}

func (s *RetryJobInstanceRequest) SetRegionId(v string) *RetryJobInstanceRequest {
	s.RegionId = &v
	return s
}

type RetryJobInstanceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// jobId=xxx is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RetryJobInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetryJobInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RetryJobInstanceResponseBody) SetCode(v int32) *RetryJobInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *RetryJobInstanceResponseBody) SetMessage(v string) *RetryJobInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *RetryJobInstanceResponseBody) SetRequestId(v string) *RetryJobInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryJobInstanceResponseBody) SetSuccess(v bool) *RetryJobInstanceResponseBody {
	s.Success = &v
	return s
}

type RetryJobInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryJobInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryJobInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryJobInstanceResponse) GoString() string {
	return s.String()
}

func (s *RetryJobInstanceResponse) SetHeaders(v map[string]*string) *RetryJobInstanceResponse {
	s.Headers = v
	return s
}

func (s *RetryJobInstanceResponse) SetStatusCode(v int32) *RetryJobInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryJobInstanceResponse) SetBody(v *RetryJobInstanceResponseBody) *RetryJobInstanceResponse {
	s.Body = v
	return s
}

type RevokePermissionRequest struct {
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.defalutGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The unique identifier (UID) of the namespace. You can obtain the namespace UID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffcdf01
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The UID of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 277641081920123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The HTTP status code.
	//
	// example:
	//
	// 400
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// Your request is denied as lack of ssl protect.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *RevokePermissionResponse) SetStatusCode(v int32) *RevokePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokePermissionResponse) SetBody(v *RevokePermissionResponseBody) *RevokePermissionResponse {
	s.Body = v
	return s
}

type SetJobInstanceSuccessRequest struct {
	// The application group ID. You can obtain the application group ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	JobInstanceId *int64 `json:"JobInstanceId,omitempty" xml:"JobInstanceId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetJobInstanceSuccessRequest) String() string {
	return tea.Prettify(s)
}

func (s SetJobInstanceSuccessRequest) GoString() string {
	return s.String()
}

func (s *SetJobInstanceSuccessRequest) SetGroupId(v string) *SetJobInstanceSuccessRequest {
	s.GroupId = &v
	return s
}

func (s *SetJobInstanceSuccessRequest) SetJobId(v int64) *SetJobInstanceSuccessRequest {
	s.JobId = &v
	return s
}

func (s *SetJobInstanceSuccessRequest) SetJobInstanceId(v int64) *SetJobInstanceSuccessRequest {
	s.JobInstanceId = &v
	return s
}

func (s *SetJobInstanceSuccessRequest) SetNamespace(v string) *SetJobInstanceSuccessRequest {
	s.Namespace = &v
	return s
}

func (s *SetJobInstanceSuccessRequest) SetNamespaceSource(v string) *SetJobInstanceSuccessRequest {
	s.NamespaceSource = &v
	return s
}

func (s *SetJobInstanceSuccessRequest) SetRegionId(v string) *SetJobInstanceSuccessRequest {
	s.RegionId = &v
	return s
}

type SetJobInstanceSuccessResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// jobId=xxx is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetJobInstanceSuccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetJobInstanceSuccessResponseBody) GoString() string {
	return s.String()
}

func (s *SetJobInstanceSuccessResponseBody) SetCode(v int32) *SetJobInstanceSuccessResponseBody {
	s.Code = &v
	return s
}

func (s *SetJobInstanceSuccessResponseBody) SetMessage(v string) *SetJobInstanceSuccessResponseBody {
	s.Message = &v
	return s
}

func (s *SetJobInstanceSuccessResponseBody) SetRequestId(v string) *SetJobInstanceSuccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetJobInstanceSuccessResponseBody) SetSuccess(v bool) *SetJobInstanceSuccessResponseBody {
	s.Success = &v
	return s
}

type SetJobInstanceSuccessResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetJobInstanceSuccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetJobInstanceSuccessResponse) String() string {
	return tea.Prettify(s)
}

func (s SetJobInstanceSuccessResponse) GoString() string {
	return s.String()
}

func (s *SetJobInstanceSuccessResponse) SetHeaders(v map[string]*string) *SetJobInstanceSuccessResponse {
	s.Headers = v
	return s
}

func (s *SetJobInstanceSuccessResponse) SetStatusCode(v int32) *SetJobInstanceSuccessResponse {
	s.StatusCode = &v
	return s
}

func (s *SetJobInstanceSuccessResponse) SetBody(v *SetJobInstanceSuccessResponseBody) *SetJobInstanceSuccessResponse {
	s.Body = v
	return s
}

type SetWfInstanceSuccessRequest struct {
	// The application group ID. You can obtain the application group ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The workflow instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789
	WfInstanceId *int64 `json:"WfInstanceId,omitempty" xml:"WfInstanceId,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s SetWfInstanceSuccessRequest) String() string {
	return tea.Prettify(s)
}

func (s SetWfInstanceSuccessRequest) GoString() string {
	return s.String()
}

func (s *SetWfInstanceSuccessRequest) SetGroupId(v string) *SetWfInstanceSuccessRequest {
	s.GroupId = &v
	return s
}

func (s *SetWfInstanceSuccessRequest) SetNamespace(v string) *SetWfInstanceSuccessRequest {
	s.Namespace = &v
	return s
}

func (s *SetWfInstanceSuccessRequest) SetNamespaceSource(v string) *SetWfInstanceSuccessRequest {
	s.NamespaceSource = &v
	return s
}

func (s *SetWfInstanceSuccessRequest) SetRegionId(v string) *SetWfInstanceSuccessRequest {
	s.RegionId = &v
	return s
}

func (s *SetWfInstanceSuccessRequest) SetWfInstanceId(v int64) *SetWfInstanceSuccessRequest {
	s.WfInstanceId = &v
	return s
}

func (s *SetWfInstanceSuccessRequest) SetWorkflowId(v int64) *SetWfInstanceSuccessRequest {
	s.WorkflowId = &v
	return s
}

type SetWfInstanceSuccessResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// wofkflowId is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetWfInstanceSuccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetWfInstanceSuccessResponseBody) GoString() string {
	return s.String()
}

func (s *SetWfInstanceSuccessResponseBody) SetCode(v int32) *SetWfInstanceSuccessResponseBody {
	s.Code = &v
	return s
}

func (s *SetWfInstanceSuccessResponseBody) SetMessage(v string) *SetWfInstanceSuccessResponseBody {
	s.Message = &v
	return s
}

func (s *SetWfInstanceSuccessResponseBody) SetRequestId(v string) *SetWfInstanceSuccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetWfInstanceSuccessResponseBody) SetSuccess(v bool) *SetWfInstanceSuccessResponseBody {
	s.Success = &v
	return s
}

type SetWfInstanceSuccessResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetWfInstanceSuccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetWfInstanceSuccessResponse) String() string {
	return tea.Prettify(s)
}

func (s SetWfInstanceSuccessResponse) GoString() string {
	return s.String()
}

func (s *SetWfInstanceSuccessResponse) SetHeaders(v map[string]*string) *SetWfInstanceSuccessResponse {
	s.Headers = v
	return s
}

func (s *SetWfInstanceSuccessResponse) SetStatusCode(v int32) *SetWfInstanceSuccessResponse {
	s.StatusCode = &v
	return s
}

func (s *SetWfInstanceSuccessResponse) SetBody(v *SetWfInstanceSuccessResponseBody) *SetWfInstanceSuccessResponse {
	s.Body = v
	return s
}

type StopInstanceRequest struct {
	// The ID of the application. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the job instance in the running state.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11111111
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the job. You can obtain the ID of the job on the Task Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the namespace. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned only if an error occurs.
	//
	// example:
	//
	// Your request is denied as lack of ssl protect.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- **true**: The call is successful.
	//
	// 	- **false**: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *StopInstanceResponse) SetStatusCode(v int32) *StopInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopInstanceResponse) SetBody(v *StopInstanceResponseBody) *StopInstanceResponse {
	s.Body = v
	return s
}

type UpdateAppGroupRequest struct {
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateAppGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppGroupRequest) SetDescription(v string) *UpdateAppGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateAppGroupRequest) SetGroupId(v string) *UpdateAppGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateAppGroupRequest) SetMaxConcurrency(v int32) *UpdateAppGroupRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *UpdateAppGroupRequest) SetNamespace(v string) *UpdateAppGroupRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateAppGroupRequest) SetRegionId(v string) *UpdateAppGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAppGroupRequest) SetVersion(v int32) *UpdateAppGroupRequest {
	s.Version = &v
	return s
}

type UpdateAppGroupResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// app is not existed, groupId=xxxx, namesapce=xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppGroupResponseBody) SetCode(v int32) *UpdateAppGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAppGroupResponseBody) SetMessage(v string) *UpdateAppGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppGroupResponseBody) SetRequestId(v string) *UpdateAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppGroupResponseBody) SetSuccess(v bool) *UpdateAppGroupResponseBody {
	s.Success = &v
	return s
}

type UpdateAppGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppGroupResponse) SetHeaders(v map[string]*string) *UpdateAppGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppGroupResponse) SetStatusCode(v int32) *UpdateAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppGroupResponse) SetBody(v *UpdateAppGroupResponseBody) *UpdateAppGroupResponse {
	s.Body = v
	return s
}

type UpdateJobRequest struct {
	// The interval of retries after a job failure. Default value: 30. Unit: seconds.
	//
	// example:
	//
	// 30
	AttemptInterval *int32 `json:"AttemptInterval,omitempty" xml:"AttemptInterval,omitempty"`
	// If you set TimeType to 1 (cron), you can specify calendar days.
	//
	// example:
	//
	// Business days
	Calendar *string `json:"Calendar,omitempty" xml:"Calendar,omitempty"`
	// The full path of the job interface class.
	//
	// This field is available only when you set the job type to java. In this case, you must enter a full path.
	//
	// example:
	//
	// com.alibaba.test.helloworld
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The number of threads that are triggered by a single worker at a time. Default value: 5. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// 5
	ConsumerSize *int32 `json:"ConsumerSize,omitempty" xml:"ConsumerSize,omitempty"`
	// The information about the alert contact.
	ContactInfo []*UpdateJobRequestContactInfo `json:"ContactInfo,omitempty" xml:"ContactInfo,omitempty" type:"Repeated"`
	// The script content. This parameter is required when you set the job type to python, shell, go, or k8s.
	//
	// example:
	//
	// echo \\"hello\\"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// If you set TimeType to 1 (cron), you can specify a time offset. Unit: seconds.
	//
	// example:
	//
	// 2400
	DataOffset *int32 `json:"DataOffset,omitempty" xml:"DataOffset,omitempty"`
	// The job description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of task distribution threads. Default value: 5. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// 5
	DispatcherSize *int32 `json:"DispatcherSize,omitempty" xml:"DispatcherSize,omitempty"`
	// The execution mode of the job. Valid values:
	//
	// 	- **Stand-alone operation**
	//
	// 	- **Broadcast run**
	//
	// 	- **Visual MapReduce**
	//
	// 	- **MapReduce**
	//
	// 	- **Shard run**
	//
	// example:
	//
	// standalone
	ExecuteMode *string `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	// Specifies whether to turn on Failure alarm. If the switch is turned on, an alert will be generated upon a failure. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	FailEnable *bool `json:"FailEnable,omitempty" xml:"FailEnable,omitempty"`
	// The number of consecutive failures. An alert will be received if the number of consecutive failures reaches the value of this parameter.
	//
	// example:
	//
	// 1
	FailTimes *int32 `json:"FailTimes,omitempty" xml:"FailTimes,omitempty"`
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID. You can obtain the job ID on the Task Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The maximum number of retries after a job failure. This parameter is specified based on your business requirements.
	//
	// example:
	//
	// 0
	MaxAttempt *int32 `json:"MaxAttempt,omitempty" xml:"MaxAttempt,omitempty"`
	// The maximum number of concurrent instances. Default value: 1. The default value indicates that only one instance is allowed to run at a time. When an instance is running, another instance is not triggered even if the scheduled time for running the instance is reached.
	//
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// Specifies whether to turn on No machine alarm available. If the switch is turned on, an alert will be generated when no machine is available for running the job. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	MissWorkerEnable *bool `json:"MissWorkerEnable,omitempty" xml:"MissWorkerEnable,omitempty"`
	// The job name.
	//
	// example:
	//
	// helloword
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The namespace source. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The number of tasks that can be pulled at a time. Default value: 100. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The user-defined parameters that you can obtain when the job is running.
	//
	// example:
	//
	// test
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The maximum number of tasks that can be queued. Default value: 10000. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// 10000
	QueueSize *int32 `json:"QueueSize,omitempty" xml:"QueueSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The method that is used to send alerts. Only Short Message Service (SMS) is supported.
	//
	// example:
	//
	// sms
	SendChannel *string `json:"SendChannel,omitempty" xml:"SendChannel,omitempty"`
	// Specifies whether to turn on Successful notice. If the switch is turned on, a notice will be sent when a job succeeds.
	//
	// example:
	//
	// false
	SuccessNoticeEnable *bool `json:"SuccessNoticeEnable,omitempty" xml:"SuccessNoticeEnable,omitempty"`
	// The interval of retries after a task failure. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// 0
	TaskAttemptInterval *int32 `json:"TaskAttemptInterval,omitempty" xml:"TaskAttemptInterval,omitempty"`
	// The job mode. Valid values: push and pull. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// push
	TaskDispatchMode *string `json:"TaskDispatchMode,omitempty" xml:"TaskDispatchMode,omitempty"`
	// The number of retries after a task failure. This parameter is an advanced configuration item of the MapReduce job.
	//
	// example:
	//
	// 0
	TaskMaxAttempt *int32  `json:"TaskMaxAttempt,omitempty" xml:"TaskMaxAttempt,omitempty"`
	Template       *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The time expression. Specify the time expression based on the value of TimeType:
	//
	// 	- If you set TimeType to **1*	- (cron), specify this parameter to a standard CRON expression.
	//
	// 	- If you set TimeType to **100*	- (api), no time expression is required.
	//
	// 	- If you set TimeType to **3*	- (fixed_rate), specify this parameter to a fixed frequency in seconds. For example, if you set this parameter to 30, the system triggers a job every 30 seconds.
	//
	// 	- If you set TimeType to **4*	- (second_delay), specify this parameter to a fixed delay after which the job is triggered. Valid values: 1 to 60. Unit: seconds.
	//
	// example:
	//
	// 30
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The time type. Valid values:
	//
	// 	- **1**: cron
	//
	// 	- **3**: fix_rate
	//
	// 	- **4**: second_delay
	//
	// 	- **100**: api
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The timeout threshold. Unit: seconds.
	//
	// example:
	//
	// 7200
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Specifies whether to turn on Timeout alarm. If the switch is turned on, an alert will be generated upon a timeout. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	TimeoutEnable *bool `json:"TimeoutEnable,omitempty" xml:"TimeoutEnable,omitempty"`
	// Specifies whether to turn on Timeout termination. If the switch is turned on, the job will be terminated upon a timeout. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	TimeoutKillEnable *bool `json:"TimeoutKillEnable,omitempty" xml:"TimeoutKillEnable,omitempty"`
	// Time zone.
	//
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// example:
	//
	// {"resource":"shell","fileFormat":"unix","templateType":"customTemplate"}
	XAttrs *string `json:"XAttrs,omitempty" xml:"XAttrs,omitempty"`
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

func (s *UpdateJobRequest) SetFailTimes(v int32) *UpdateJobRequest {
	s.FailTimes = &v
	return s
}

func (s *UpdateJobRequest) SetGroupId(v string) *UpdateJobRequest {
	s.GroupId = &v
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

func (s *UpdateJobRequest) SetSuccessNoticeEnable(v bool) *UpdateJobRequest {
	s.SuccessNoticeEnable = &v
	return s
}

func (s *UpdateJobRequest) SetTaskAttemptInterval(v int32) *UpdateJobRequest {
	s.TaskAttemptInterval = &v
	return s
}

func (s *UpdateJobRequest) SetTaskDispatchMode(v string) *UpdateJobRequest {
	s.TaskDispatchMode = &v
	return s
}

func (s *UpdateJobRequest) SetTaskMaxAttempt(v int32) *UpdateJobRequest {
	s.TaskMaxAttempt = &v
	return s
}

func (s *UpdateJobRequest) SetTemplate(v string) *UpdateJobRequest {
	s.Template = &v
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

func (s *UpdateJobRequest) SetTimezone(v string) *UpdateJobRequest {
	s.Timezone = &v
	return s
}

func (s *UpdateJobRequest) SetXAttrs(v string) *UpdateJobRequest {
	s.XAttrs = &v
	return s
}

type UpdateJobRequestContactInfo struct {
	// The webhook URL of the DingTalk chatbot.[](https://open.dingtalk.com/document/org/application-types)
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=**********
	Ding *string `json:"Ding,omitempty" xml:"Ding,omitempty"`
	// The email address of the alert contact.
	//
	// example:
	//
	// test***@***.com
	UserMail *string `json:"UserMail,omitempty" xml:"UserMail,omitempty"`
	// The name of the alert contact.
	//
	// example:
	//
	// userA
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The mobile phone number of the alert contact.
	//
	// example:
	//
	// 1381111****
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information returned only if an error occurs.
	//
	// example:
	//
	// job type is java className can not be blank
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
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

type UpdateWorkflowRequest struct {
	// The description of the workflow.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application group ID. You can obtain the application group ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the workflow.
	//
	// example:
	//
	// helloworld
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time expression. You can set the time expression based on the selected method that is used to specify time.
	//
	// 	- If you set TimeType to cron, you need to enter a standard cron expression. Online verification is supported.
	//
	// 	- If you set TimeType to api, no time expression is required.
	//
	// example:
	//
	// 0 0/10 	- 	- 	- ?
	TimeExpression *string `json:"TimeExpression,omitempty" xml:"TimeExpression,omitempty"`
	// The method that is used to specify the time. Valid values:
	//
	// 	- 1: cron
	//
	// 	- 100: api
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s UpdateWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowRequest) SetDescription(v string) *UpdateWorkflowRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkflowRequest) SetGroupId(v string) *UpdateWorkflowRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateWorkflowRequest) SetName(v string) *UpdateWorkflowRequest {
	s.Name = &v
	return s
}

func (s *UpdateWorkflowRequest) SetNamespace(v string) *UpdateWorkflowRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateWorkflowRequest) SetNamespaceSource(v string) *UpdateWorkflowRequest {
	s.NamespaceSource = &v
	return s
}

func (s *UpdateWorkflowRequest) SetRegionId(v string) *UpdateWorkflowRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateWorkflowRequest) SetTimeExpression(v string) *UpdateWorkflowRequest {
	s.TimeExpression = &v
	return s
}

func (s *UpdateWorkflowRequest) SetTimeType(v int32) *UpdateWorkflowRequest {
	s.TimeType = &v
	return s
}

func (s *UpdateWorkflowRequest) SetWorkflowId(v string) *UpdateWorkflowRequest {
	s.WorkflowId = &v
	return s
}

type UpdateWorkflowResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// timetype is invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowResponseBody) SetCode(v int32) *UpdateWorkflowResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWorkflowResponseBody) SetMessage(v string) *UpdateWorkflowResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWorkflowResponseBody) SetRequestId(v string) *UpdateWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkflowResponseBody) SetSuccess(v bool) *UpdateWorkflowResponseBody {
	s.Success = &v
	return s
}

type UpdateWorkflowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowResponse) SetHeaders(v map[string]*string) *UpdateWorkflowResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkflowResponse) SetStatusCode(v int32) *UpdateWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkflowResponse) SetBody(v *UpdateWorkflowResponseBody) *UpdateWorkflowResponse {
	s.Body = v
	return s
}

type UpdateWorkflowDagRequest struct {
	// The directed acyclic graph (DAG) of the workflow, including the information about the nodes and the edges. Specify the value of this parameter in the JSON format.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"nodes":[{"id":2300691},{"id":10518089},{"id":1758851}],"edges":[{"source":10518089,"target":1758851},{"source":10518089,"target":2300691}]}
	DagJson *string `json:"DagJson,omitempty" xml:"DagJson,omitempty"`
	// The application group ID. You can obtain the application group ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s UpdateWorkflowDagRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowDagRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDagRequest) SetDagJson(v string) *UpdateWorkflowDagRequest {
	s.DagJson = &v
	return s
}

func (s *UpdateWorkflowDagRequest) SetGroupId(v string) *UpdateWorkflowDagRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateWorkflowDagRequest) SetNamespace(v string) *UpdateWorkflowDagRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateWorkflowDagRequest) SetNamespaceSource(v string) *UpdateWorkflowDagRequest {
	s.NamespaceSource = &v
	return s
}

func (s *UpdateWorkflowDagRequest) SetRegionId(v string) *UpdateWorkflowDagRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateWorkflowDagRequest) SetWorkflowId(v string) *UpdateWorkflowDagRequest {
	s.WorkflowId = &v
	return s
}

type UpdateWorkflowDagResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// workflowId=xxxx is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkflowDagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowDagResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDagResponseBody) SetCode(v int32) *UpdateWorkflowDagResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWorkflowDagResponseBody) SetMessage(v string) *UpdateWorkflowDagResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWorkflowDagResponseBody) SetRequestId(v string) *UpdateWorkflowDagResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkflowDagResponseBody) SetSuccess(v bool) *UpdateWorkflowDagResponseBody {
	s.Success = &v
	return s
}

type UpdateWorkflowDagResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkflowDagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkflowDagResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowDagResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDagResponse) SetHeaders(v map[string]*string) *UpdateWorkflowDagResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkflowDagResponse) SetStatusCode(v int32) *UpdateWorkflowDagResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkflowDagResponse) SetBody(v *UpdateWorkflowDagResponseBody) *UpdateWorkflowDagResponse {
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

// Summary:
//
// Deletes multiple jobs at a time.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
// ```xml
//
// <dependency>
//
//     <groupId>com.aliyun</groupId>
//
//     <artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//     <version>1.0.4</version>
//
// </dependency>
//
// ```
//
// @param request - BatchDeleteJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteJobsResponse
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

// Summary:
//
// Deletes multiple jobs at a time.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
// ```xml
//
// <dependency>
//
//     <groupId>com.aliyun</groupId>
//
//     <artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//     <version>1.0.4</version>
//
// </dependency>
//
// ```
//
// @param request - BatchDeleteJobsRequest
//
// @return BatchDeleteJobsResponse
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

// Summary:
//
// The additional information that is returned.
//
// @param request - BatchDeleteRouteStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteRouteStrategyResponse
func (client *Client) BatchDeleteRouteStrategyWithOptions(request *BatchDeleteRouteStrategyRequest, runtime *util.RuntimeOptions) (_result *BatchDeleteRouteStrategyResponse, _err error) {
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
		Action:      tea.String("BatchDeleteRouteStrategy"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteRouteStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The additional information that is returned.
//
// @param request - BatchDeleteRouteStrategyRequest
//
// @return BatchDeleteRouteStrategyResponse
func (client *Client) BatchDeleteRouteStrategy(request *BatchDeleteRouteStrategyRequest) (_result *BatchDeleteRouteStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeleteRouteStrategyResponse{}
	_body, _err := client.BatchDeleteRouteStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables multiple jobs at a time.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
// ```xml
//
// <dependency>
//
//     <groupId>com.aliyun</groupId>
//
//     <artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//     <version>1.0.4</version>
//
// </dependency>
//
// ```
//
// @param request - BatchDisableJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDisableJobsResponse
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

// Summary:
//
// Disables multiple jobs at a time.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
// ```xml
//
// <dependency>
//
//     <groupId>com.aliyun</groupId>
//
//     <artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//     <version>1.0.4</version>
//
// </dependency>
//
// ```
//
// @param request - BatchDisableJobsRequest
//
// @return BatchDisableJobsResponse
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

// Summary:
//
// Enables multiple jobs at a time.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
// ```xml
//
// <dependency>
//
//     <groupId>com.aliyun</groupId>
//
//     <artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//     <version>1.0.4</version>
//
// </dependency>
//
// ```
//
// @param request - BatchEnableJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchEnableJobsResponse
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

// Summary:
//
// Enables multiple jobs at a time.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
// ```xml
//
// <dependency>
//
//     <groupId>com.aliyun</groupId>
//
//     <artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//     <version>1.0.4</version>
//
// </dependency>
//
// ```
//
// @param request - BatchEnableJobsRequest
//
// @return BatchEnableJobsResponse
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

// Summary:
//
// Creates an application group. The AppKey is returned.
//
// @param request - CreateAppGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppGroupResponse
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

// Summary:
//
// Creates an application group. The AppKey is returned.
//
// @param request - CreateAppGroupRequest
//
// @return CreateAppGroupResponse
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

// Summary:
//
// Creates a job and obtains the job ID.
//
// @param request - CreateJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobResponse
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

	if !tea.BoolValue(util.IsUnset(request.FailTimes)) {
		body["FailTimes"] = request.FailTimes
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
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

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.SuccessNoticeEnable)) {
		body["SuccessNoticeEnable"] = request.SuccessNoticeEnable
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

	if !tea.BoolValue(util.IsUnset(request.Timezone)) {
		body["Timezone"] = request.Timezone
	}

	if !tea.BoolValue(util.IsUnset(request.XAttrs)) {
		body["XAttrs"] = request.XAttrs
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

// Summary:
//
// Creates a job and obtains the job ID.
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
// Creates a namespace.
//
// @param request - CreateNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespaceWithOptions(request *CreateNamespaceRequest, runtime *util.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNamespace"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a namespace.
//
// @param request - CreateNamespaceRequest
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespace(request *CreateNamespaceRequest) (_result *CreateNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CreateNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a routing policy.
//
// @param request - CreateRouteStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRouteStrategyResponse
func (client *Client) CreateRouteStrategyWithOptions(request *CreateRouteStrategyRequest, runtime *util.RuntimeOptions) (_result *CreateRouteStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.StrategyContent)) {
		query["StrategyContent"] = request.StrategyContent
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRouteStrategy"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRouteStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a routing policy.
//
// @param request - CreateRouteStrategyRequest
//
// @return CreateRouteStrategyResponse
func (client *Client) CreateRouteStrategy(request *CreateRouteStrategyRequest) (_result *CreateRouteStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRouteStrategyResponse{}
	_body, _err := client.CreateRouteStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workflow. By default, the created workflow is disabled. After you update the directed acyclic graph (DAG) of the workflow, you must manually or call the corresponding operation to enable the workflow. You can call this operation only in the professional edition.
//
// @param request - CreateWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkflowResponse
func (client *Client) CreateWorkflowWithOptions(request *CreateWorkflowRequest, runtime *util.RuntimeOptions) (_result *CreateWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxConcurrency)) {
		body["MaxConcurrency"] = request.MaxConcurrency
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
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
		Action:      tea.String("CreateWorkflow"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workflow. By default, the created workflow is disabled. After you update the directed acyclic graph (DAG) of the workflow, you must manually or call the corresponding operation to enable the workflow. You can call this operation only in the professional edition.
//
// @param request - CreateWorkflowRequest
//
// @return CreateWorkflowResponse
func (client *Client) CreateWorkflow(request *CreateWorkflowRequest) (_result *CreateWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWorkflowResponse{}
	_body, _err := client.CreateWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The additional information that is returned.
//
// @param request - DeleteAppGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppGroupResponse
func (client *Client) DeleteAppGroupWithOptions(request *DeleteAppGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteAppGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeleteJobs)) {
		query["DeleteJobs"] = request.DeleteJobs
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAppGroup"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The additional information that is returned.
//
// @param request - DeleteAppGroupRequest
//
// @return DeleteAppGroupResponse
func (client *Client) DeleteAppGroup(request *DeleteAppGroupRequest) (_result *DeleteAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppGroupResponse{}
	_body, _err := client.DeleteAppGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified job.
//
// @param request - DeleteJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobResponse
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

// Summary:
//
// Deletes a specified job.
//
// @param request - DeleteJobRequest
//
// @return DeleteJobResponse
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

// Summary:
//
// Deletes a routing policy.
//
// @param request - DeleteRouteStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRouteStrategyResponse
func (client *Client) DeleteRouteStrategyWithOptions(request *DeleteRouteStrategyRequest, runtime *util.RuntimeOptions) (_result *DeleteRouteStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRouteStrategy"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRouteStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a routing policy.
//
// @param request - DeleteRouteStrategyRequest
//
// @return DeleteRouteStrategyResponse
func (client *Client) DeleteRouteStrategy(request *DeleteRouteStrategyRequest) (_result *DeleteRouteStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRouteStrategyResponse{}
	_body, _err := client.DeleteRouteStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a workflow.
//
// @param request - DeleteWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkflowResponse
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

// Summary:
//
// Deletes a workflow.
//
// @param request - DeleteWorkflowRequest
//
// @return DeleteWorkflowResponse
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

// Summary:
//
// Returns available regions.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
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

// Summary:
//
// Returns available regions.
//
// @return DescribeRegionsResponse
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

// Summary:
//
// Designates machines.
//
// @param request - DesignateWorkersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DesignateWorkersResponse
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

// Summary:
//
// Designates machines.
//
// @param request - DesignateWorkersRequest
//
// @return DesignateWorkersResponse
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

// Summary:
//
// Disables a job.
//
// @param request - DisableJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableJobResponse
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

// Summary:
//
// Disables a job.
//
// @param request - DisableJobRequest
//
// @return DisableJobResponse
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

// Summary:
//
// Disables a specified workflow.
//
// @param request - DisableWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableWorkflowResponse
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

// Summary:
//
// Disables a specified workflow.
//
// @param request - DisableWorkflowRequest
//
// @return DisableWorkflowResponse
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

// Summary:
//
// Enables a job.
//
// @param request - EnableJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableJobResponse
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

// Summary:
//
// Enables a job.
//
// @param request - EnableJobRequest
//
// @return EnableJobResponse
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

// Summary:
//
// Enables a specified workflow.
//
// @param request - EnableWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableWorkflowResponse
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

// Summary:
//
// Enables a specified workflow.
//
// @param request - EnableWorkflowRequest
//
// @return EnableWorkflowResponse
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

// Summary:
//
// Triggers a job to immediately run once.
//
// Description:
//
// > The combination of the `JobID` and `ScheduleTime` parameters serves as a unique index. Therefore, after the ExecuteJob operation is called to run a job once, a sleep for one second is required before the ExecuteJob operation is called to run the job again. Otherwise, the job may fail.
//
// @param request - ExecuteJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteJobResponse
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

// Summary:
//
// Triggers a job to immediately run once.
//
// Description:
//
// > The combination of the `JobID` and `ScheduleTime` parameters serves as a unique index. Therefore, after the ExecuteJob operation is called to run a job once, a sleep for one second is required before the ExecuteJob operation is called to run the job again. Otherwise, the job may fail.
//
// @param request - ExecuteJobRequest
//
// @return ExecuteJobResponse
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

// Summary:
//
// Immediately triggers a workflow.
//
// @param request - ExecuteWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteWorkflowResponse
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

// Summary:
//
// Immediately triggers a workflow.
//
// @param request - ExecuteWorkflowRequest
//
// @return ExecuteWorkflowResponse
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

// Summary:
//
// The configuration of the alert. The value is a JSON string. For more information, see **the additional information about response parameters below this table**.
//
// @param request - GetAppGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppGroupResponse
func (client *Client) GetAppGroupWithOptions(request *GetAppGroupRequest, runtime *util.RuntimeOptions) (_result *GetAppGroupResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppGroup"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The configuration of the alert. The value is a JSON string. For more information, see **the additional information about response parameters below this table**.
//
// @param request - GetAppGroupRequest
//
// @return GetAppGroupResponse
func (client *Client) GetAppGroup(request *GetAppGroupRequest) (_result *GetAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppGroupResponse{}
	_body, _err := client.GetAppGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a job based on the job ID. In most cases, the obtained information is used to update jobs.
//
// @param request - GetJobInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobInfoResponse
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

// Summary:
//
// Queries the details of a job based on the job ID. In most cases, the obtained information is used to update jobs.
//
// @param request - GetJobInfoRequest
//
// @return GetJobInfoResponse
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

// Summary:
//
// Queries the information about a job instance. You can view the status and progress of the job instance.
//
// @param request - GetJobInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobInstanceResponse
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

// Summary:
//
// Queries the information about a job instance. You can view the status and progress of the job instance.
//
// @param request - GetJobInstanceRequest
//
// @return GetJobInstanceResponse
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

// Summary:
//
// Queries the most recent 10 execution instances of a job.
//
// @param request - GetJobInstanceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobInstanceListResponse
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

// Summary:
//
// Queries the most recent 10 execution instances of a job.
//
// @param request - GetJobInstanceListRequest
//
// @return GetJobInstanceListResponse
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

// Summary:
//
// Queries the operational logs of a job. You can call this operation only in the professional edition.
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
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operational logs of a job. You can call this operation only in the professional edition.
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
// 查询概览数据信息
//
// @param request - GetOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOverviewResponse
func (client *Client) GetOverviewWithOptions(request *GetOverviewRequest, runtime *util.RuntimeOptions) (_result *GetOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceSource)) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !tea.BoolValue(util.IsUnset(request.Operate)) {
		query["Operate"] = request.Operate
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOverview"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询概览数据信息
//
// @param request - GetOverviewRequest
//
// @return GetOverviewResponse
func (client *Client) GetOverview(request *GetOverviewRequest) (_result *GetOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOverviewResponse{}
	_body, _err := client.GetOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information about a workflow.
//
// @param request - GetWorkFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkFlowResponse
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

// Summary:
//
// Obtains the information about a workflow.
//
// @param request - GetWorkFlowRequest
//
// @return GetWorkFlowResponse
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

// Summary:
//
// Obtains the list of workers that are connected to an application.
//
// @param request - GetWorkerListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkerListResponse
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

// Summary:
//
// Obtains the list of workers that are connected to an application.
//
// @param request - GetWorkerListRequest
//
// @return GetWorkerListResponse
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

// Summary:
//
// Queries the details of a specified workflow instance, including the state of the workflow instance, the state of each job instance, and the dependencies between job instances. You can call this operation only in the professional edition.
//
// @param request - GetWorkflowInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowInstanceResponse
func (client *Client) GetWorkflowInstanceWithOptions(request *GetWorkflowInstanceRequest, runtime *util.RuntimeOptions) (_result *GetWorkflowInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkflowInstance"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkflowInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified workflow instance, including the state of the workflow instance, the state of each job instance, and the dependencies between job instances. You can call this operation only in the professional edition.
//
// @param request - GetWorkflowInstanceRequest
//
// @return GetWorkflowInstanceResponse
func (client *Client) GetWorkflowInstance(request *GetWorkflowInstanceRequest) (_result *GetWorkflowInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWorkflowInstanceResponse{}
	_body, _err := client.GetWorkflowInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants permissions to an application group.
//
// @param request - GrantPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantPermissionResponse
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

// Summary:
//
// Grants permissions to an application group.
//
// @param request - GrantPermissionRequest
//
// @return GrantPermissionResponse
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

// Summary:
//
// Queries a list of applications.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
// ```xml
//
// <dependency>
//
//     <groupId>com.aliyun</groupId>
//
//     <artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//     <version>1.0.5</version>
//
// </dependency>
//
// ```
//
// @param request - ListGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsResponse
func (client *Client) ListGroupsWithOptions(request *ListGroupsRequest, runtime *util.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppGroupName)) {
		query["AppGroupName"] = request.AppGroupName
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroups"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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

// Summary:
//
// Queries a list of applications.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
// ```xml
//
// <dependency>
//
//     <groupId>com.aliyun</groupId>
//
//     <artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//     <version>1.0.5</version>
//
// </dependency>
//
// ```
//
// @param request - ListGroupsRequest
//
// @return ListGroupsResponse
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

// Summary:
//
// Queries jobs.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
//     <dependency>
//
//           <groupId>com.aliyun</groupId>
//
//           <artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//           <version>1.0.5</version>
//
//     </dependency>
//
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

// Summary:
//
// Queries jobs.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
//     <dependency>
//
//           <groupId>com.aliyun</groupId>
//
//           <artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//           <version>1.0.5</version>
//
//     </dependency>
//
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
// Queries namespaces.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
//     <dependency>
//
//         <groupId>com.aliyun</groupId>
//
//         <artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//         <version>1.0.5</version>
//
//     </dependency>
//
// @param request - ListNamespacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNamespacesResponse
func (client *Client) ListNamespacesWithOptions(request *ListNamespacesRequest, runtime *util.RuntimeOptions) (_result *ListNamespacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNamespaces"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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

// Summary:
//
// Queries namespaces.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
//     <dependency>
//
//         <groupId>com.aliyun</groupId>
//
//         <artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//         <version>1.0.5</version>
//
//     </dependency>
//
// @param request - ListNamespacesRequest
//
// @return ListNamespacesResponse
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

// Summary:
//
// Queries the execution history of a workflow. You can call this operation only in the professional edition.
//
// @param request - ListWorkflowInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowInstanceResponse
func (client *Client) ListWorkflowInstanceWithOptions(request *ListWorkflowInstanceRequest, runtime *util.RuntimeOptions) (_result *ListWorkflowInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkflowInstance"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkflowInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution history of a workflow. You can call this operation only in the professional edition.
//
// @param request - ListWorkflowInstanceRequest
//
// @return ListWorkflowInstanceResponse
func (client *Client) ListWorkflowInstance(request *ListWorkflowInstanceRequest) (_result *ListWorkflowInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWorkflowInstanceResponse{}
	_body, _err := client.ListWorkflowInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reruns a job to obtain the historical data of the job. You can call this operation only in the professional edition.
//
// @param request - RerunJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RerunJobResponse
func (client *Client) RerunJobWithOptions(request *RerunJobRequest, runtime *util.RuntimeOptions) (_result *RerunJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataTime)) {
		body["DataTime"] = request.DataTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceSource)) {
		body["NamespaceSource"] = request.NamespaceSource
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RerunJob"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RerunJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reruns a job to obtain the historical data of the job. You can call this operation only in the professional edition.
//
// @param request - RerunJobRequest
//
// @return RerunJobResponse
func (client *Client) RerunJob(request *RerunJobRequest) (_result *RerunJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RerunJobResponse{}
	_body, _err := client.RerunJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reruns a successful or failed job instance. You can call this operation only in the professional edition.
//
// @param request - RetryJobInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryJobInstanceResponse
func (client *Client) RetryJobInstanceWithOptions(request *RetryJobInstanceRequest, runtime *util.RuntimeOptions) (_result *RetryJobInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.JobInstanceId)) {
		query["JobInstanceId"] = request.JobInstanceId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RetryJobInstance"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RetryJobInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reruns a successful or failed job instance. You can call this operation only in the professional edition.
//
// @param request - RetryJobInstanceRequest
//
// @return RetryJobInstanceResponse
func (client *Client) RetryJobInstance(request *RetryJobInstanceRequest) (_result *RetryJobInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetryJobInstanceResponse{}
	_body, _err := client.RetryJobInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the permissions that are granted to an Alibaba Cloud Resource Access Management (RAM) user.
//
// @param request - RevokePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokePermissionResponse
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

// Summary:
//
// Revokes the permissions that are granted to an Alibaba Cloud Resource Access Management (RAM) user.
//
// @param request - RevokePermissionRequest
//
// @return RevokePermissionResponse
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

// Summary:
//
// Forcibly sets the state of a job instance to successful. You can call this operation only in the professional edition.
//
// @param request - SetJobInstanceSuccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetJobInstanceSuccessResponse
func (client *Client) SetJobInstanceSuccessWithOptions(request *SetJobInstanceSuccessRequest, runtime *util.RuntimeOptions) (_result *SetJobInstanceSuccessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.JobInstanceId)) {
		query["JobInstanceId"] = request.JobInstanceId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetJobInstanceSuccess"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetJobInstanceSuccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Forcibly sets the state of a job instance to successful. You can call this operation only in the professional edition.
//
// @param request - SetJobInstanceSuccessRequest
//
// @return SetJobInstanceSuccessResponse
func (client *Client) SetJobInstanceSuccess(request *SetJobInstanceSuccessRequest) (_result *SetJobInstanceSuccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetJobInstanceSuccessResponse{}
	_body, _err := client.SetJobInstanceSuccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Forcibly sets the state of a workflow instance to successful. You can call this operation only in the professional edition.
//
// @param request - SetWfInstanceSuccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetWfInstanceSuccessResponse
func (client *Client) SetWfInstanceSuccessWithOptions(request *SetWfInstanceSuccessRequest, runtime *util.RuntimeOptions) (_result *SetWfInstanceSuccessResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.WfInstanceId)) {
		query["WfInstanceId"] = request.WfInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkflowId)) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetWfInstanceSuccess"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetWfInstanceSuccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Forcibly sets the state of a workflow instance to successful. You can call this operation only in the professional edition.
//
// @param request - SetWfInstanceSuccessRequest
//
// @return SetWfInstanceSuccessResponse
func (client *Client) SetWfInstanceSuccess(request *SetWfInstanceSuccessRequest) (_result *SetWfInstanceSuccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetWfInstanceSuccessResponse{}
	_body, _err := client.SetWfInstanceSuccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a job instance in the running state.
//
// @param request - StopInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInstanceResponse
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

// Summary:
//
// Stops a job instance in the running state.
//
// @param request - StopInstanceRequest
//
// @return StopInstanceResponse
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

// Summary:
//
// The additional information that is returned.
//
// @param request - UpdateAppGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppGroupResponse
func (client *Client) UpdateAppGroupWithOptions(request *UpdateAppGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateAppGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxConcurrency)) {
		query["MaxConcurrency"] = request.MaxConcurrency
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["Version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAppGroup"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The additional information that is returned.
//
// @param request - UpdateAppGroupRequest
//
// @return UpdateAppGroupResponse
func (client *Client) UpdateAppGroup(request *UpdateAppGroupRequest) (_result *UpdateAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppGroupResponse{}
	_body, _err := client.UpdateAppGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration information about a job. By default, you need to call the GetJobInfo operation to obtain the original configuration of the job before you call this operation to modify the configuration as required.
//
// @param request - UpdateJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateJobResponse
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

	if !tea.BoolValue(util.IsUnset(request.FailTimes)) {
		body["FailTimes"] = request.FailTimes
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
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

	if !tea.BoolValue(util.IsUnset(request.SuccessNoticeEnable)) {
		body["SuccessNoticeEnable"] = request.SuccessNoticeEnable
	}

	if !tea.BoolValue(util.IsUnset(request.TaskAttemptInterval)) {
		body["TaskAttemptInterval"] = request.TaskAttemptInterval
	}

	if !tea.BoolValue(util.IsUnset(request.TaskDispatchMode)) {
		body["TaskDispatchMode"] = request.TaskDispatchMode
	}

	if !tea.BoolValue(util.IsUnset(request.TaskMaxAttempt)) {
		body["TaskMaxAttempt"] = request.TaskMaxAttempt
	}

	if !tea.BoolValue(util.IsUnset(request.Template)) {
		body["Template"] = request.Template
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

	if !tea.BoolValue(util.IsUnset(request.Timezone)) {
		body["Timezone"] = request.Timezone
	}

	if !tea.BoolValue(util.IsUnset(request.XAttrs)) {
		body["XAttrs"] = request.XAttrs
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

// Summary:
//
// Updates the configuration information about a job. By default, you need to call the GetJobInfo operation to obtain the original configuration of the job before you call this operation to modify the configuration as required.
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

// Summary:
//
// Updates the basic information about a workflow. You can call this operation only in the professional edition.
//
// @param request - UpdateWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowResponse
func (client *Client) UpdateWorkflowWithOptions(request *UpdateWorkflowRequest, runtime *util.RuntimeOptions) (_result *UpdateWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
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

	if !tea.BoolValue(util.IsUnset(request.TimeExpression)) {
		body["TimeExpression"] = request.TimeExpression
	}

	if !tea.BoolValue(util.IsUnset(request.TimeType)) {
		body["TimeType"] = request.TimeType
	}

	if !tea.BoolValue(util.IsUnset(request.WorkflowId)) {
		body["WorkflowId"] = request.WorkflowId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkflow"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about a workflow. You can call this operation only in the professional edition.
//
// @param request - UpdateWorkflowRequest
//
// @return UpdateWorkflowResponse
func (client *Client) UpdateWorkflow(request *UpdateWorkflowRequest) (_result *UpdateWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWorkflowResponse{}
	_body, _err := client.UpdateWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the nodes and dependencies of a workflow. You can call this operation only in the professional edition.
//
// @param request - UpdateWorkflowDagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowDagResponse
func (client *Client) UpdateWorkflowDagWithOptions(request *UpdateWorkflowDagRequest, runtime *util.RuntimeOptions) (_result *UpdateWorkflowDagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DagJson)) {
		body["DagJson"] = request.DagJson
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceSource)) {
		body["NamespaceSource"] = request.NamespaceSource
	}

	if !tea.BoolValue(util.IsUnset(request.WorkflowId)) {
		body["WorkflowId"] = request.WorkflowId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkflowDag"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkflowDagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the nodes and dependencies of a workflow. You can call this operation only in the professional edition.
//
// @param request - UpdateWorkflowDagRequest
//
// @return UpdateWorkflowDagResponse
func (client *Client) UpdateWorkflowDag(request *UpdateWorkflowDagRequest) (_result *UpdateWorkflowDagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWorkflowDagResponse{}
	_body, _err := client.UpdateWorkflowDagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
