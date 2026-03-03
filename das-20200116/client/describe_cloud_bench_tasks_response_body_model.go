// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudBenchTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCloudBenchTasksResponseBody
	GetCode() *string
	SetData(v *DescribeCloudBenchTasksResponseBodyData) *DescribeCloudBenchTasksResponseBody
	GetData() *DescribeCloudBenchTasksResponseBodyData
	SetMessage(v string) *DescribeCloudBenchTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCloudBenchTasksResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeCloudBenchTasksResponseBody
	GetSuccess() *string
}

type DescribeCloudBenchTasksResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information, including the error codes and the number of entries that are returned.
	Data *DescribeCloudBenchTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudBenchTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudBenchTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudBenchTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCloudBenchTasksResponseBody) GetData() *DescribeCloudBenchTasksResponseBodyData {
	return s.Data
}

func (s *DescribeCloudBenchTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCloudBenchTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudBenchTasksResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeCloudBenchTasksResponseBody) SetCode(v string) *DescribeCloudBenchTasksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBody) SetData(v *DescribeCloudBenchTasksResponseBodyData) *DescribeCloudBenchTasksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudBenchTasksResponseBody) SetMessage(v string) *DescribeCloudBenchTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBody) SetRequestId(v string) *DescribeCloudBenchTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBody) SetSuccess(v string) *DescribeCloudBenchTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudBenchTasksResponseBodyData struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	Extra *string                                      `json:"Extra,omitempty" xml:"Extra,omitempty"`
	List  *DescribeCloudBenchTasksResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCloudBenchTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudBenchTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudBenchTasksResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *DescribeCloudBenchTasksResponseBodyData) GetList() *DescribeCloudBenchTasksResponseBodyDataList {
	return s.List
}

func (s *DescribeCloudBenchTasksResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeCloudBenchTasksResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudBenchTasksResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeCloudBenchTasksResponseBodyData) SetExtra(v string) *DescribeCloudBenchTasksResponseBodyData {
	s.Extra = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyData) SetList(v *DescribeCloudBenchTasksResponseBodyDataList) *DescribeCloudBenchTasksResponseBodyData {
	s.List = v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyData) SetPageNo(v int32) *DescribeCloudBenchTasksResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyData) SetPageSize(v int32) *DescribeCloudBenchTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyData) SetTotal(v int64) *DescribeCloudBenchTasksResponseBodyData {
	s.Total = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyData) Validate() error {
	if s.List != nil {
		if err := s.List.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudBenchTasksResponseBodyDataList struct {
	CloudbenchTasks []*DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks `json:"cloudbenchTasks,omitempty" xml:"cloudbenchTasks,omitempty" type:"Repeated"`
}

func (s DescribeCloudBenchTasksResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudBenchTasksResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeCloudBenchTasksResponseBodyDataList) GetCloudbenchTasks() []*DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	return s.CloudbenchTasks
}

func (s *DescribeCloudBenchTasksResponseBodyDataList) SetCloudbenchTasks(v []*DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) *DescribeCloudBenchTasksResponseBodyDataList {
	s.CloudbenchTasks = v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataList) Validate() error {
	if s.CloudbenchTasks != nil {
		for _, item := range s.CloudbenchTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks struct {
	ArchiveJobId        *string `json:"ArchiveJobId,omitempty" xml:"ArchiveJobId,omitempty"`
	ArchiveOssTableName *string `json:"ArchiveOssTableName,omitempty" xml:"ArchiveOssTableName,omitempty"`
	ArchiveState        *int32  `json:"ArchiveState,omitempty" xml:"ArchiveState,omitempty"`
	BackupId            *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupType          *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BenchStep           *string `json:"BenchStep,omitempty" xml:"BenchStep,omitempty"`
	BenchStepStatus     *string `json:"BenchStepStatus,omitempty" xml:"BenchStepStatus,omitempty"`
	ClientGatewayId     *string `json:"ClientGatewayId,omitempty" xml:"ClientGatewayId,omitempty"`
	ClientType          *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DstInstanceUuid     *string `json:"DstInstanceUuid,omitempty" xml:"DstInstanceUuid,omitempty"`
	DstIp               *string `json:"DstIp,omitempty" xml:"DstIp,omitempty"`
	DstPort             *int32  `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	DstType             *string `json:"DstType,omitempty" xml:"DstType,omitempty"`
	DtsJobClass         *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	DtsJobId            *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	DtsJobName          *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	DtsJobState         *int32  `json:"DtsJobState,omitempty" xml:"DtsJobState,omitempty"`
	DtsJobStatus        *string `json:"DtsJobStatus,omitempty" xml:"DtsJobStatus,omitempty"`
	EcsInstanceId       *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	EndState            *string `json:"EndState,omitempty" xml:"EndState,omitempty"`
	ErrorCode           *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage        *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	External            *string `json:"External,omitempty" xml:"External,omitempty"`
	Rate                *int32  `json:"Rate,omitempty" xml:"Rate,omitempty"`
	RequestDuration     *int64  `json:"RequestDuration,omitempty" xml:"RequestDuration,omitempty"`
	SmartPressureTime   *int32  `json:"SmartPressureTime,omitempty" xml:"SmartPressureTime,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SqlCompleteReuse    *string `json:"SqlCompleteReuse,omitempty" xml:"SqlCompleteReuse,omitempty"`
	SrcInstanceArea     *string `json:"SrcInstanceArea,omitempty" xml:"SrcInstanceArea,omitempty"`
	SrcInstanceUuid     *string `json:"SrcInstanceUuid,omitempty" xml:"SrcInstanceUuid,omitempty"`
	SrcPublicIp         *string `json:"SrcPublicIp,omitempty" xml:"SrcPublicIp,omitempty"`
	State               *string `json:"State,omitempty" xml:"State,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TableSchema         *string `json:"TableSchema,omitempty" xml:"TableSchema,omitempty"`
	TaskId              *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	Topic               *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Version             *string `json:"Version,omitempty" xml:"Version,omitempty"`
	WorkDir             *string `json:"WorkDir,omitempty" xml:"WorkDir,omitempty"`
}

func (s DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GoString() string {
	return s.String()
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetArchiveJobId() *string {
	return s.ArchiveJobId
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetArchiveOssTableName() *string {
	return s.ArchiveOssTableName
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetArchiveState() *int32 {
	return s.ArchiveState
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetBenchStep() *string {
	return s.BenchStep
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetBenchStepStatus() *string {
	return s.BenchStepStatus
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetClientGatewayId() *string {
	return s.ClientGatewayId
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetClientType() *string {
	return s.ClientType
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDescription() *string {
	return s.Description
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDstInstanceUuid() *string {
	return s.DstInstanceUuid
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDstIp() *string {
	return s.DstIp
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDstPort() *int32 {
	return s.DstPort
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDstType() *string {
	return s.DstType
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDtsJobClass() *string {
	return s.DtsJobClass
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDtsJobState() *int32 {
	return s.DtsJobState
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDtsJobStatus() *string {
	return s.DtsJobStatus
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetEndState() *string {
	return s.EndState
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetExternal() *string {
	return s.External
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetRate() *int32 {
	return s.Rate
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetRequestDuration() *int64 {
	return s.RequestDuration
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetSmartPressureTime() *int32 {
	return s.SmartPressureTime
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetSource() *string {
	return s.Source
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetSqlCompleteReuse() *string {
	return s.SqlCompleteReuse
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetSrcInstanceArea() *string {
	return s.SrcInstanceArea
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetSrcInstanceUuid() *string {
	return s.SrcInstanceUuid
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetSrcPublicIp() *string {
	return s.SrcPublicIp
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetState() *string {
	return s.State
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetStatus() *string {
	return s.Status
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetTableSchema() *string {
	return s.TableSchema
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetTopic() *string {
	return s.Topic
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetUserId() *string {
	return s.UserId
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetVersion() *string {
	return s.Version
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetWorkDir() *string {
	return s.WorkDir
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetArchiveJobId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ArchiveJobId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetArchiveOssTableName(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ArchiveOssTableName = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetArchiveState(v int32) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ArchiveState = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetBackupId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.BackupId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetBackupType(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.BackupType = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetBenchStep(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.BenchStep = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetBenchStepStatus(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.BenchStepStatus = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetClientGatewayId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ClientGatewayId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetClientType(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ClientType = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDescription(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Description = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDstInstanceUuid(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DstInstanceUuid = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDstIp(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DstIp = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDstPort(v int32) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DstPort = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDstType(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DstType = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDtsJobClass(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDtsJobId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DtsJobId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDtsJobName(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DtsJobName = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDtsJobState(v int32) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DtsJobState = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDtsJobStatus(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DtsJobStatus = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetEcsInstanceId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetEndState(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.EndState = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetErrorCode(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ErrorCode = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetErrorMessage(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetExternal(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.External = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetRate(v int32) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Rate = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetRequestDuration(v int64) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.RequestDuration = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSmartPressureTime(v int32) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.SmartPressureTime = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSource(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Source = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSqlCompleteReuse(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.SqlCompleteReuse = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSrcInstanceArea(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.SrcInstanceArea = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSrcInstanceUuid(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.SrcInstanceUuid = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSrcPublicIp(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.SrcPublicIp = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetState(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.State = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetStatus(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Status = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetTableSchema(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.TableSchema = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetTaskId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetTaskType(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.TaskType = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetTopic(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Topic = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetUserId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.UserId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetVersion(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Version = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetWorkDir(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.WorkDir = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) Validate() error {
	return dara.Validate(s)
}
