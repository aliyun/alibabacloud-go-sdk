// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowlogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlowLogs(v *DescribeFlowlogsResponseBodyFlowLogs) *DescribeFlowlogsResponseBody
	GetFlowLogs() *DescribeFlowlogsResponseBodyFlowLogs
	SetPageNumber(v string) *DescribeFlowlogsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeFlowlogsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeFlowlogsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeFlowlogsResponseBody
	GetSuccess() *string
	SetTotalCount(v string) *DescribeFlowlogsResponseBody
	GetTotalCount() *string
}

type DescribeFlowlogsResponseBody struct {
	// The information about the flow log.
	FlowLogs *DescribeFlowlogsResponseBodyFlowLogs `json:"FlowLogs,omitempty" xml:"FlowLogs,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F7DDDC17-FA06-4AC2-8F35-59D2470FCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFlowlogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowlogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsResponseBody) GetFlowLogs() *DescribeFlowlogsResponseBodyFlowLogs {
	return s.FlowLogs
}

func (s *DescribeFlowlogsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeFlowlogsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeFlowlogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFlowlogsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeFlowlogsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeFlowlogsResponseBody) SetFlowLogs(v *DescribeFlowlogsResponseBodyFlowLogs) *DescribeFlowlogsResponseBody {
	s.FlowLogs = v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetPageNumber(v string) *DescribeFlowlogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetPageSize(v string) *DescribeFlowlogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetRequestId(v string) *DescribeFlowlogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetSuccess(v string) *DescribeFlowlogsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFlowlogsResponseBody) SetTotalCount(v string) *DescribeFlowlogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFlowlogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFlowlogsResponseBodyFlowLogs struct {
	FlowLog []*DescribeFlowlogsResponseBodyFlowLogsFlowLog `json:"FlowLog,omitempty" xml:"FlowLog,omitempty" type:"Repeated"`
}

func (s DescribeFlowlogsResponseBodyFlowLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowlogsResponseBodyFlowLogs) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsResponseBodyFlowLogs) GetFlowLog() []*DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	return s.FlowLog
}

func (s *DescribeFlowlogsResponseBodyFlowLogs) SetFlowLog(v []*DescribeFlowlogsResponseBodyFlowLogsFlowLog) *DescribeFlowlogsResponseBodyFlowLogs {
	s.FlowLog = v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogs) Validate() error {
	return dara.Validate(s)
}

type DescribeFlowlogsResponseBodyFlowLogsFlowLog struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The time when the flow log was created.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-07-24T13:00:52Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the flow log.
	//
	// example:
	//
	// myFlowlog
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the flow log.
	//
	// example:
	//
	// flowlog-m5evbtbpt****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// The name of the flow log.
	//
	// example:
	//
	// myFlowlog
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// The flow log version.
	//
	// Flow logs are automatically created in the latest version, which is **3**.
	//
	// example:
	//
	// 3
	FlowLogVersion *string `json:"FlowLogVersion,omitempty" xml:"FlowLogVersion,omitempty"`
	// The time window for collecting log data. Unit: seconds. Valid values: **60*	- or **600*	- Default value: **600**.
	//
	// example:
	//
	// 60
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The string that defines the format of the flow log. Format:
	//
	// `${Field 1}${Field 2}${Field 3}`
	//
	// example:
	//
	// ${srcaddr}${dstaddr}${bytes}
	LogFormatString *string `json:"LogFormatString,omitempty" xml:"LogFormatString,omitempty"`
	// The Logstore that stores the captured traffic data.
	//
	// example:
	//
	// FlowLogStore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The name of the project that stores the captured traffic data.
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
	// The status of the flow log. Valid values:
	//
	// 	- **Active**: The flow log is enabled.
	//
	// 	- **Inactive**: The flow log is disabled.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags *DescribeFlowlogsResponseBodyFlowLogsFlowLogTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the network instance connection
	//
	// example:
	//
	// tr-attach-5x4o4ynzuqbv6g****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-bp1g9313sx675zr1lajmj
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s DescribeFlowlogsResponseBodyFlowLogsFlowLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowlogsResponseBodyFlowLogsFlowLog) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) GetCenId() *string {
	return s.CenId
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) GetDescription() *string {
	return s.Description
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) GetFlowLogName() *string {
	return s.FlowLogName
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) GetFlowLogVersion() *string {
	return s.FlowLogVersion
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) GetLogFormatString() *string {
	return s.LogFormatString
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) GetStatus() *string {
	return s.Status
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) GetTags() *DescribeFlowlogsResponseBodyFlowLogsFlowLogTags {
	return s.Tags
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetCenId(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.CenId = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetCreationTime(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.CreationTime = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetDescription(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.Description = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetFlowLogId(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.FlowLogId = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetFlowLogName(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.FlowLogName = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetFlowLogVersion(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.FlowLogVersion = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetInterval(v int64) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.Interval = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetLogFormatString(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.LogFormatString = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetLogStoreName(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.LogStoreName = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetProjectName(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.ProjectName = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetRegionId(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetStatus(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.Status = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetTags(v *DescribeFlowlogsResponseBodyFlowLogsFlowLogTags) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.Tags = v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetTransitRouterAttachmentId(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) SetTransitRouterId(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLog {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLog) Validate() error {
	return dara.Validate(s)
}

type DescribeFlowlogsResponseBodyFlowLogsFlowLogTags struct {
	Tag []*DescribeFlowlogsResponseBodyFlowLogsFlowLogTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeFlowlogsResponseBodyFlowLogsFlowLogTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowlogsResponseBodyFlowLogsFlowLogTags) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLogTags) GetTag() []*DescribeFlowlogsResponseBodyFlowLogsFlowLogTagsTag {
	return s.Tag
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLogTags) SetTag(v []*DescribeFlowlogsResponseBodyFlowLogsFlowLogTagsTag) *DescribeFlowlogsResponseBodyFlowLogsFlowLogTags {
	s.Tag = v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLogTags) Validate() error {
	return dara.Validate(s)
}

type DescribeFlowlogsResponseBodyFlowLogsFlowLogTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFlowlogsResponseBodyFlowLogsFlowLogTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowlogsResponseBodyFlowLogsFlowLogTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLogTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLogTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLogTagsTag) SetKey(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLogTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLogTagsTag) SetValue(v string) *DescribeFlowlogsResponseBodyFlowLogsFlowLogTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeFlowlogsResponseBodyFlowLogsFlowLogTagsTag) Validate() error {
	return dara.Validate(s)
}
