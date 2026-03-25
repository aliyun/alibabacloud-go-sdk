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
	if s.FlowLogs != nil {
		if err := s.FlowLogs.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type DescribeFlowlogsResponseBodyFlowLogsFlowLog struct {
	CenId                     *string                                          `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CreationTime              *string                                          `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description               *string                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	FlowLogId                 *string                                          `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	FlowLogName               *string                                          `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	FlowLogVersion            *string                                          `json:"FlowLogVersion,omitempty" xml:"FlowLogVersion,omitempty"`
	Interval                  *int64                                           `json:"Interval,omitempty" xml:"Interval,omitempty"`
	LogFormatString           *string                                          `json:"LogFormatString,omitempty" xml:"LogFormatString,omitempty"`
	LogStoreName              *string                                          `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	ProjectName               *string                                          `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RegionId                  *string                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status                    *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                      *DescribeFlowlogsResponseBodyFlowLogsFlowLogTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TransitRouterAttachmentId *string                                          `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterId           *string                                          `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
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
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type DescribeFlowlogsResponseBodyFlowLogsFlowLogTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
