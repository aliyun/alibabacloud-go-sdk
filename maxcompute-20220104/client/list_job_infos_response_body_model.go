// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListJobInfosResponseBodyData) *ListJobInfosResponseBody
	GetData() *ListJobInfosResponseBodyData
	SetHttpCode(v int32) *ListJobInfosResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListJobInfosResponseBody
	GetRequestId() *string
}

type ListJobInfosResponseBody struct {
	// The data returned.
	Data *ListJobInfosResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Indicates whether the request was successful. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc13a9516807484336515320e38f5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListJobInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobInfosResponseBody) GetData() *ListJobInfosResponseBodyData {
	return s.Data
}

func (s *ListJobInfosResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListJobInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobInfosResponseBody) SetData(v *ListJobInfosResponseBodyData) *ListJobInfosResponseBody {
	s.Data = v
	return s
}

func (s *ListJobInfosResponseBody) SetHttpCode(v int32) *ListJobInfosResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListJobInfosResponseBody) SetRequestId(v string) *ListJobInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobInfosResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobInfosResponseBodyData struct {
	// The information about the jobs.
	JobInfoList []*ListJobInfosResponseBodyDataJobInfoList `json:"jobInfoList,omitempty" xml:"jobInfoList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 2
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 64
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListJobInfosResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListJobInfosResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobInfosResponseBodyData) GetJobInfoList() []*ListJobInfosResponseBodyDataJobInfoList {
	return s.JobInfoList
}

func (s *ListJobInfosResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListJobInfosResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListJobInfosResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListJobInfosResponseBodyData) SetJobInfoList(v []*ListJobInfosResponseBodyDataJobInfoList) *ListJobInfosResponseBodyData {
	s.JobInfoList = v
	return s
}

func (s *ListJobInfosResponseBodyData) SetPageNumber(v int64) *ListJobInfosResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListJobInfosResponseBodyData) SetPageSize(v int64) *ListJobInfosResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListJobInfosResponseBodyData) SetTotalCount(v int64) *ListJobInfosResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListJobInfosResponseBodyData) Validate() error {
	if s.JobInfoList != nil {
		for _, item := range s.JobInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobInfosResponseBodyDataJobInfoList struct {
	// The cluster ID.
	//
	// example:
	//
	// AY20A
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The CU snapshot proportion of the job.
	//
	// example:
	//
	// 0.48
	CuSnapshot *float64 `json:"cuSnapshot,omitempty" xml:"cuSnapshot,omitempty"`
	// The amount of resources consumed by the job. This parameter is returned only for jobs that are complete.Unit: 100\\*Core\\*s.
	//
	// example:
	//
	// 10
	CuUsage *int64 `json:"cuUsage,omitempty" xml:"cuUsage,omitempty"`
	// The time when the job stops running.
	//
	// example:
	//
	// 0
	EndAtTime *int64 `json:"endAtTime,omitempty" xml:"endAtTime,omitempty"`
	// The node ID of DataWorks.
	//
	// example:
	//
	// node_4
	ExtNodeId *string `json:"extNodeId,omitempty" xml:"extNodeId,omitempty"`
	// The account of the node owner.
	//
	// example:
	//
	// duty_2
	ExtNodeOnDuty *string `json:"extNodeOnDuty,omitempty" xml:"extNodeOnDuty,omitempty"`
	// The upstream platform.
	//
	// example:
	//
	// platform_3
	ExtPlantFrom *string `json:"extPlantFrom,omitempty" xml:"extPlantFrom,omitempty"`
	// The amount of scanned data for the job. Unit: byte.
	//
	// example:
	//
	// 1234
	InputBytes *float64 `json:"inputBytes,omitempty" xml:"inputBytes,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 20230410050036549gfmsdwf60gg
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The account that commits the job.
	//
	// example:
	//
	// ALIYUN$xxx@test.aliyunid.com
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// The type of the job.
	//
	// example:
	//
	// SQL
	JobType *string `json:"jobType,omitempty" xml:"jobType,omitempty"`
	// The memory snapshot proportion of the job.
	//
	// example:
	//
	// 0.42
	MemorySnapshot *float64 `json:"memorySnapshot,omitempty" xml:"memorySnapshot,omitempty"`
	// The number of memory consumed by the job. This parameter is returned only for jobs that are complete.Unit: MB\\*s.
	//
	// example:
	//
	// 40
	MemoryUsage *int64 `json:"memoryUsage,omitempty" xml:"memoryUsage,omitempty"`
	// The priority of the job.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// The name of the MaxCompute project.
	//
	// example:
	//
	// openrec_new
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The nickname of the quota that is used by the job.
	//
	// example:
	//
	// my_quota
	QuotaNickname *string `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
	// The type of the quota.
	//
	// example:
	//
	// subscription
	QuotaType *string `json:"quotaType,omitempty" xml:"quotaType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The time when the job starts to run.
	//
	// example:
	//
	// 1672112113
	RunningAtTime *int64 `json:"runningAtTime,omitempty" xml:"runningAtTime,omitempty"`
	// The period for which the job runs.
	//
	// example:
	//
	// 800
	RunningTime *int64 `json:"runningTime,omitempty" xml:"runningTime,omitempty"`
	// The intelligent diagnostics results.
	SceneResults []*ListJobInfosResponseBodyDataJobInfoListSceneResults `json:"sceneResults,omitempty" xml:"sceneResults,omitempty" type:"Repeated"`
	// The signature of the SQL job.
	//
	// example:
	//
	// i094KijGrN3kOXZ74kbexB77XQY=
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// The status of the job.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The status of the snapshot.
	//
	// example:
	//
	// running
	StatusSnapshot *string `json:"statusSnapshot,omitempty" xml:"statusSnapshot,omitempty"`
	// The time when the job was committed.
	//
	// example:
	//
	// 1672112013
	SubmittedAtTime *int64 `json:"submittedAtTime,omitempty" xml:"submittedAtTime,omitempty"`
	// The tags.
	//
	// example:
	//
	// []
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 213065738244354
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The total period for which the job runs.
	//
	// example:
	//
	// 900
	TotalTime *int64 `json:"totalTime,omitempty" xml:"totalTime,omitempty"`
	// The duration for which the job waits to start.
	//
	// example:
	//
	// 100
	WaitingTime *int64 `json:"waitingTime,omitempty" xml:"waitingTime,omitempty"`
}

func (s ListJobInfosResponseBodyDataJobInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListJobInfosResponseBodyDataJobInfoList) GoString() string {
	return s.String()
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetCluster() *string {
	return s.Cluster
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetCuSnapshot() *float64 {
	return s.CuSnapshot
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetCuUsage() *int64 {
	return s.CuUsage
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetEndAtTime() *int64 {
	return s.EndAtTime
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetExtNodeId() *string {
	return s.ExtNodeId
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetExtNodeOnDuty() *string {
	return s.ExtNodeOnDuty
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetExtPlantFrom() *string {
	return s.ExtPlantFrom
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetInputBytes() *float64 {
	return s.InputBytes
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetJobOwner() *string {
	return s.JobOwner
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetJobType() *string {
	return s.JobType
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetMemorySnapshot() *float64 {
	return s.MemorySnapshot
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetMemoryUsage() *int64 {
	return s.MemoryUsage
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetPriority() *int64 {
	return s.Priority
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetProject() *string {
	return s.Project
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetQuotaNickname() *string {
	return s.QuotaNickname
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetQuotaType() *string {
	return s.QuotaType
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetRegion() *string {
	return s.Region
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetRunningAtTime() *int64 {
	return s.RunningAtTime
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetRunningTime() *int64 {
	return s.RunningTime
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetSceneResults() []*ListJobInfosResponseBodyDataJobInfoListSceneResults {
	return s.SceneResults
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetSignature() *string {
	return s.Signature
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetStatus() *string {
	return s.Status
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetStatusSnapshot() *string {
	return s.StatusSnapshot
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetSubmittedAtTime() *int64 {
	return s.SubmittedAtTime
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetTags() *string {
	return s.Tags
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetTotalTime() *int64 {
	return s.TotalTime
}

func (s *ListJobInfosResponseBodyDataJobInfoList) GetWaitingTime() *int64 {
	return s.WaitingTime
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetCluster(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Cluster = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetCuSnapshot(v float64) *ListJobInfosResponseBodyDataJobInfoList {
	s.CuSnapshot = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetCuUsage(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.CuUsage = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetEndAtTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.EndAtTime = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetExtNodeId(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.ExtNodeId = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetExtNodeOnDuty(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.ExtNodeOnDuty = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetExtPlantFrom(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.ExtPlantFrom = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetInputBytes(v float64) *ListJobInfosResponseBodyDataJobInfoList {
	s.InputBytes = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetInstanceId(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.InstanceId = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetJobOwner(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.JobOwner = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetJobType(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.JobType = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetMemorySnapshot(v float64) *ListJobInfosResponseBodyDataJobInfoList {
	s.MemorySnapshot = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetMemoryUsage(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.MemoryUsage = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetPriority(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.Priority = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetProject(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Project = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetQuotaNickname(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.QuotaNickname = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetQuotaType(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.QuotaType = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetRegion(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Region = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetRunningAtTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.RunningAtTime = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetRunningTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.RunningTime = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetSceneResults(v []*ListJobInfosResponseBodyDataJobInfoListSceneResults) *ListJobInfosResponseBodyDataJobInfoList {
	s.SceneResults = v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetSignature(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Signature = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetStatus(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Status = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetStatusSnapshot(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.StatusSnapshot = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetSubmittedAtTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.SubmittedAtTime = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetTags(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Tags = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetTenantId(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.TenantId = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetTotalTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.TotalTime = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetWaitingTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.WaitingTime = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) Validate() error {
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

type ListJobInfosResponseBodyDataJobInfoListSceneResults struct {
	// The intelligent diagnostics result description.
	//
	// example:
	//
	// This job uses annual and monthly computing resources. It may be that the job is waiting for resources due to the large amount of overall job running data, many resources requested, and low job priority. Please go to Resource Consumption to view the specific situation. You can also go to Cost Optimization to see if you need to adjust resource configuration.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Information about the nodes where data skew or data expansion is detected. This parameter is returned only when the diagnostics scenario is data skew or data expansion.
	Params map[string]*string `json:"params,omitempty" xml:"params,omitempty"`
	// The intelligent diagnostics result scenario.
	//
	// example:
	//
	// LackResource
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The intelligent diagnostics result tag.
	//
	// example:
	//
	// SubscriptionLackResource
	SceneTag *string `json:"sceneTag,omitempty" xml:"sceneTag,omitempty"`
	// The intelligent diagnostics result summary.
	//
	// example:
	//
	// Insufficient computing resources available for the job. Click to view details.
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// The intelligent diagnostics result type.
	//
	// example:
	//
	// warning
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListJobInfosResponseBodyDataJobInfoListSceneResults) String() string {
	return dara.Prettify(s)
}

func (s ListJobInfosResponseBodyDataJobInfoListSceneResults) GoString() string {
	return s.String()
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) GetDescription() *string {
	return s.Description
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) GetParams() map[string]*string {
	return s.Params
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) GetScene() *string {
	return s.Scene
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) GetSceneTag() *string {
	return s.SceneTag
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) GetSummary() *string {
	return s.Summary
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) GetType() *string {
	return s.Type
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) SetDescription(v string) *ListJobInfosResponseBodyDataJobInfoListSceneResults {
	s.Description = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) SetParams(v map[string]*string) *ListJobInfosResponseBodyDataJobInfoListSceneResults {
	s.Params = v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) SetScene(v string) *ListJobInfosResponseBodyDataJobInfoListSceneResults {
	s.Scene = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) SetSceneTag(v string) *ListJobInfosResponseBodyDataJobInfoListSceneResults {
	s.SceneTag = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) SetSummary(v string) *ListJobInfosResponseBodyDataJobInfoListSceneResults {
	s.Summary = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) SetType(v string) *ListJobInfosResponseBodyDataJobInfoListSceneResults {
	s.Type = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoListSceneResults) Validate() error {
	return dara.Validate(s)
}
