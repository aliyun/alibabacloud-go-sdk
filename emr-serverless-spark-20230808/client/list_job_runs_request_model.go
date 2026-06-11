// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobRunsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfigs(v string) *ListJobRunsRequest
	GetApplicationConfigs() *string
	SetCreator(v string) *ListJobRunsRequest
	GetCreator() *string
	SetEndTime(v *ListJobRunsRequestEndTime) *ListJobRunsRequest
	GetEndTime() *ListJobRunsRequestEndTime
	SetIsWorkflow(v string) *ListJobRunsRequest
	GetIsWorkflow() *string
	SetJobRunDeploymentId(v string) *ListJobRunsRequest
	GetJobRunDeploymentId() *string
	SetJobRunId(v string) *ListJobRunsRequest
	GetJobRunId() *string
	SetMaxResults(v int32) *ListJobRunsRequest
	GetMaxResults() *int32
	SetMinDuration(v int64) *ListJobRunsRequest
	GetMinDuration() *int64
	SetName(v string) *ListJobRunsRequest
	GetName() *string
	SetNextToken(v string) *ListJobRunsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListJobRunsRequest
	GetRegionId() *string
	SetResourceQueueId(v string) *ListJobRunsRequest
	GetResourceQueueId() *string
	SetRuntimeConfigs(v string) *ListJobRunsRequest
	GetRuntimeConfigs() *string
	SetStartTime(v *ListJobRunsRequestStartTime) *ListJobRunsRequest
	GetStartTime() *ListJobRunsRequestStartTime
	SetStates(v []*string) *ListJobRunsRequest
	GetStates() []*string
	SetTags(v []*ListJobRunsRequestTags) *ListJobRunsRequest
	GetTags() []*ListJobRunsRequestTags
}

type ListJobRunsRequest struct {
	// The Spark configurations.
	//
	// example:
	//
	// [{\\"key\\":\\"spark.app.name\\",\\"value\\":\\"test\\"}]
	ApplicationConfigs *string `json:"applicationConfigs,omitempty" xml:"applicationConfigs,omitempty"`
	// The UID of the user who created the job.
	//
	// example:
	//
	// 150976534701****
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The time range when the job run ended.
	EndTime *ListJobRunsRequestEndTime `json:"endTime,omitempty" xml:"endTime,omitempty" type:"Struct"`
	// Specifies whether the job is a workflow task.
	//
	// example:
	//
	// false
	IsWorkflow *string `json:"isWorkflow,omitempty" xml:"isWorkflow,omitempty"`
	// The deployment ID of the streaming job.
	//
	// example:
	//
	// jd-b6d003f1930f****
	JobRunDeploymentId *string `json:"jobRunDeploymentId,omitempty" xml:"jobRunDeploymentId,omitempty"`
	// The job run ID.
	//
	// example:
	//
	// j-xxx
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// The maximum number of entries to return. The maximum value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The minimum runtime of the job run, in milliseconds.
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
	// The token that specifies the position from which to start the next read.
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
	// The ID of the resource queue on which the Spark job runs.
	//
	// example:
	//
	// dev_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	// The runtime configurations.
	//
	// example:
	//
	// [{\\"key\\":\\"mainClass\\",\\"value\\":\\"yourClass\\"}]
	RuntimeConfigs *string `json:"runtimeConfigs,omitempty" xml:"runtimeConfigs,omitempty"`
	// The time range when the job run started.
	StartTime *ListJobRunsRequestStartTime `json:"startTime,omitempty" xml:"startTime,omitempty" type:"Struct"`
	// The job run states.
	//
	// example:
	//
	// ["Running","Submitted"]
	States []*string `json:"states,omitempty" xml:"states,omitempty" type:"Repeated"`
	// The list of tags.
	Tags []*ListJobRunsRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListJobRunsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobRunsRequest) GoString() string {
	return s.String()
}

func (s *ListJobRunsRequest) GetApplicationConfigs() *string {
	return s.ApplicationConfigs
}

func (s *ListJobRunsRequest) GetCreator() *string {
	return s.Creator
}

func (s *ListJobRunsRequest) GetEndTime() *ListJobRunsRequestEndTime {
	return s.EndTime
}

func (s *ListJobRunsRequest) GetIsWorkflow() *string {
	return s.IsWorkflow
}

func (s *ListJobRunsRequest) GetJobRunDeploymentId() *string {
	return s.JobRunDeploymentId
}

func (s *ListJobRunsRequest) GetJobRunId() *string {
	return s.JobRunId
}

func (s *ListJobRunsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListJobRunsRequest) GetMinDuration() *int64 {
	return s.MinDuration
}

func (s *ListJobRunsRequest) GetName() *string {
	return s.Name
}

func (s *ListJobRunsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListJobRunsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListJobRunsRequest) GetResourceQueueId() *string {
	return s.ResourceQueueId
}

func (s *ListJobRunsRequest) GetRuntimeConfigs() *string {
	return s.RuntimeConfigs
}

func (s *ListJobRunsRequest) GetStartTime() *ListJobRunsRequestStartTime {
	return s.StartTime
}

func (s *ListJobRunsRequest) GetStates() []*string {
	return s.States
}

func (s *ListJobRunsRequest) GetTags() []*ListJobRunsRequestTags {
	return s.Tags
}

func (s *ListJobRunsRequest) SetApplicationConfigs(v string) *ListJobRunsRequest {
	s.ApplicationConfigs = &v
	return s
}

func (s *ListJobRunsRequest) SetCreator(v string) *ListJobRunsRequest {
	s.Creator = &v
	return s
}

func (s *ListJobRunsRequest) SetEndTime(v *ListJobRunsRequestEndTime) *ListJobRunsRequest {
	s.EndTime = v
	return s
}

func (s *ListJobRunsRequest) SetIsWorkflow(v string) *ListJobRunsRequest {
	s.IsWorkflow = &v
	return s
}

func (s *ListJobRunsRequest) SetJobRunDeploymentId(v string) *ListJobRunsRequest {
	s.JobRunDeploymentId = &v
	return s
}

func (s *ListJobRunsRequest) SetJobRunId(v string) *ListJobRunsRequest {
	s.JobRunId = &v
	return s
}

func (s *ListJobRunsRequest) SetMaxResults(v int32) *ListJobRunsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListJobRunsRequest) SetMinDuration(v int64) *ListJobRunsRequest {
	s.MinDuration = &v
	return s
}

func (s *ListJobRunsRequest) SetName(v string) *ListJobRunsRequest {
	s.Name = &v
	return s
}

func (s *ListJobRunsRequest) SetNextToken(v string) *ListJobRunsRequest {
	s.NextToken = &v
	return s
}

func (s *ListJobRunsRequest) SetRegionId(v string) *ListJobRunsRequest {
	s.RegionId = &v
	return s
}

func (s *ListJobRunsRequest) SetResourceQueueId(v string) *ListJobRunsRequest {
	s.ResourceQueueId = &v
	return s
}

func (s *ListJobRunsRequest) SetRuntimeConfigs(v string) *ListJobRunsRequest {
	s.RuntimeConfigs = &v
	return s
}

func (s *ListJobRunsRequest) SetStartTime(v *ListJobRunsRequestStartTime) *ListJobRunsRequest {
	s.StartTime = v
	return s
}

func (s *ListJobRunsRequest) SetStates(v []*string) *ListJobRunsRequest {
	s.States = v
	return s
}

func (s *ListJobRunsRequest) SetTags(v []*ListJobRunsRequestTags) *ListJobRunsRequest {
	s.Tags = v
	return s
}

func (s *ListJobRunsRequest) Validate() error {
	if s.EndTime != nil {
		if err := s.EndTime.Validate(); err != nil {
			return err
		}
	}
	if s.StartTime != nil {
		if err := s.StartTime.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobRunsRequestEndTime struct {
	// The end of the time range.
	//
	// example:
	//
	// 1710432000000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start of the time range.
	//
	// example:
	//
	// 1709740800000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListJobRunsRequestEndTime) String() string {
	return dara.Prettify(s)
}

func (s ListJobRunsRequestEndTime) GoString() string {
	return s.String()
}

func (s *ListJobRunsRequestEndTime) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListJobRunsRequestEndTime) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListJobRunsRequestEndTime) SetEndTime(v int64) *ListJobRunsRequestEndTime {
	s.EndTime = &v
	return s
}

func (s *ListJobRunsRequestEndTime) SetStartTime(v int64) *ListJobRunsRequestEndTime {
	s.StartTime = &v
	return s
}

func (s *ListJobRunsRequestEndTime) Validate() error {
	return dara.Validate(s)
}

type ListJobRunsRequestStartTime struct {
	// The end of the time range.
	//
	// example:
	//
	// 1710432000000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start of the time range.
	//
	// example:
	//
	// 1709740800000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListJobRunsRequestStartTime) String() string {
	return dara.Prettify(s)
}

func (s ListJobRunsRequestStartTime) GoString() string {
	return s.String()
}

func (s *ListJobRunsRequestStartTime) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListJobRunsRequestStartTime) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListJobRunsRequestStartTime) SetEndTime(v int64) *ListJobRunsRequestStartTime {
	s.EndTime = &v
	return s
}

func (s *ListJobRunsRequestStartTime) SetStartTime(v int64) *ListJobRunsRequestStartTime {
	s.StartTime = &v
	return s
}

func (s *ListJobRunsRequestStartTime) Validate() error {
	return dara.Validate(s)
}

type ListJobRunsRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// tag_key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListJobRunsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListJobRunsRequestTags) GoString() string {
	return s.String()
}

func (s *ListJobRunsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListJobRunsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListJobRunsRequestTags) SetKey(v string) *ListJobRunsRequestTags {
	s.Key = &v
	return s
}

func (s *ListJobRunsRequestTags) SetValue(v string) *ListJobRunsRequestTags {
	s.Value = &v
	return s
}

func (s *ListJobRunsRequestTags) Validate() error {
	return dara.Validate(s)
}
