// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSampleConsistencyJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSampleConsistencyJobsResponseBody
	GetRequestId() *string
	SetSampleConsistencyJobs(v []*ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) *ListSampleConsistencyJobsResponseBody
	GetSampleConsistencyJobs() []*ListSampleConsistencyJobsResponseBodySampleConsistencyJobs
	SetTotalCount(v int64) *ListSampleConsistencyJobsResponseBody
	GetTotalCount() *int64
}

type ListSampleConsistencyJobsResponseBody struct {
	RequestId             *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SampleConsistencyJobs []*ListSampleConsistencyJobsResponseBodySampleConsistencyJobs `json:"SampleConsistencyJobs,omitempty" xml:"SampleConsistencyJobs,omitempty" type:"Repeated"`
	TotalCount            *int64                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSampleConsistencyJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSampleConsistencyJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSampleConsistencyJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSampleConsistencyJobsResponseBody) GetSampleConsistencyJobs() []*ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	return s.SampleConsistencyJobs
}

func (s *ListSampleConsistencyJobsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSampleConsistencyJobsResponseBody) SetRequestId(v string) *ListSampleConsistencyJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBody) SetSampleConsistencyJobs(v []*ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) *ListSampleConsistencyJobsResponseBody {
	s.SampleConsistencyJobs = v
	return s
}

func (s *ListSampleConsistencyJobsResponseBody) SetTotalCount(v int64) *ListSampleConsistencyJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSampleConsistencyJobsResponseBodySampleConsistencyJobs struct {
	Config                 *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Duration               *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EasModelServiceName    *string `json:"EasModelServiceName,omitempty" xml:"EasModelServiceName,omitempty"`
	EndTime                *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FeatureSaveResourceId  *string `json:"FeatureSaveResourceId,omitempty" xml:"FeatureSaveResourceId,omitempty"`
	GmtCreateTime          *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime        *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ItemIdField            *string `json:"ItemIdField,omitempty" xml:"ItemIdField,omitempty"`
	Logs                   *string `json:"Logs,omitempty" xml:"Logs,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PartitionField         *string `json:"PartitionField,omitempty" xml:"PartitionField,omitempty"`
	PartitionFieldFormat   *string `json:"PartitionFieldFormat,omitempty" xml:"PartitionFieldFormat,omitempty"`
	RequestIdField         *string `json:"RequestIdField,omitempty" xml:"RequestIdField,omitempty"`
	SampleConsistencyJobId *string `json:"SampleConsistencyJobId,omitempty" xml:"SampleConsistencyJobId,omitempty"`
	SampleRate             *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	SampleTableName        *string `json:"SampleTableName,omitempty" xml:"SampleTableName,omitempty"`
	SceneId                *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName              *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	StartTime              *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserIdField            *string `json:"UserIdField,omitempty" xml:"UserIdField,omitempty"`
}

func (s ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) String() string {
	return dara.Prettify(s)
}

func (s ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GoString() string {
	return s.String()
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetConfig() *string {
	return s.Config
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetDuration() *string {
	return s.Duration
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetEasModelServiceName() *string {
	return s.EasModelServiceName
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetFeatureSaveResourceId() *string {
	return s.FeatureSaveResourceId
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetItemIdField() *string {
	return s.ItemIdField
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetLogs() *string {
	return s.Logs
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetName() *string {
	return s.Name
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetPartitionField() *string {
	return s.PartitionField
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetPartitionFieldFormat() *string {
	return s.PartitionFieldFormat
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetRequestIdField() *string {
	return s.RequestIdField
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetSampleConsistencyJobId() *string {
	return s.SampleConsistencyJobId
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetSampleRate() *string {
	return s.SampleRate
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetSampleTableName() *string {
	return s.SampleTableName
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetSceneId() *string {
	return s.SceneId
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetSceneName() *string {
	return s.SceneName
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetStatus() *string {
	return s.Status
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) GetUserIdField() *string {
	return s.UserIdField
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetConfig(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.Config = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetDuration(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.Duration = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetEasModelServiceName(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.EasModelServiceName = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetEndTime(v int64) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.EndTime = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetFeatureSaveResourceId(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.FeatureSaveResourceId = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetGmtCreateTime(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.GmtCreateTime = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetGmtModifiedTime(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetItemIdField(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.ItemIdField = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetLogs(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.Logs = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetName(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.Name = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetPartitionField(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.PartitionField = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetPartitionFieldFormat(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.PartitionFieldFormat = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetRequestIdField(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.RequestIdField = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetSampleConsistencyJobId(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.SampleConsistencyJobId = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetSampleRate(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.SampleRate = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetSampleTableName(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.SampleTableName = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetSceneId(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.SceneId = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetSceneName(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.SceneName = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetStartTime(v int64) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.StartTime = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetStatus(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.Status = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) SetUserIdField(v string) *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs {
	s.UserIdField = &v
	return s
}

func (s *ListSampleConsistencyJobsResponseBodySampleConsistencyJobs) Validate() error {
	return dara.Validate(s)
}
