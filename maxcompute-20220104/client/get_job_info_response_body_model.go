// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetJobInfoResponseBodyData) *GetJobInfoResponseBody
	GetData() *GetJobInfoResponseBodyData
	SetErrorCode(v string) *GetJobInfoResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetJobInfoResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *GetJobInfoResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetJobInfoResponseBody
	GetRequestId() *string
}

type GetJobInfoResponseBody struct {
	Data      *GetJobInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrorCode *string                     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                     `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HttpCode  *int32                      `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetJobInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBody) GetData() *GetJobInfoResponseBodyData {
	return s.Data
}

func (s *GetJobInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetJobInfoResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetJobInfoResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetJobInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobInfoResponseBody) SetData(v *GetJobInfoResponseBodyData) *GetJobInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetJobInfoResponseBody) SetErrorCode(v string) *GetJobInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetJobInfoResponseBody) SetErrorMsg(v string) *GetJobInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetJobInfoResponseBody) SetHttpCode(v int32) *GetJobInfoResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetJobInfoResponseBody) SetRequestId(v string) *GetJobInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobInfoResponseBodyData struct {
	CuUsage          *int64                                        `json:"cuUsage,omitempty" xml:"cuUsage,omitempty"`
	EndAtTime        *int64                                        `json:"endAtTime,omitempty" xml:"endAtTime,omitempty"`
	ExtNodeId        *string                                       `json:"extNodeId,omitempty" xml:"extNodeId,omitempty"`
	ExtNodeOnDuty    *string                                       `json:"extNodeOnDuty,omitempty" xml:"extNodeOnDuty,omitempty"`
	ExtPlantFrom     *string                                       `json:"extPlantFrom,omitempty" xml:"extPlantFrom,omitempty"`
	InputBytes       *float64                                      `json:"inputBytes,omitempty" xml:"inputBytes,omitempty"`
	InstanceId       *string                                       `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	JobOwner         *string                                       `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	JobSubStatusList []*GetJobInfoResponseBodyDataJobSubStatusList `json:"jobSubStatusList,omitempty" xml:"jobSubStatusList,omitempty" type:"Repeated"`
	JobType          *string                                       `json:"jobType,omitempty" xml:"jobType,omitempty"`
	MemoryUsage      *int64                                        `json:"memoryUsage,omitempty" xml:"memoryUsage,omitempty"`
	Priority         *int64                                        `json:"priority,omitempty" xml:"priority,omitempty"`
	Project          *string                                       `json:"project,omitempty" xml:"project,omitempty"`
	QuotaNickname    *string                                       `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
	QuotaType        *string                                       `json:"quotaType,omitempty" xml:"quotaType,omitempty"`
	Region           *string                                       `json:"region,omitempty" xml:"region,omitempty"`
	RunningAtTime    *int64                                        `json:"runningAtTime,omitempty" xml:"runningAtTime,omitempty"`
	RunningTime      *int64                                        `json:"runningTime,omitempty" xml:"runningTime,omitempty"`
	SceneResults     []*GetJobInfoResponseBodyDataSceneResults     `json:"sceneResults,omitempty" xml:"sceneResults,omitempty" type:"Repeated"`
	Signature        *string                                       `json:"signature,omitempty" xml:"signature,omitempty"`
	Status           *string                                       `json:"status,omitempty" xml:"status,omitempty"`
	SubmittedAtTime  *int64                                        `json:"submittedAtTime,omitempty" xml:"submittedAtTime,omitempty"`
	TenantId         *string                                       `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	TotalTime        *int64                                        `json:"totalTime,omitempty" xml:"totalTime,omitempty"`
	WaitingTime      *int64                                        `json:"waitingTime,omitempty" xml:"waitingTime,omitempty"`
}

func (s GetJobInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyData) GetCuUsage() *int64 {
	return s.CuUsage
}

func (s *GetJobInfoResponseBodyData) GetEndAtTime() *int64 {
	return s.EndAtTime
}

func (s *GetJobInfoResponseBodyData) GetExtNodeId() *string {
	return s.ExtNodeId
}

func (s *GetJobInfoResponseBodyData) GetExtNodeOnDuty() *string {
	return s.ExtNodeOnDuty
}

func (s *GetJobInfoResponseBodyData) GetExtPlantFrom() *string {
	return s.ExtPlantFrom
}

func (s *GetJobInfoResponseBodyData) GetInputBytes() *float64 {
	return s.InputBytes
}

func (s *GetJobInfoResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetJobInfoResponseBodyData) GetJobOwner() *string {
	return s.JobOwner
}

func (s *GetJobInfoResponseBodyData) GetJobSubStatusList() []*GetJobInfoResponseBodyDataJobSubStatusList {
	return s.JobSubStatusList
}

func (s *GetJobInfoResponseBodyData) GetJobType() *string {
	return s.JobType
}

func (s *GetJobInfoResponseBodyData) GetMemoryUsage() *int64 {
	return s.MemoryUsage
}

func (s *GetJobInfoResponseBodyData) GetPriority() *int64 {
	return s.Priority
}

func (s *GetJobInfoResponseBodyData) GetProject() *string {
	return s.Project
}

func (s *GetJobInfoResponseBodyData) GetQuotaNickname() *string {
	return s.QuotaNickname
}

func (s *GetJobInfoResponseBodyData) GetQuotaType() *string {
	return s.QuotaType
}

func (s *GetJobInfoResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *GetJobInfoResponseBodyData) GetRunningAtTime() *int64 {
	return s.RunningAtTime
}

func (s *GetJobInfoResponseBodyData) GetRunningTime() *int64 {
	return s.RunningTime
}

func (s *GetJobInfoResponseBodyData) GetSceneResults() []*GetJobInfoResponseBodyDataSceneResults {
	return s.SceneResults
}

func (s *GetJobInfoResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GetJobInfoResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetJobInfoResponseBodyData) GetSubmittedAtTime() *int64 {
	return s.SubmittedAtTime
}

func (s *GetJobInfoResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *GetJobInfoResponseBodyData) GetTotalTime() *int64 {
	return s.TotalTime
}

func (s *GetJobInfoResponseBodyData) GetWaitingTime() *int64 {
	return s.WaitingTime
}

func (s *GetJobInfoResponseBodyData) SetCuUsage(v int64) *GetJobInfoResponseBodyData {
	s.CuUsage = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetEndAtTime(v int64) *GetJobInfoResponseBodyData {
	s.EndAtTime = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetExtNodeId(v string) *GetJobInfoResponseBodyData {
	s.ExtNodeId = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetExtNodeOnDuty(v string) *GetJobInfoResponseBodyData {
	s.ExtNodeOnDuty = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetExtPlantFrom(v string) *GetJobInfoResponseBodyData {
	s.ExtPlantFrom = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetInputBytes(v float64) *GetJobInfoResponseBodyData {
	s.InputBytes = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetInstanceId(v string) *GetJobInfoResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetJobOwner(v string) *GetJobInfoResponseBodyData {
	s.JobOwner = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetJobSubStatusList(v []*GetJobInfoResponseBodyDataJobSubStatusList) *GetJobInfoResponseBodyData {
	s.JobSubStatusList = v
	return s
}

func (s *GetJobInfoResponseBodyData) SetJobType(v string) *GetJobInfoResponseBodyData {
	s.JobType = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetMemoryUsage(v int64) *GetJobInfoResponseBodyData {
	s.MemoryUsage = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetPriority(v int64) *GetJobInfoResponseBodyData {
	s.Priority = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetProject(v string) *GetJobInfoResponseBodyData {
	s.Project = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetQuotaNickname(v string) *GetJobInfoResponseBodyData {
	s.QuotaNickname = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetQuotaType(v string) *GetJobInfoResponseBodyData {
	s.QuotaType = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetRegion(v string) *GetJobInfoResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetRunningAtTime(v int64) *GetJobInfoResponseBodyData {
	s.RunningAtTime = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetRunningTime(v int64) *GetJobInfoResponseBodyData {
	s.RunningTime = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetSceneResults(v []*GetJobInfoResponseBodyDataSceneResults) *GetJobInfoResponseBodyData {
	s.SceneResults = v
	return s
}

func (s *GetJobInfoResponseBodyData) SetSignature(v string) *GetJobInfoResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetStatus(v string) *GetJobInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetSubmittedAtTime(v int64) *GetJobInfoResponseBodyData {
	s.SubmittedAtTime = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetTenantId(v string) *GetJobInfoResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetTotalTime(v int64) *GetJobInfoResponseBodyData {
	s.TotalTime = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetWaitingTime(v int64) *GetJobInfoResponseBodyData {
	s.WaitingTime = &v
	return s
}

func (s *GetJobInfoResponseBodyData) Validate() error {
	if s.JobSubStatusList != nil {
		for _, item := range s.JobSubStatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SceneResults != nil {
		for _, item := range s.SceneResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetJobInfoResponseBodyDataJobSubStatusList struct {
	Code        *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	StartTime   *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetJobInfoResponseBodyDataJobSubStatusList) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponseBodyDataJobSubStatusList) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyDataJobSubStatusList) GetCode() *int32 {
	return s.Code
}

func (s *GetJobInfoResponseBodyDataJobSubStatusList) GetDescription() *string {
	return s.Description
}

func (s *GetJobInfoResponseBodyDataJobSubStatusList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetJobInfoResponseBodyDataJobSubStatusList) SetCode(v int32) *GetJobInfoResponseBodyDataJobSubStatusList {
	s.Code = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobSubStatusList) SetDescription(v string) *GetJobInfoResponseBodyDataJobSubStatusList {
	s.Description = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobSubStatusList) SetStartTime(v string) *GetJobInfoResponseBodyDataJobSubStatusList {
	s.StartTime = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobSubStatusList) Validate() error {
	return dara.Validate(s)
}

type GetJobInfoResponseBodyDataSceneResults struct {
	Description *string            `json:"description,omitempty" xml:"description,omitempty"`
	Params      map[string]*string `json:"params,omitempty" xml:"params,omitempty"`
	Scene       *string            `json:"scene,omitempty" xml:"scene,omitempty"`
	SceneTag    *string            `json:"sceneTag,omitempty" xml:"sceneTag,omitempty"`
	Summary     *string            `json:"summary,omitempty" xml:"summary,omitempty"`
	Type        *string            `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetJobInfoResponseBodyDataSceneResults) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponseBodyDataSceneResults) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyDataSceneResults) GetDescription() *string {
	return s.Description
}

func (s *GetJobInfoResponseBodyDataSceneResults) GetParams() map[string]*string {
	return s.Params
}

func (s *GetJobInfoResponseBodyDataSceneResults) GetScene() *string {
	return s.Scene
}

func (s *GetJobInfoResponseBodyDataSceneResults) GetSceneTag() *string {
	return s.SceneTag
}

func (s *GetJobInfoResponseBodyDataSceneResults) GetSummary() *string {
	return s.Summary
}

func (s *GetJobInfoResponseBodyDataSceneResults) GetType() *string {
	return s.Type
}

func (s *GetJobInfoResponseBodyDataSceneResults) SetDescription(v string) *GetJobInfoResponseBodyDataSceneResults {
	s.Description = &v
	return s
}

func (s *GetJobInfoResponseBodyDataSceneResults) SetParams(v map[string]*string) *GetJobInfoResponseBodyDataSceneResults {
	s.Params = v
	return s
}

func (s *GetJobInfoResponseBodyDataSceneResults) SetScene(v string) *GetJobInfoResponseBodyDataSceneResults {
	s.Scene = &v
	return s
}

func (s *GetJobInfoResponseBodyDataSceneResults) SetSceneTag(v string) *GetJobInfoResponseBodyDataSceneResults {
	s.SceneTag = &v
	return s
}

func (s *GetJobInfoResponseBodyDataSceneResults) SetSummary(v string) *GetJobInfoResponseBodyDataSceneResults {
	s.Summary = &v
	return s
}

func (s *GetJobInfoResponseBodyDataSceneResults) SetType(v string) *GetJobInfoResponseBodyDataSceneResults {
	s.Type = &v
	return s
}

func (s *GetJobInfoResponseBodyDataSceneResults) Validate() error {
	return dara.Validate(s)
}
