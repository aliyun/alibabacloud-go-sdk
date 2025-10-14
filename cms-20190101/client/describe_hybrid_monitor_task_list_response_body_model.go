// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridMonitorTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHybridMonitorTaskListResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeHybridMonitorTaskListResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeHybridMonitorTaskListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHybridMonitorTaskListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHybridMonitorTaskListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeHybridMonitorTaskListResponseBody
	GetSuccess() *string
	SetTaskList(v []*DescribeHybridMonitorTaskListResponseBodyTaskList) *DescribeHybridMonitorTaskListResponseBody
	GetTaskList() []*DescribeHybridMonitorTaskListResponseBodyTaskList
	SetTotal(v int32) *DescribeHybridMonitorTaskListResponseBody
	GetTotal() *int32
}

type DescribeHybridMonitorTaskListResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// 	- If the request was successful, the value `successful` is returned.
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 11145B76-566A-5D80-A8A3-FAD98D310079
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The metric import tasks.
	TaskList []*DescribeHybridMonitorTaskListResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeHybridMonitorTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorTaskListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHybridMonitorTaskListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHybridMonitorTaskListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHybridMonitorTaskListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHybridMonitorTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridMonitorTaskListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeHybridMonitorTaskListResponseBody) GetTaskList() []*DescribeHybridMonitorTaskListResponseBodyTaskList {
	return s.TaskList
}

func (s *DescribeHybridMonitorTaskListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeHybridMonitorTaskListResponseBody) SetCode(v string) *DescribeHybridMonitorTaskListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBody) SetMessage(v string) *DescribeHybridMonitorTaskListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBody) SetPageNumber(v int32) *DescribeHybridMonitorTaskListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBody) SetPageSize(v int32) *DescribeHybridMonitorTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBody) SetRequestId(v string) *DescribeHybridMonitorTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBody) SetSuccess(v string) *DescribeHybridMonitorTaskListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBody) SetTaskList(v []*DescribeHybridMonitorTaskListResponseBodyTaskList) *DescribeHybridMonitorTaskListResponseBody {
	s.TaskList = v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBody) SetTotal(v int32) *DescribeHybridMonitorTaskListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBody) Validate() error {
	if s.TaskList != nil {
		for _, item := range s.TaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHybridMonitorTaskListResponseBodyTaskList struct {
	// The tags of the metric import task.
	AttachLabels []*DescribeHybridMonitorTaskListResponseBodyTaskListAttachLabels `json:"AttachLabels,omitempty" xml:"AttachLabels,omitempty" type:"Repeated"`
	// The interval at which the CloudMonitor agent collects host monitoring data. Valid values:
	//
	// 	- 15
	//
	// 	- 30
	//
	// 	- 60
	//
	// Unit: seconds.
	//
	// example:
	//
	// 60
	CollectInterval *int32 `json:"CollectInterval,omitempty" xml:"CollectInterval,omitempty"`
	// The URL of the destination from which the CloudMonitor agent collects host monitoring data.
	//
	// example:
	//
	// http://localhost
	CollectTargetEndpoint *string `json:"CollectTargetEndpoint,omitempty" xml:"CollectTargetEndpoint,omitempty"`
	// The relative path from which the CloudMonitor agent collects monitoring data.
	//
	// example:
	//
	// /metrics
	CollectTargetPath *string `json:"CollectTargetPath,omitempty" xml:"CollectTargetPath,omitempty"`
	// The type of the monitoring data. Valid values: Spring, Tomcat, Nginx, Tengine, JVM, Redis, and MySQL.
	//
	// example:
	//
	// nginx
	CollectTargetType *string `json:"CollectTargetType,omitempty" xml:"CollectTargetType,omitempty"`
	// The timeout period during which the CloudMonitor agent collects host monitoring data. Valid values:
	//
	// 	- 0
	//
	// 	- 15
	//
	// 	- 30
	//
	// 	- 60
	//
	// Unit: seconds.
	//
	// example:
	//
	// 15
	CollectTimout *int32 `json:"CollectTimout,omitempty" xml:"CollectTimout,omitempty"`
	// The timestamp when the metric import task was created.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1639382496000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the metric import task.
	//
	// example:
	//
	// aliyun
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The additional information of the instance.
	//
	// example:
	//
	// test
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 3607****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instances where monitoring data is collected in batches.
	Instances []*string `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// C:\\UserData\\log\\*.Log
	LogFilePath *string `json:"LogFilePath,omitempty" xml:"LogFilePath,omitempty"`
	// The method that is used to aggregate on-premises log data.
	//
	// example:
	//
	// [{"metric": "metric1","filters": [{"column": "filed2","type": "include","values": ["222222"]}],"groupBys": [{"column": "filed3","name": "filed3"}],"calculates": [{"column": "field1","name": "avg","type": "avg"}]},{"metric": "metric2","filters": [{"column": "field1","type": "include","values": ["11111"]}],"groupBys": [{"column": "filed2","name": "filed2"}],"calculates": [{"column": "field1","name": "avg","type": "avg"}]}]
	LogProcess *string `json:"LogProcess,omitempty" xml:"LogProcess,omitempty"`
	// The sample on-premises log.
	//
	// example:
	//
	// {"logContent":"100.116.134.26 1119 - [13/Aug/2019:16:55:46 +0800] POST metrichub-cn-hongkong.aliyun.com /agent/metrics/putLines 200 0 \\"-\\" \\"127.0.0.1:7001\\" \\"Go-http-client/1.1\\" \\"-\\" \\"-\\" \\"0a98a21a15656865460656276e\\"","addData":{"field1":["1119","1119"],"filed2":["POSTx","POST"],"filed3":["true","200"]}}
	LogSample *string `json:"LogSample,omitempty" xml:"LogSample,omitempty"`
	// The result that is returned after on-premises log data is split based on the specified matching mode.
	//
	// > The matching modes of on-premises log data include full regex mode, delimiter mode, and JSON mode.
	//
	// example:
	//
	// {"type": "regex","regex": "\\\\d+\\\\.\\\\d+\\\\.\\\\d+\\\\.\\\\d+\\\\s+(\\\\d+)\\\\s+\\\\S+\\\\s+\[\\\\d+/\\\\S+/\\\\d+:\\\\d+:\\\\d+:\\\\d+\\\\s+\\\\+\\\\d+\\\\]\\\\s+(\\\\S+)\\\\s+\\\\S+\\\\s+/\\\\S+/\\\\S+/\\\\S+\\\\s+(\\\\d+)","columns": [{"name": "field1"},{"name": "filed2","translate": {"default": "-","mappings": [{"from": "(\\\\w+)","to": "$1x","type": "regex"}]}},{"name": "filed3","translate": {"default": "-","mappings": [{"from": "NumberRange(100,300)","to": "true","type": "function"}]}}]}
	LogSplit *string `json:"LogSplit,omitempty" xml:"LogSplit,omitempty"`
	// The conditions that are used to match the instances in the application group.
	MatchExpress []*DescribeHybridMonitorTaskListResponseBodyTaskListMatchExpress `json:"MatchExpress,omitempty" xml:"MatchExpress,omitempty" type:"Repeated"`
	// The relationship between the conditions that are used to filter metric import tasks. Valid values:
	//
	// 	- or
	//
	// 	- and
	//
	// example:
	//
	// or
	MatchExpressRelation *string `json:"MatchExpressRelation,omitempty" xml:"MatchExpressRelation,omitempty"`
	// The namespace to which the host belongs.
	//
	// example:
	//
	// aliyun
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The network type of the host. Valid values:
	//
	// 	- `vpc`
	//
	// 	- `Internet`
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The configurations of the logs that are imported from Log Service.
	//
	// example:
	//
	// {"express": [],"filter": {"filters": [{"key": "task_type","operator": "=","value": "1"}]},"groupby": [{"alias": "isp","key": "isp","sqlKey": "t.`isp`","valueKey": "isp"}],"interval": 60,"labels": [{"name": "__cms_app__","type": 0,"value": "sitemonitor"}],"statistics": [{"alias": "http_dns_time_avg","function": "avg","key": "http_dns_time"}]}
	SLSProcess *string `json:"SLSProcess,omitempty" xml:"SLSProcess,omitempty"`
	// The configurations of the logs that are imported from Log Service.
	//
	// > This parameter is returned only if the `TaskType` parameter is set to `aliyun_sls`.
	SLSProcessConfig *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig `json:"SLSProcessConfig,omitempty" xml:"SLSProcessConfig,omitempty" type:"Struct"`
	// The ID of the member account.
	//
	// > This parameter is displayed only when you call this operation by using a management account.
	//
	// example:
	//
	// 120886317861****
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
	// The ID of the metric import task.
	//
	// example:
	//
	// 36****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the metric import task.
	//
	// example:
	//
	// aliyun_task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the metric import task. Valid values:
	//
	// 	- aliyun_fc: metric import tasks for Alibaba Cloud services
	//
	// 	- aliyun_sls: metrics for logs imported from Log Service
	//
	// example:
	//
	// aliyun_sls
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The region where the host resides.
	//
	// example:
	//
	// cn-hangzhou
	UploadRegion *string `json:"UploadRegion,omitempty" xml:"UploadRegion,omitempty"`
	// The configuration file of the Alibaba Cloud service that you want to monitor by using Hybrid Cloud Monitoring.
	//
	// 	- namespace: the namespace of the Alibaba Cloud service.
	//
	// 	- metric_list: the metrics of the Alibaba Cloud service.
	//
	// example:
	//
	// products:- namespace: acs_ecs_dashboard metric_info: - metric_list: - cpu_total
	YARMConfig *string `json:"YARMConfig,omitempty" xml:"YARMConfig,omitempty"`
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetAttachLabels() []*DescribeHybridMonitorTaskListResponseBodyTaskListAttachLabels {
	return s.AttachLabels
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetCollectInterval() *int32 {
	return s.CollectInterval
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetCollectTargetEndpoint() *string {
	return s.CollectTargetEndpoint
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetCollectTargetPath() *string {
	return s.CollectTargetPath
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetCollectTargetType() *string {
	return s.CollectTargetType
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetCollectTimout() *int32 {
	return s.CollectTimout
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetDescription() *string {
	return s.Description
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetInstances() []*string {
	return s.Instances
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetLogFilePath() *string {
	return s.LogFilePath
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetLogProcess() *string {
	return s.LogProcess
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetLogSample() *string {
	return s.LogSample
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetLogSplit() *string {
	return s.LogSplit
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetMatchExpress() []*DescribeHybridMonitorTaskListResponseBodyTaskListMatchExpress {
	return s.MatchExpress
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetMatchExpressRelation() *string {
	return s.MatchExpressRelation
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetSLSProcess() *string {
	return s.SLSProcess
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetSLSProcessConfig() *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig {
	return s.SLSProcessConfig
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetTargetUserId() *string {
	return s.TargetUserId
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetUploadRegion() *string {
	return s.UploadRegion
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) GetYARMConfig() *string {
	return s.YARMConfig
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetAttachLabels(v []*DescribeHybridMonitorTaskListResponseBodyTaskListAttachLabels) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.AttachLabels = v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetCollectInterval(v int32) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.CollectInterval = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetCollectTargetEndpoint(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.CollectTargetEndpoint = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetCollectTargetPath(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.CollectTargetPath = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetCollectTargetType(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.CollectTargetType = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetCollectTimout(v int32) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.CollectTimout = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetCreateTime(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.CreateTime = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetDescription(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.Description = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetExtraInfo(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.ExtraInfo = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetGroupId(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.GroupId = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetInstances(v []*string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.Instances = v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetLogFilePath(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.LogFilePath = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetLogProcess(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.LogProcess = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetLogSample(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.LogSample = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetLogSplit(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.LogSplit = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetMatchExpress(v []*DescribeHybridMonitorTaskListResponseBodyTaskListMatchExpress) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.MatchExpress = v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetMatchExpressRelation(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.MatchExpressRelation = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetNamespace(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.Namespace = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetNetworkType(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.NetworkType = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetSLSProcess(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.SLSProcess = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetSLSProcessConfig(v *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.SLSProcessConfig = v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetTargetUserId(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.TargetUserId = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetTaskId(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetTaskName(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.TaskName = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetTaskType(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.TaskType = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetUploadRegion(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.UploadRegion = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) SetYARMConfig(v string) *DescribeHybridMonitorTaskListResponseBodyTaskList {
	s.YARMConfig = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskList) Validate() error {
	if s.AttachLabels != nil {
		for _, item := range s.AttachLabels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MatchExpress != nil {
		for _, item := range s.MatchExpress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SLSProcessConfig != nil {
		if err := s.SLSProcessConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHybridMonitorTaskListResponseBodyTaskListAttachLabels struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskListAttachLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskListAttachLabels) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListAttachLabels) GetName() *string {
	return s.Name
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListAttachLabels) GetValue() *string {
	return s.Value
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListAttachLabels) SetName(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListAttachLabels {
	s.Name = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListAttachLabels) SetValue(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListAttachLabels {
	s.Value = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListAttachLabels) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridMonitorTaskListResponseBodyTaskListMatchExpress struct {
	// The method that is used to match the instance name. Valid values:
	//
	// 	- startWith: starts with a prefix
	//
	// 	- endWith: ends with a suffix
	//
	// 	- all: matches all
	//
	// 	- equals: equals
	//
	// 	- contains: contains
	//
	// 	- notContains: does not contain
	//
	// example:
	//
	// all
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The instance name.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The keyword that corresponds to the instance name.
	//
	// example:
	//
	// ECS_instance1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskListMatchExpress) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskListMatchExpress) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListMatchExpress) GetFunction() *string {
	return s.Function
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListMatchExpress) GetName() *string {
	return s.Name
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListMatchExpress) GetValue() *string {
	return s.Value
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListMatchExpress) SetFunction(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListMatchExpress {
	s.Function = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListMatchExpress) SetName(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListMatchExpress {
	s.Name = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListMatchExpress) SetValue(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListMatchExpress {
	s.Value = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListMatchExpress) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig struct {
	// The extended fields that indicate the results of basic operations that are performed on aggregation results.
	Express []*DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigExpress `json:"Express,omitempty" xml:"Express,omitempty" type:"Repeated"`
	// The conditions that are used to filter logs imported from Log Service.
	Filter *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The dimensions based on which data is aggregated. This parameter is equivalent to the GROUP BY clause in SQL.
	GroupBy []*DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigGroupBy `json:"GroupBy,omitempty" xml:"GroupBy,omitempty" type:"Repeated"`
	// The methods that are used to aggregate logs imported from Log Service.
	Statistics []*DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig) GetExpress() []*DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigExpress {
	return s.Express
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig) GetFilter() *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilter {
	return s.Filter
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig) GetGroupBy() []*DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigGroupBy {
	return s.GroupBy
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig) GetStatistics() []*DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics {
	return s.Statistics
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig) SetExpress(v []*DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigExpress) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig {
	s.Express = v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig) SetFilter(v *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilter) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig {
	s.Filter = v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig) SetGroupBy(v []*DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigGroupBy) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig {
	s.GroupBy = v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig) SetStatistics(v []*DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig {
	s.Statistics = v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfig) Validate() error {
	if s.Express != nil {
		for _, item := range s.Express {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	if s.GroupBy != nil {
		for _, item := range s.GroupBy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Statistics != nil {
		for _, item := range s.Statistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigExpress struct {
	// The alias of the extended field that indicates the result of basic operations that are performed on aggregation results.
	//
	// example:
	//
	// SuccRate
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The extended field that indicates the result of basic operations that are performed on aggregation results.
	//
	// example:
	//
	// success_count
	Express *string `json:"Express,omitempty" xml:"Express,omitempty"`
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigExpress) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigExpress) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigExpress) GetAlias() *string {
	return s.Alias
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigExpress) GetExpress() *string {
	return s.Express
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigExpress) SetAlias(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigExpress {
	s.Alias = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigExpress) SetExpress(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigExpress {
	s.Express = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigExpress) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilter struct {
	// The conditions that are used to filter logs imported from Log Service.
	Filters []*DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilterFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The relationship between multiple filter conditions. Valid values:
	//
	// 	- and (default): Logs are processed only if all filter conditions are met.
	//
	// 	- or: Logs are processed if one of the filter conditions is met.
	//
	// example:
	//
	// and
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilter) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilter) GetFilters() []*DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilterFilters {
	return s.Filters
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilter) GetRelation() *string {
	return s.Relation
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilter) SetFilters(v []*DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilterFilters) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilter {
	s.Filters = v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilter) SetRelation(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilter {
	s.Relation = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilter) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilterFilters struct {
	// The method that is used to filter logs imported from Log Service. Valid values:
	//
	// 	- `contain`: contains
	//
	// 	- `notContain`: does not contain
	//
	// 	- `>`: greater than
	//
	// 	- `<`: less than
	//
	// 	- `=`: equal to
	//
	// 	- `! =`: not equal to
	//
	// 	- `>=`: greater than or equal to
	//
	// 	- `<=`: less than or equal to
	//
	// example:
	//
	// =
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The name of the key that is used to filter logs imported from Log Service.
	//
	// example:
	//
	// host
	SLSKeyName *string `json:"SLSKeyName,omitempty" xml:"SLSKeyName,omitempty"`
	// The value of the key that is used to filter logs imported from Log Service.
	//
	// example:
	//
	// 200
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilterFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilterFilters) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilterFilters) GetOperator() *string {
	return s.Operator
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilterFilters) GetSLSKeyName() *string {
	return s.SLSKeyName
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilterFilters) GetValue() *string {
	return s.Value
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilterFilters) SetOperator(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilterFilters {
	s.Operator = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilterFilters) SetSLSKeyName(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilterFilters {
	s.SLSKeyName = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilterFilters) SetValue(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilterFilters {
	s.Value = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigFilterFilters) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigGroupBy struct {
	// The alias of the aggregation result.
	//
	// example:
	//
	// ApiResult
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The name of the key that is used to aggregate logs imported from Log Service.
	//
	// example:
	//
	// code
	SLSKeyName *string `json:"SLSKeyName,omitempty" xml:"SLSKeyName,omitempty"`
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigGroupBy) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigGroupBy) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigGroupBy) GetAlias() *string {
	return s.Alias
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigGroupBy) GetSLSKeyName() *string {
	return s.SLSKeyName
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigGroupBy) SetAlias(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigGroupBy {
	s.Alias = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigGroupBy) SetSLSKeyName(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigGroupBy {
	s.SLSKeyName = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigGroupBy) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics struct {
	// The alias of the aggregation result.
	//
	// example:
	//
	// level_count
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The function that is used to aggregate log data within a statistical period. Valid values:
	//
	// 	- count: counts the number.
	//
	// 	- sum: calculates the total value.
	//
	// 	- avg: calculates the average value.
	//
	// 	- max: calculates the maximum value.
	//
	// 	- min: calculates the minimum value.
	//
	// 	- value: collects samples within the statistical period.
	//
	// 	- countps: calculates the average number of the specified field per second by using the following formula: Counted number of the specified field/Total number of seconds within the statistical period.
	//
	// 	- sumps: calculates the average number of the specified field per second by using the following formula: Total value of the specified field/Total number of seconds within the statistical period.
	//
	// 	- distinct: counts the number of logs where the specified field appears within the statistical period.
	//
	// 	- distribution: counts the number of logs that meet a specified condition within the statistical period.
	//
	// 	- percentile: sorts the values of the specified field in ascending order, and then returns the value that is at the specified percentile within the statistical period. Example: P50.
	//
	// example:
	//
	// count
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The value of the function that is used to aggregate logs imported from Log Service.
	//
	// 	- If the `Function` parameter is set to `distribution`, this parameter indicates the lower limit of the statistical interval. For example, 200 indicates that the number of HTTP requests whose status code is 2XX is calculated.
	//
	// 	- If the `Function` parameter is set to `percentile`, this parameter specifies the percentile at which the expected value is. For example, 0.5 specifies P50.
	//
	// example:
	//
	// 200
	Parameter1 *string `json:"Parameter1,omitempty" xml:"Parameter1,omitempty"`
	// The value of the function that is used to aggregate logs imported from Log Service.
	//
	// > This parameter is returned only if the `Function` parameter is set to `distribution`. This parameter indicates the upper limit of the statistical interval. For example, 299 indicates that the number of HTTP requests whose status code is 2XX is calculated.
	//
	// example:
	//
	// 299
	Parameter2 *string `json:"Parameter2,omitempty" xml:"Parameter2,omitempty"`
	// The name of the key that is used to aggregate logs imported from Log Service.
	//
	// example:
	//
	// name
	SLSKeyName *string `json:"SLSKeyName,omitempty" xml:"SLSKeyName,omitempty"`
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics) GetAlias() *string {
	return s.Alias
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics) GetFunction() *string {
	return s.Function
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics) GetParameter1() *string {
	return s.Parameter1
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics) GetParameter2() *string {
	return s.Parameter2
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics) GetSLSKeyName() *string {
	return s.SLSKeyName
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics) SetAlias(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics {
	s.Alias = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics) SetFunction(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics {
	s.Function = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics) SetParameter1(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics {
	s.Parameter1 = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics) SetParameter2(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics {
	s.Parameter2 = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics) SetSLSKeyName(v string) *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics {
	s.SLSKeyName = &v
	return s
}

func (s *DescribeHybridMonitorTaskListResponseBodyTaskListSLSProcessConfigStatistics) Validate() error {
	return dara.Validate(s)
}
