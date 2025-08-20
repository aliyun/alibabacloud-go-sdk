// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobRunsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfigs(v string) *ListJobRunsShrinkRequest
	GetApplicationConfigs() *string
	SetCreator(v string) *ListJobRunsShrinkRequest
	GetCreator() *string
	SetEndTimeShrink(v string) *ListJobRunsShrinkRequest
	GetEndTimeShrink() *string
	SetIsWorkflow(v string) *ListJobRunsShrinkRequest
	GetIsWorkflow() *string
	SetJobRunDeploymentId(v string) *ListJobRunsShrinkRequest
	GetJobRunDeploymentId() *string
	SetJobRunId(v string) *ListJobRunsShrinkRequest
	GetJobRunId() *string
	SetMaxResults(v int32) *ListJobRunsShrinkRequest
	GetMaxResults() *int32
	SetMinDuration(v int64) *ListJobRunsShrinkRequest
	GetMinDuration() *int64
	SetName(v string) *ListJobRunsShrinkRequest
	GetName() *string
	SetNextToken(v string) *ListJobRunsShrinkRequest
	GetNextToken() *string
	SetRegionId(v string) *ListJobRunsShrinkRequest
	GetRegionId() *string
	SetResourceQueueId(v string) *ListJobRunsShrinkRequest
	GetResourceQueueId() *string
	SetRuntimeConfigs(v string) *ListJobRunsShrinkRequest
	GetRuntimeConfigs() *string
	SetStartTimeShrink(v string) *ListJobRunsShrinkRequest
	GetStartTimeShrink() *string
	SetStatesShrink(v string) *ListJobRunsShrinkRequest
	GetStatesShrink() *string
	SetTagsShrink(v string) *ListJobRunsShrinkRequest
	GetTagsShrink() *string
}

type ListJobRunsShrinkRequest struct {
	ApplicationConfigs *string `json:"applicationConfigs,omitempty" xml:"applicationConfigs,omitempty"`
	// The ID of the user who created the job.
	//
	// example:
	//
	// 1509789347011222
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The range of end time.
	EndTimeShrink *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	IsWorkflow    *string `json:"isWorkflow,omitempty" xml:"isWorkflow,omitempty"`
	// The job run ID.
	//
	// example:
	//
	// jd-b6d003f1930f****
	JobRunDeploymentId *string `json:"jobRunDeploymentId,omitempty" xml:"jobRunDeploymentId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// j-xxx
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The minimum running duration of the job. Unit: ms.
	//
	// example:
	//
	// 60000
	MinDuration *int64 `json:"minDuration,omitempty" xml:"minDuration,omitempty"`
	// The job name.
	//
	// example:
	//
	// emr-spark-demo-job
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The name of the resource queue on which the Spark jobs run.
	//
	// example:
	//
	// dev_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	RuntimeConfigs  *string `json:"runtimeConfigs,omitempty" xml:"runtimeConfigs,omitempty"`
	// The range of start time.
	StartTimeShrink *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The job states.
	//
	// example:
	//
	// ["Running","Submitted"]
	StatesShrink *string `json:"states,omitempty" xml:"states,omitempty"`
	// The tags of the job.
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListJobRunsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobRunsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListJobRunsShrinkRequest) GetApplicationConfigs() *string {
	return s.ApplicationConfigs
}

func (s *ListJobRunsShrinkRequest) GetCreator() *string {
	return s.Creator
}

func (s *ListJobRunsShrinkRequest) GetEndTimeShrink() *string {
	return s.EndTimeShrink
}

func (s *ListJobRunsShrinkRequest) GetIsWorkflow() *string {
	return s.IsWorkflow
}

func (s *ListJobRunsShrinkRequest) GetJobRunDeploymentId() *string {
	return s.JobRunDeploymentId
}

func (s *ListJobRunsShrinkRequest) GetJobRunId() *string {
	return s.JobRunId
}

func (s *ListJobRunsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListJobRunsShrinkRequest) GetMinDuration() *int64 {
	return s.MinDuration
}

func (s *ListJobRunsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListJobRunsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListJobRunsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListJobRunsShrinkRequest) GetResourceQueueId() *string {
	return s.ResourceQueueId
}

func (s *ListJobRunsShrinkRequest) GetRuntimeConfigs() *string {
	return s.RuntimeConfigs
}

func (s *ListJobRunsShrinkRequest) GetStartTimeShrink() *string {
	return s.StartTimeShrink
}

func (s *ListJobRunsShrinkRequest) GetStatesShrink() *string {
	return s.StatesShrink
}

func (s *ListJobRunsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListJobRunsShrinkRequest) SetApplicationConfigs(v string) *ListJobRunsShrinkRequest {
	s.ApplicationConfigs = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetCreator(v string) *ListJobRunsShrinkRequest {
	s.Creator = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetEndTimeShrink(v string) *ListJobRunsShrinkRequest {
	s.EndTimeShrink = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetIsWorkflow(v string) *ListJobRunsShrinkRequest {
	s.IsWorkflow = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetJobRunDeploymentId(v string) *ListJobRunsShrinkRequest {
	s.JobRunDeploymentId = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetJobRunId(v string) *ListJobRunsShrinkRequest {
	s.JobRunId = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetMaxResults(v int32) *ListJobRunsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetMinDuration(v int64) *ListJobRunsShrinkRequest {
	s.MinDuration = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetName(v string) *ListJobRunsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetNextToken(v string) *ListJobRunsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetRegionId(v string) *ListJobRunsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetResourceQueueId(v string) *ListJobRunsShrinkRequest {
	s.ResourceQueueId = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetRuntimeConfigs(v string) *ListJobRunsShrinkRequest {
	s.RuntimeConfigs = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetStartTimeShrink(v string) *ListJobRunsShrinkRequest {
	s.StartTimeShrink = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetStatesShrink(v string) *ListJobRunsShrinkRequest {
	s.StatesShrink = &v
	return s
}

func (s *ListJobRunsShrinkRequest) SetTagsShrink(v string) *ListJobRunsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListJobRunsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
