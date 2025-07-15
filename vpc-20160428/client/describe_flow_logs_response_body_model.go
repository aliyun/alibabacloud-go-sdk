// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlowLogs(v *DescribeFlowLogsResponseBodyFlowLogs) *DescribeFlowLogsResponseBody
	GetFlowLogs() *DescribeFlowLogsResponseBodyFlowLogs
	SetPageNumber(v string) *DescribeFlowLogsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeFlowLogsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeFlowLogsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeFlowLogsResponseBody
	GetSuccess() *string
	SetTotalCount(v string) *DescribeFlowLogsResponseBody
	GetTotalCount() *string
}

type DescribeFlowLogsResponseBody struct {
	// The information about the flow logs.
	FlowLogs *DescribeFlowLogsResponseBodyFlowLogs `json:"FlowLogs,omitempty" xml:"FlowLogs,omitempty" type:"Struct"`
	// The number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F7DDDC17-FA06-4AC2-8F35-59D2470FCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The number of flow logs that are queried.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFlowLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponseBody) GetFlowLogs() *DescribeFlowLogsResponseBodyFlowLogs {
	return s.FlowLogs
}

func (s *DescribeFlowLogsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeFlowLogsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeFlowLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFlowLogsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeFlowLogsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeFlowLogsResponseBody) SetFlowLogs(v *DescribeFlowLogsResponseBodyFlowLogs) *DescribeFlowLogsResponseBody {
	s.FlowLogs = v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetPageNumber(v string) *DescribeFlowLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetPageSize(v string) *DescribeFlowLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetRequestId(v string) *DescribeFlowLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetSuccess(v string) *DescribeFlowLogsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) SetTotalCount(v string) *DescribeFlowLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFlowLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFlowLogsResponseBodyFlowLogs struct {
	FlowLog []*DescribeFlowLogsResponseBodyFlowLogsFlowLog `json:"FlowLog,omitempty" xml:"FlowLog,omitempty" type:"Repeated"`
}

func (s DescribeFlowLogsResponseBodyFlowLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsResponseBodyFlowLogs) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) GetFlowLog() []*DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	return s.FlowLog
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) SetFlowLog(v []*DescribeFlowLogsResponseBodyFlowLogsFlowLog) *DescribeFlowLogsResponseBodyFlowLogs {
	s.FlowLog = v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogs) Validate() error {
	return dara.Validate(s)
}

type DescribeFlowLogsResponseBodyFlowLogsFlowLog struct {
	// The sampling interval of the flow log. Unit: minutes.
	//
	// example:
	//
	// 10
	AggregationInterval *int32 `json:"AggregationInterval,omitempty" xml:"AggregationInterval,omitempty"`
	// The business status of the flow log. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **FinancialLocked**
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The time when the flow log was created.
	//
	// example:
	//
	// 2022-01-21T03:08:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the flow log.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// If the flow log failed to be delivered, you can troubleshoot based on the following error messages that may be returned:
	//
	// 	- **UnavaliableTarget**: The Logstore of SLS is unavailable and cannot receive logs. Check whether the Logstore is available.
	//
	// 	- **ProjectNotExist**: The project of SLS does not exist. We recommend that you delete the project and create a new one.
	//
	// 	- **UnknownError**: An internal error occurred. Try again later.
	//
	// example:
	//
	// UnavaliableTarget
	FlowLogDeliverErrorMessage *string `json:"FlowLogDeliverErrorMessage,omitempty" xml:"FlowLogDeliverErrorMessage,omitempty"`
	// Indicates whether the flow log is delivered. Valid values:
	//
	// - **SUCCESS*	-
	//
	// - **FAILED**
	//
	// example:
	//
	// FAILED
	FlowLogDeliverStatus *string `json:"FlowLogDeliverStatus,omitempty" xml:"FlowLogDeliverStatus,omitempty"`
	// The ID of the flow log.
	//
	// example:
	//
	// fl-bp1f6qqhsrc2c12ta****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// The name of the flow log.
	//
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	IpVersion   *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The Logstore that stores the captured traffic data.
	//
	// example:
	//
	// FlowLogStore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The project that manages the captured traffic data.
	//
	// example:
	//
	// FlowLogProject
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The region ID of the flow log.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the flow log belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource from which traffic is captured.
	//
	// example:
	//
	// eni-askldfas****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource from which traffic is captured. Valid values:
	//
	// 	- **NetworkInterface**: ENI
	//
	// 	- **VSwitch**: all ENIs in a vSwitch
	//
	// 	- **VPC**: all ENIs in a VPC
	//
	// example:
	//
	// NetworkInterface
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The hosting type of the cloud service.
	//
	// 	- This parameter can be empty, which indicates that the flow log is created by the user.
	//
	// 	- If this parameter is not empty, the value is set to **sls**. The value sls indicates that the flow log is created in the Simple Log Service (SLS) console.
	//
	// > A flow log that is created in the SLS console can be displayed in the VPC list. However, you cannot modify, start, stop, or delete the flow log in the VPC console. If you want to manage the flow log, you can log on to the [SLS console](https://sls.console.aliyun.com) and perform required operations.
	//
	// example:
	//
	// sls
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The status of the flow log. Valid values:
	//
	// 	- **Active**
	//
	// 	- **Activating**
	//
	// 	- **Inactive**
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags *DescribeFlowLogsResponseBodyFlowLogsFlowLogTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The sampling scope of the traffic that is collected. Valid values:
	//
	// 	- **all*	- (default value): all traffic
	//
	// 	- **internetGateway**: Internet traffic
	//
	// > By default, the traffic path feature is unavailable. To use this feature, [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.11182188.console-base-top.dworkorder.18ae4882n3v6ZW#/ticket/createIndex).
	TrafficPath *DescribeFlowLogsResponseBodyFlowLogsFlowLogTrafficPath `json:"TrafficPath,omitempty" xml:"TrafficPath,omitempty" type:"Struct"`
	// The type of traffic that is captured by the flow log. Valid values:
	//
	// 	- **All**: all traffic
	//
	// 	- **Allow**: traffic that is allowed by access control
	//
	// 	- **Drop**: traffic that is denied by access control
	//
	// example:
	//
	// All
	TrafficType *string `json:"TrafficType,omitempty" xml:"TrafficType,omitempty"`
}

func (s DescribeFlowLogsResponseBodyFlowLogsFlowLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsResponseBodyFlowLogsFlowLog) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetAggregationInterval() *int32 {
	return s.AggregationInterval
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetDescription() *string {
	return s.Description
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetFlowLogDeliverErrorMessage() *string {
	return s.FlowLogDeliverErrorMessage
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetFlowLogDeliverStatus() *string {
	return s.FlowLogDeliverStatus
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetFlowLogName() *string {
	return s.FlowLogName
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetServiceType() *string {
	return s.ServiceType
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetStatus() *string {
	return s.Status
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetTags() *DescribeFlowLogsResponseBodyFlowLogsFlowLogTags {
	return s.Tags
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetTrafficPath() *DescribeFlowLogsResponseBodyFlowLogsFlowLogTrafficPath {
	return s.TrafficPath
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) GetTrafficType() *string {
	return s.TrafficType
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetAggregationInterval(v int32) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.AggregationInterval = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetBusinessStatus(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetCreationTime(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.CreationTime = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetDescription(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.Description = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetFlowLogDeliverErrorMessage(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.FlowLogDeliverErrorMessage = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetFlowLogDeliverStatus(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.FlowLogDeliverStatus = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetFlowLogId(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.FlowLogId = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetFlowLogName(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.FlowLogName = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetIpVersion(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.IpVersion = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetLogStoreName(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.LogStoreName = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetProjectName(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.ProjectName = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetRegionId(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetResourceGroupId(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetResourceId(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.ResourceId = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetResourceType(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.ResourceType = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetServiceType(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.ServiceType = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetStatus(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.Status = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetTags(v *DescribeFlowLogsResponseBodyFlowLogsFlowLogTags) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.Tags = v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetTrafficPath(v *DescribeFlowLogsResponseBodyFlowLogsFlowLogTrafficPath) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.TrafficPath = v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) SetTrafficType(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLog {
	s.TrafficType = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLog) Validate() error {
	return dara.Validate(s)
}

type DescribeFlowLogsResponseBodyFlowLogsFlowLogTags struct {
	Tag []*DescribeFlowLogsResponseBodyFlowLogsFlowLogTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeFlowLogsResponseBodyFlowLogsFlowLogTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsResponseBodyFlowLogsFlowLogTags) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogTags) GetTag() []*DescribeFlowLogsResponseBodyFlowLogsFlowLogTagsTag {
	return s.Tag
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogTags) SetTag(v []*DescribeFlowLogsResponseBodyFlowLogsFlowLogTagsTag) *DescribeFlowLogsResponseBodyFlowLogsFlowLogTags {
	s.Tag = v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogTags) Validate() error {
	return dara.Validate(s)
}

type DescribeFlowLogsResponseBodyFlowLogsFlowLogTagsTag struct {
	// The key of tag N.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFlowLogsResponseBodyFlowLogsFlowLogTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsResponseBodyFlowLogsFlowLogTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogTagsTag) SetKey(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLogTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogTagsTag) SetValue(v string) *DescribeFlowLogsResponseBodyFlowLogsFlowLogTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeFlowLogsResponseBodyFlowLogsFlowLogTrafficPath struct {
	TrafficPathList []*string `json:"TrafficPathList,omitempty" xml:"TrafficPathList,omitempty" type:"Repeated"`
}

func (s DescribeFlowLogsResponseBodyFlowLogsFlowLogTrafficPath) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsResponseBodyFlowLogsFlowLogTrafficPath) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogTrafficPath) GetTrafficPathList() []*string {
	return s.TrafficPathList
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogTrafficPath) SetTrafficPathList(v []*string) *DescribeFlowLogsResponseBodyFlowLogsFlowLogTrafficPath {
	s.TrafficPathList = v
	return s
}

func (s *DescribeFlowLogsResponseBodyFlowLogsFlowLogTrafficPath) Validate() error {
	return dara.Validate(s)
}
