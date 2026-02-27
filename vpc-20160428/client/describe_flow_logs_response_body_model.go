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
	FlowLogs *DescribeFlowLogsResponseBodyFlowLogs `json:"FlowLogs,omitempty" xml:"FlowLogs,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of items per page in a paginated query.
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
	// Indicates whether the call was successful. Values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The number of entries in the queried flow log list.
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
	if s.FlowLogs != nil {
		if err := s.FlowLogs.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.FlowLog != nil {
		for _, item := range s.FlowLog {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFlowLogsResponseBodyFlowLogsFlowLog struct {
	AggregationInterval        *int32                                                  `json:"AggregationInterval,omitempty" xml:"AggregationInterval,omitempty"`
	BusinessStatus             *string                                                 `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	CreationTime               *string                                                 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description                *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	FlowLogDeliverErrorMessage *string                                                 `json:"FlowLogDeliverErrorMessage,omitempty" xml:"FlowLogDeliverErrorMessage,omitempty"`
	FlowLogDeliverStatus       *string                                                 `json:"FlowLogDeliverStatus,omitempty" xml:"FlowLogDeliverStatus,omitempty"`
	FlowLogId                  *string                                                 `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	FlowLogName                *string                                                 `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	IpVersion                  *string                                                 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	LogStoreName               *string                                                 `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	ProjectName                *string                                                 `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RegionId                   *string                                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId            *string                                                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId                 *string                                                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType               *string                                                 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ServiceType                *string                                                 `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Status                     *string                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                       *DescribeFlowLogsResponseBodyFlowLogsFlowLogTags        `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TrafficPath                *DescribeFlowLogsResponseBodyFlowLogsFlowLogTrafficPath `json:"TrafficPath,omitempty" xml:"TrafficPath,omitempty" type:"Struct"`
	TrafficType                *string                                                 `json:"TrafficType,omitempty" xml:"TrafficType,omitempty"`
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
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	if s.TrafficPath != nil {
		if err := s.TrafficPath.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFlowLogsResponseBodyFlowLogsFlowLogTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
