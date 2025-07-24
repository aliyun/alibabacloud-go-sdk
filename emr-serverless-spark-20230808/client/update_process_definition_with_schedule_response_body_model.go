// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProcessDefinitionWithScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateProcessDefinitionWithScheduleResponseBody
	GetCode() *int32
	SetData(v *UpdateProcessDefinitionWithScheduleResponseBodyData) *UpdateProcessDefinitionWithScheduleResponseBody
	GetData() *UpdateProcessDefinitionWithScheduleResponseBodyData
	SetFailed(v string) *UpdateProcessDefinitionWithScheduleResponseBody
	GetFailed() *string
	SetHttpStatusCode(v int32) *UpdateProcessDefinitionWithScheduleResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *UpdateProcessDefinitionWithScheduleResponseBody
	GetMsg() *string
	SetRequestId(v string) *UpdateProcessDefinitionWithScheduleResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateProcessDefinitionWithScheduleResponseBody
	GetSuccess() *string
}

type UpdateProcessDefinitionWithScheduleResponseBody struct {
	// The code that is returned by the backend server.
	//
	// example:
	//
	// 1400009
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *UpdateProcessDefinitionWithScheduleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Indicates whether the request failed.
	//
	// example:
	//
	// false
	Failed *string `json:"failed,omitempty" xml:"failed,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The description of the returned code.
	//
	// example:
	//
	// No permission for resource action
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) GetData() *UpdateProcessDefinitionWithScheduleResponseBodyData {
	return s.Data
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) GetFailed() *string {
	return s.Failed
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) SetCode(v int32) *UpdateProcessDefinitionWithScheduleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) SetData(v *UpdateProcessDefinitionWithScheduleResponseBodyData) *UpdateProcessDefinitionWithScheduleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) SetFailed(v string) *UpdateProcessDefinitionWithScheduleResponseBody {
	s.Failed = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) SetHttpStatusCode(v int32) *UpdateProcessDefinitionWithScheduleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) SetMsg(v string) *UpdateProcessDefinitionWithScheduleResponseBody {
	s.Msg = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) SetRequestId(v string) *UpdateProcessDefinitionWithScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) SetSuccess(v string) *UpdateProcessDefinitionWithScheduleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateProcessDefinitionWithScheduleResponseBodyData struct {
	// The email address to receive alerts.
	//
	// example:
	//
	// foo_bar@spark.alert.invalid.com
	AlertEmailAddress *string `json:"alertEmailAddress,omitempty" xml:"alertEmailAddress,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// alicloud_ack_one_cluster
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// 12***********
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The time when the workflow was created.
	//
	// example:
	//
	// 2024-09-05T02:03:19Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The CRON expression that is used for scheduling.
	//
	// example:
	//
	// 0 0 0 	- 	- ?
	Crontab *string `json:"crontab,omitempty" xml:"crontab,omitempty"`
	// The node description.
	//
	// example:
	//
	// 1
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The end of the end time range.
	//
	// example:
	//
	// 1710432000000
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The execution policy.
	//
	// example:
	//
	// SERIAL
	ExecutionType *string `json:"executionType,omitempty" xml:"executionType,omitempty"`
	// The serial number of the workflow.
	//
	// example:
	//
	// 123223
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the workflow.
	//
	// example:
	//
	// ods_batch_workflow
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the project to which the workflow belongs.
	//
	// example:
	//
	// w-********
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// The status of the workflow.
	//
	// example:
	//
	// ONLINE
	ReleaseState *string `json:"releaseState,omitempty" xml:"releaseState,omitempty"`
	// The start time of the scheduling.
	//
	// example:
	//
	// 0
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The ID of the time zone.
	//
	// example:
	//
	// Asia/Shanghai
	TimezoneId *string `json:"timezoneId,omitempty" xml:"timezoneId,omitempty"`
	// The time when the workflow was updated.
	//
	// example:
	//
	// 2024-03-05T06:24:27Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The ID of the user that is used to initiate a scheduling.
	//
	// example:
	//
	// 113*********
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The name of the user that is used to initiate a scheduling.
	//
	// example:
	//
	// w-********
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
	// The hash code of the version.
	//
	// example:
	//
	// dwerf*********
	VersionHashCode *string `json:"versionHashCode,omitempty" xml:"versionHashCode,omitempty"`
}

func (s UpdateProcessDefinitionWithScheduleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionWithScheduleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetAlertEmailAddress() *string {
	return s.AlertEmailAddress
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetBizId() *string {
	return s.BizId
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetCrontab() *string {
	return s.Crontab
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetExecutionType() *string {
	return s.ExecutionType
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetId() *string {
	return s.Id
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetReleaseState() *string {
	return s.ReleaseState
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetTimezoneId() *string {
	return s.TimezoneId
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetVersion() *int32 {
	return s.Version
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) GetVersionHashCode() *string {
	return s.VersionHashCode
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetAlertEmailAddress(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.AlertEmailAddress = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetBizId(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.BizId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetCode(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.Code = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetCreateTime(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetCrontab(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.Crontab = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetDescription(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetEndTime(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetExecutionType(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.ExecutionType = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetId(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetName(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetProjectName(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetReleaseState(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.ReleaseState = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetStartTime(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetTimezoneId(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.TimezoneId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetUpdateTime(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetUserId(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.UserId = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetUserName(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.UserName = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetVersion(v int32) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.Version = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) SetVersionHashCode(v string) *UpdateProcessDefinitionWithScheduleResponseBodyData {
	s.VersionHashCode = &v
	return s
}

func (s *UpdateProcessDefinitionWithScheduleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
