// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ListAgentlessTaskResponseBodyList) *ListAgentlessTaskResponseBody
	GetList() []*ListAgentlessTaskResponseBodyList
	SetPageInfo(v *ListAgentlessTaskResponseBodyPageInfo) *ListAgentlessTaskResponseBody
	GetPageInfo() *ListAgentlessTaskResponseBodyPageInfo
	SetRequestId(v string) *ListAgentlessTaskResponseBody
	GetRequestId() *string
}

type ListAgentlessTaskResponseBody struct {
	// The tasks.
	List []*ListAgentlessTaskResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListAgentlessTaskResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1A975D03-5F49-5354-B2CB-3918D5DA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAgentlessTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentlessTaskResponseBody) GetList() []*ListAgentlessTaskResponseBodyList {
	return s.List
}

func (s *ListAgentlessTaskResponseBody) GetPageInfo() *ListAgentlessTaskResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListAgentlessTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentlessTaskResponseBody) SetList(v []*ListAgentlessTaskResponseBodyList) *ListAgentlessTaskResponseBody {
	s.List = v
	return s
}

func (s *ListAgentlessTaskResponseBody) SetPageInfo(v *ListAgentlessTaskResponseBodyPageInfo) *ListAgentlessTaskResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListAgentlessTaskResponseBody) SetRequestId(v string) *ListAgentlessTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentlessTaskResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentlessTaskResponseBodyList struct {
	// The end timestamp of the task. Unit: milliseconds.
	//
	// example:
	//
	// 1678895999999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID of the asset.
	//
	// example:
	//
	// s-wz95vuqky0ada4******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// sql-test-0****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 8.210.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 172.25.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The amount of data detected. Unit: MB.
	//
	// example:
	//
	// 154.11
	MeasureSpace *int64 `json:"MeasureSpace,omitempty" xml:"MeasureSpace,omitempty"`
	// The progress of the task.
	//
	// example:
	//
	// 60
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The execution progress of the check items.
	//
	// example:
	//
	// "{\\"scaVul\\":100,\\"binary\\":100,\\"baseline\\":100,\\"vul\\":100,\\"webshell\\":100,\\"script\\":100,\\"sensitiveInfo\\":100}"
	ProgressByProject *string `json:"ProgressByProject,omitempty" xml:"ProgressByProject,omitempty"`
	// The download URL of the report.
	//
	// example:
	//
	// http://newsas-oss-bucket.oss-cn-hangzhou.aliyuncs.com/agent_less_single_report****
	ReportDownloadUrl *string `json:"ReportDownloadUrl,omitempty" xml:"ReportDownloadUrl,omitempty"`
	// The status of the report. Valid values:
	//
	// 	- **PREPARED**: preparing
	//
	// 	- **RUNNING**: running
	//
	// 	- **SUCCESS**: succeeded
	//
	// 	- **TIMEOUT**: timed out
	//
	// 	- **FAILED**: failed
	//
	// example:
	//
	// SUCCESS
	ReportStatus *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	// The result of the detection.
	//
	// example:
	//
	// True
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The start timestamp of the task. Unit: milliseconds.
	//
	// example:
	//
	// 1672741657897
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the detection task.
	//
	// 	- **1**: The detection task is in progress.
	//
	// 	- **2**: The detection task is complete.
	//
	// 	- **3**: The detection task fails.
	//
	// 	- **4**: The detection task times out.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the asset that is detected.
	//
	// example:
	//
	// hkdevt****
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the asset that is detected. Valid values:
	//
	// 	- **1**: snapshot
	//
	// 	- **2**: image
	//
	// example:
	//
	// 2
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 1538****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the detection task.
	//
	// example:
	//
	// AGENTLESS_SCAN
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The UUID of the asset.
	//
	// example:
	//
	// 49e25e0f-bb51-4a5a-a1b3-13a4ddaa****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListAgentlessTaskResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessTaskResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListAgentlessTaskResponseBodyList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAgentlessTaskResponseBodyList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAgentlessTaskResponseBodyList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListAgentlessTaskResponseBodyList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListAgentlessTaskResponseBodyList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListAgentlessTaskResponseBodyList) GetMeasureSpace() *int64 {
	return s.MeasureSpace
}

func (s *ListAgentlessTaskResponseBodyList) GetProgress() *int32 {
	return s.Progress
}

func (s *ListAgentlessTaskResponseBodyList) GetProgressByProject() *string {
	return s.ProgressByProject
}

func (s *ListAgentlessTaskResponseBodyList) GetReportDownloadUrl() *string {
	return s.ReportDownloadUrl
}

func (s *ListAgentlessTaskResponseBodyList) GetReportStatus() *string {
	return s.ReportStatus
}

func (s *ListAgentlessTaskResponseBodyList) GetResult() *string {
	return s.Result
}

func (s *ListAgentlessTaskResponseBodyList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAgentlessTaskResponseBodyList) GetStatus() *int32 {
	return s.Status
}

func (s *ListAgentlessTaskResponseBodyList) GetTargetName() *string {
	return s.TargetName
}

func (s *ListAgentlessTaskResponseBodyList) GetTargetType() *int32 {
	return s.TargetType
}

func (s *ListAgentlessTaskResponseBodyList) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAgentlessTaskResponseBodyList) GetTaskName() *string {
	return s.TaskName
}

func (s *ListAgentlessTaskResponseBodyList) GetUuid() *string {
	return s.Uuid
}

func (s *ListAgentlessTaskResponseBodyList) SetEndTime(v int64) *ListAgentlessTaskResponseBodyList {
	s.EndTime = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetInstanceId(v string) *ListAgentlessTaskResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetInstanceName(v string) *ListAgentlessTaskResponseBodyList {
	s.InstanceName = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetInternetIp(v string) *ListAgentlessTaskResponseBodyList {
	s.InternetIp = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetIntranetIp(v string) *ListAgentlessTaskResponseBodyList {
	s.IntranetIp = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetMeasureSpace(v int64) *ListAgentlessTaskResponseBodyList {
	s.MeasureSpace = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetProgress(v int32) *ListAgentlessTaskResponseBodyList {
	s.Progress = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetProgressByProject(v string) *ListAgentlessTaskResponseBodyList {
	s.ProgressByProject = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetReportDownloadUrl(v string) *ListAgentlessTaskResponseBodyList {
	s.ReportDownloadUrl = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetReportStatus(v string) *ListAgentlessTaskResponseBodyList {
	s.ReportStatus = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetResult(v string) *ListAgentlessTaskResponseBodyList {
	s.Result = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetStartTime(v int64) *ListAgentlessTaskResponseBodyList {
	s.StartTime = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetStatus(v int32) *ListAgentlessTaskResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetTargetName(v string) *ListAgentlessTaskResponseBodyList {
	s.TargetName = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetTargetType(v int32) *ListAgentlessTaskResponseBodyList {
	s.TargetType = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetTaskId(v string) *ListAgentlessTaskResponseBodyList {
	s.TaskId = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetTaskName(v string) *ListAgentlessTaskResponseBodyList {
	s.TaskName = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) SetUuid(v string) *ListAgentlessTaskResponseBodyList {
	s.Uuid = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListAgentlessTaskResponseBodyPageInfo struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 55
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgentlessTaskResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessTaskResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListAgentlessTaskResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAgentlessTaskResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentlessTaskResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAgentlessTaskResponseBodyPageInfo) SetCurrentPage(v int32) *ListAgentlessTaskResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyPageInfo) SetPageSize(v int32) *ListAgentlessTaskResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyPageInfo) SetTotalCount(v int32) *ListAgentlessTaskResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListAgentlessTaskResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
