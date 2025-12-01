// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeRestoreTaskListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeRestoreTaskListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeRestoreTaskListResponseBody
	GetHttpStatusCode() *int32
	SetItems(v *DescribeRestoreTaskListResponseBodyItems) *DescribeRestoreTaskListResponseBody
	GetItems() *DescribeRestoreTaskListResponseBodyItems
	SetPageNum(v int32) *DescribeRestoreTaskListResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeRestoreTaskListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRestoreTaskListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeRestoreTaskListResponseBody
	GetSuccess() *bool
	SetTotalElements(v int32) *DescribeRestoreTaskListResponseBody
	GetTotalElements() *int32
	SetTotalPages(v int32) *DescribeRestoreTaskListResponseBody
	GetTotalPages() *int32
}

type DescribeRestoreTaskListResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backup schedule.
	Items *DescribeRestoreTaskListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9C397502-B4F2-4E22-AD97-C81F0049F3F3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of restore tasks.
	//
	// example:
	//
	// 1
	TotalElements *int32 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeRestoreTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTaskListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeRestoreTaskListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeRestoreTaskListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeRestoreTaskListResponseBody) GetItems() *DescribeRestoreTaskListResponseBodyItems {
	return s.Items
}

func (s *DescribeRestoreTaskListResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeRestoreTaskListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestoreTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRestoreTaskListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRestoreTaskListResponseBody) GetTotalElements() *int32 {
	return s.TotalElements
}

func (s *DescribeRestoreTaskListResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeRestoreTaskListResponseBody) SetErrCode(v string) *DescribeRestoreTaskListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetErrMessage(v string) *DescribeRestoreTaskListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetHttpStatusCode(v int32) *DescribeRestoreTaskListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetItems(v *DescribeRestoreTaskListResponseBodyItems) *DescribeRestoreTaskListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetPageNum(v int32) *DescribeRestoreTaskListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetPageSize(v int32) *DescribeRestoreTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetRequestId(v string) *DescribeRestoreTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetSuccess(v bool) *DescribeRestoreTaskListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetTotalElements(v int32) *DescribeRestoreTaskListResponseBody {
	s.TotalElements = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) SetTotalPages(v int32) *DescribeRestoreTaskListResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRestoreTaskListResponseBodyItems struct {
	RestoreTaskDetail []*DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail `json:"RestoreTaskDetail,omitempty" xml:"RestoreTaskDetail,omitempty" type:"Repeated"`
}

func (s DescribeRestoreTaskListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTaskListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTaskListResponseBodyItems) GetRestoreTaskDetail() []*DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	return s.RestoreTaskDetail
}

func (s *DescribeRestoreTaskListResponseBodyItems) SetRestoreTaskDetail(v []*DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) *DescribeRestoreTaskListResponseBodyItems {
	s.RestoreTaskDetail = v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItems) Validate() error {
	if s.RestoreTaskDetail != nil {
		for _, item := range s.RestoreTaskDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail struct {
	// The ID of the backup gateway.
	//
	// example:
	//
	// 324234332
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// The ID of the backup schedule.
	//
	// example:
	//
	// dbs1hvb0wXXXX
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the full backup set that is used in the restore task.
	//
	// example:
	//
	// 1ibh2f5uXXX
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The restore progress of the incremental log files.
	//
	// example:
	//
	// 100
	ContinuousRestoreProgress *int32 `json:"ContinuousRestoreProgress,omitempty" xml:"ContinuousRestoreProgress,omitempty"`
	// The unique ID (UID) of the Alibaba Cloud account to which the backup schedule belongs.
	//
	// example:
	//
	// 2xxx7778xxxxxxxxxx
	CrossAliyunId *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	// The name of the RAM role that can be used to perform backups across Alibaba Cloud accounts.
	//
	// example:
	//
	// test123
	CrossRoleName *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test
	DestinationEndpointDatabaseName *string `json:"DestinationEndpointDatabaseName,omitempty" xml:"DestinationEndpointDatabaseName,omitempty"`
	// The ID of the database instance.
	//
	// example:
	//
	// rm-bp1p8c2947XXX
	DestinationEndpointInstanceID *string `json:"DestinationEndpointInstanceID,omitempty" xml:"DestinationEndpointInstanceID,omitempty"`
	// The location of the database. Valid values:
	//
	// 	- RDS
	//
	// 	- ECS
	//
	// 	- Express: The database is connected to DBS by using Express Connect, VPN Gateway, or Smart Access Gateway.
	//
	// 	- Agent: The database is connected to DBS over a DBS backup gateway.
	//
	// 	- DDS: The database is an ApsaraDB for MongoDB database.
	//
	// 	- Other: The database is connected to DBS by using the IP address and port of the database.
	//
	// example:
	//
	// RDS
	DestinationEndpointInstanceType *string `json:"DestinationEndpointInstanceType,omitempty" xml:"DestinationEndpointInstanceType,omitempty"`
	// The endpoint that is used to connect to the database.
	//
	// example:
	//
	// 100.X.X.10:33204
	DestinationEndpointIpPort *string `json:"DestinationEndpointIpPort,omitempty" xml:"DestinationEndpointIpPort,omitempty"`
	// The SID of the Oracle database.
	//
	// example:
	//
	// test
	DestinationEndpointOracleSID *string `json:"DestinationEndpointOracleSID,omitempty" xml:"DestinationEndpointOracleSID,omitempty"`
	// The region in which the database is deployed.
	//
	// example:
	//
	// cn-hangzhou
	DestinationEndpointRegion *string `json:"DestinationEndpointRegion,omitempty" xml:"DestinationEndpointRegion,omitempty"`
	// The username of the account that is used to connect to the database.
	//
	// example:
	//
	// test
	DestinationEndpointUserName *string `json:"DestinationEndpointUserName,omitempty" xml:"DestinationEndpointUserName,omitempty"`
	// The error message.
	//
	// example:
	//
	// NULL
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The restore progress of the full backup data.
	//
	// example:
	//
	// 100
	FullDataRestoreProgress *int32 `json:"FullDataRestoreProgress,omitempty" xml:"FullDataRestoreProgress,omitempty"`
	// The progress of schema restore after full data restore.
	//
	// example:
	//
	// 100
	FullStruAfterRestoreProgress *int32 `json:"FullStruAfterRestoreProgress,omitempty" xml:"FullStruAfterRestoreProgress,omitempty"`
	// The progress of schema restore before full data restore.
	//
	// example:
	//
	// 0
	FullStruforeRestoreProgress *int32 `json:"FullStruforeRestoreProgress,omitempty" xml:"FullStruforeRestoreProgress,omitempty"`
	// The directory of the destination database to which the objects were restored.
	//
	// example:
	//
	// 100.X.X.10:33204
	RestoreDir *string `json:"RestoreDir,omitempty" xml:"RestoreDir,omitempty"`
	// The objects to be restored.
	//
	// example:
	//
	// [     {         \\"DBName\\":\\"test\\"       } ]
	RestoreObjects *string `json:"RestoreObjects,omitempty" xml:"RestoreObjects,omitempty"`
	// The status of the restore task. Valid values:
	//
	// 	- init: The restore task is not started or does not pass the precheck.
	//
	// 	- running: The restore task is running.
	//
	// 	- stop: The restore task failed.
	//
	// 	- pause: The restore task is stopped.
	//
	// 	- check_pass: The restore task passed the precheck.
	//
	// example:
	//
	// running
	RestoreStatus *string `json:"RestoreStatus,omitempty" xml:"RestoreStatus,omitempty"`
	// The time when the restore task was created, such as 1554560477000.
	//
	// example:
	//
	// 1554560477000
	RestoreTaskCreateTime *int64 `json:"RestoreTaskCreateTime,omitempty" xml:"RestoreTaskCreateTime,omitempty"`
	// The time when the restore task was complete, such as 1554560477000.
	//
	// example:
	//
	// 1554560477000
	RestoreTaskFinishTime *int64 `json:"RestoreTaskFinishTime,omitempty" xml:"RestoreTaskFinishTime,omitempty"`
	// The ID of the restore task.
	//
	// example:
	//
	// s1XXXXXX
	RestoreTaskId *string `json:"RestoreTaskId,omitempty" xml:"RestoreTaskId,omitempty"`
	// The name of the restore task.
	//
	// example:
	//
	// test
	RestoreTaskName *string `json:"RestoreTaskName,omitempty" xml:"RestoreTaskName,omitempty"`
	// The time to run the restore task, such as 1554560477000.
	//
	// example:
	//
	// 1554560477000
	RestoreTime *int64 `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
}

func (s DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetBackupGatewayId() *int64 {
	return s.BackupGatewayId
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetContinuousRestoreProgress() *int32 {
	return s.ContinuousRestoreProgress
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetCrossAliyunId() *string {
	return s.CrossAliyunId
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetCrossRoleName() *string {
	return s.CrossRoleName
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetDestinationEndpointDatabaseName() *string {
	return s.DestinationEndpointDatabaseName
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetDestinationEndpointInstanceID() *string {
	return s.DestinationEndpointInstanceID
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetDestinationEndpointInstanceType() *string {
	return s.DestinationEndpointInstanceType
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetDestinationEndpointIpPort() *string {
	return s.DestinationEndpointIpPort
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetDestinationEndpointOracleSID() *string {
	return s.DestinationEndpointOracleSID
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetDestinationEndpointRegion() *string {
	return s.DestinationEndpointRegion
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetDestinationEndpointUserName() *string {
	return s.DestinationEndpointUserName
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetFullDataRestoreProgress() *int32 {
	return s.FullDataRestoreProgress
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetFullStruAfterRestoreProgress() *int32 {
	return s.FullStruAfterRestoreProgress
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetFullStruforeRestoreProgress() *int32 {
	return s.FullStruforeRestoreProgress
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetRestoreDir() *string {
	return s.RestoreDir
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetRestoreObjects() *string {
	return s.RestoreObjects
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetRestoreStatus() *string {
	return s.RestoreStatus
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetRestoreTaskCreateTime() *int64 {
	return s.RestoreTaskCreateTime
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetRestoreTaskFinishTime() *int64 {
	return s.RestoreTaskFinishTime
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetRestoreTaskId() *string {
	return s.RestoreTaskId
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetRestoreTaskName() *string {
	return s.RestoreTaskName
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) GetRestoreTime() *int64 {
	return s.RestoreTime
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetBackupGatewayId(v int64) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.BackupGatewayId = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetBackupPlanId(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetBackupSetId(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.BackupSetId = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetContinuousRestoreProgress(v int32) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.ContinuousRestoreProgress = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetCrossAliyunId(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.CrossAliyunId = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetCrossRoleName(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.CrossRoleName = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetDestinationEndpointDatabaseName(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.DestinationEndpointDatabaseName = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetDestinationEndpointInstanceID(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.DestinationEndpointInstanceID = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetDestinationEndpointInstanceType(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.DestinationEndpointInstanceType = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetDestinationEndpointIpPort(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.DestinationEndpointIpPort = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetDestinationEndpointOracleSID(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.DestinationEndpointOracleSID = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetDestinationEndpointRegion(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.DestinationEndpointRegion = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetDestinationEndpointUserName(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.DestinationEndpointUserName = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetErrMessage(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.ErrMessage = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetFullDataRestoreProgress(v int32) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.FullDataRestoreProgress = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetFullStruAfterRestoreProgress(v int32) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.FullStruAfterRestoreProgress = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetFullStruforeRestoreProgress(v int32) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.FullStruforeRestoreProgress = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetRestoreDir(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.RestoreDir = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetRestoreObjects(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.RestoreObjects = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetRestoreStatus(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.RestoreStatus = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetRestoreTaskCreateTime(v int64) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.RestoreTaskCreateTime = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetRestoreTaskFinishTime(v int64) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.RestoreTaskFinishTime = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetRestoreTaskId(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.RestoreTaskId = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetRestoreTaskName(v string) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.RestoreTaskName = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) SetRestoreTime(v int64) *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail {
	s.RestoreTime = &v
	return s
}

func (s *DescribeRestoreTaskListResponseBodyItemsRestoreTaskDetail) Validate() error {
	return dara.Validate(s)
}
