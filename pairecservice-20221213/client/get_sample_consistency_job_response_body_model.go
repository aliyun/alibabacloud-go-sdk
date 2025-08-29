// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSampleConsistencyJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *GetSampleConsistencyJobResponseBody
	GetConfig() *string
	SetDuration(v string) *GetSampleConsistencyJobResponseBody
	GetDuration() *string
	SetEasModelServiceName(v string) *GetSampleConsistencyJobResponseBody
	GetEasModelServiceName() *string
	SetEndTime(v string) *GetSampleConsistencyJobResponseBody
	GetEndTime() *string
	SetFeatureSaveResourceId(v string) *GetSampleConsistencyJobResponseBody
	GetFeatureSaveResourceId() *string
	SetGmtCreateTime(v string) *GetSampleConsistencyJobResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetSampleConsistencyJobResponseBody
	GetGmtModifiedTime() *string
	SetItemIdField(v string) *GetSampleConsistencyJobResponseBody
	GetItemIdField() *string
	SetLogs(v string) *GetSampleConsistencyJobResponseBody
	GetLogs() *string
	SetName(v string) *GetSampleConsistencyJobResponseBody
	GetName() *string
	SetPartitionField(v string) *GetSampleConsistencyJobResponseBody
	GetPartitionField() *string
	SetPartitionFieldFormat(v string) *GetSampleConsistencyJobResponseBody
	GetPartitionFieldFormat() *string
	SetRequestId(v string) *GetSampleConsistencyJobResponseBody
	GetRequestId() *string
	SetRequestIdField(v string) *GetSampleConsistencyJobResponseBody
	GetRequestIdField() *string
	SetSampleConsistencyJobId(v string) *GetSampleConsistencyJobResponseBody
	GetSampleConsistencyJobId() *string
	SetSampleRate(v string) *GetSampleConsistencyJobResponseBody
	GetSampleRate() *string
	SetSampleTableName(v string) *GetSampleConsistencyJobResponseBody
	GetSampleTableName() *string
	SetSceneId(v string) *GetSampleConsistencyJobResponseBody
	GetSceneId() *string
	SetSceneName(v string) *GetSampleConsistencyJobResponseBody
	GetSceneName() *string
	SetStartTime(v string) *GetSampleConsistencyJobResponseBody
	GetStartTime() *string
	SetStatus(v string) *GetSampleConsistencyJobResponseBody
	GetStatus() *string
	SetUserIdField(v string) *GetSampleConsistencyJobResponseBody
	GetUserIdField() *string
}

type GetSampleConsistencyJobResponseBody struct {
	Config                 *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Duration               *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EasModelServiceName    *string `json:"EasModelServiceName,omitempty" xml:"EasModelServiceName,omitempty"`
	EndTime                *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FeatureSaveResourceId  *string `json:"FeatureSaveResourceId,omitempty" xml:"FeatureSaveResourceId,omitempty"`
	GmtCreateTime          *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime        *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ItemIdField            *string `json:"ItemIdField,omitempty" xml:"ItemIdField,omitempty"`
	Logs                   *string `json:"Logs,omitempty" xml:"Logs,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PartitionField         *string `json:"PartitionField,omitempty" xml:"PartitionField,omitempty"`
	PartitionFieldFormat   *string `json:"PartitionFieldFormat,omitempty" xml:"PartitionFieldFormat,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestIdField         *string `json:"RequestIdField,omitempty" xml:"RequestIdField,omitempty"`
	SampleConsistencyJobId *string `json:"SampleConsistencyJobId,omitempty" xml:"SampleConsistencyJobId,omitempty"`
	SampleRate             *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	SampleTableName        *string `json:"SampleTableName,omitempty" xml:"SampleTableName,omitempty"`
	SceneId                *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName              *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	StartTime              *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserIdField            *string `json:"UserIdField,omitempty" xml:"UserIdField,omitempty"`
}

func (s GetSampleConsistencyJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSampleConsistencyJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetSampleConsistencyJobResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GetSampleConsistencyJobResponseBody) GetDuration() *string {
	return s.Duration
}

func (s *GetSampleConsistencyJobResponseBody) GetEasModelServiceName() *string {
	return s.EasModelServiceName
}

func (s *GetSampleConsistencyJobResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetSampleConsistencyJobResponseBody) GetFeatureSaveResourceId() *string {
	return s.FeatureSaveResourceId
}

func (s *GetSampleConsistencyJobResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetSampleConsistencyJobResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetSampleConsistencyJobResponseBody) GetItemIdField() *string {
	return s.ItemIdField
}

func (s *GetSampleConsistencyJobResponseBody) GetLogs() *string {
	return s.Logs
}

func (s *GetSampleConsistencyJobResponseBody) GetName() *string {
	return s.Name
}

func (s *GetSampleConsistencyJobResponseBody) GetPartitionField() *string {
	return s.PartitionField
}

func (s *GetSampleConsistencyJobResponseBody) GetPartitionFieldFormat() *string {
	return s.PartitionFieldFormat
}

func (s *GetSampleConsistencyJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSampleConsistencyJobResponseBody) GetRequestIdField() *string {
	return s.RequestIdField
}

func (s *GetSampleConsistencyJobResponseBody) GetSampleConsistencyJobId() *string {
	return s.SampleConsistencyJobId
}

func (s *GetSampleConsistencyJobResponseBody) GetSampleRate() *string {
	return s.SampleRate
}

func (s *GetSampleConsistencyJobResponseBody) GetSampleTableName() *string {
	return s.SampleTableName
}

func (s *GetSampleConsistencyJobResponseBody) GetSceneId() *string {
	return s.SceneId
}

func (s *GetSampleConsistencyJobResponseBody) GetSceneName() *string {
	return s.SceneName
}

func (s *GetSampleConsistencyJobResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetSampleConsistencyJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSampleConsistencyJobResponseBody) GetUserIdField() *string {
	return s.UserIdField
}

func (s *GetSampleConsistencyJobResponseBody) SetConfig(v string) *GetSampleConsistencyJobResponseBody {
	s.Config = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetDuration(v string) *GetSampleConsistencyJobResponseBody {
	s.Duration = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetEasModelServiceName(v string) *GetSampleConsistencyJobResponseBody {
	s.EasModelServiceName = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetEndTime(v string) *GetSampleConsistencyJobResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetFeatureSaveResourceId(v string) *GetSampleConsistencyJobResponseBody {
	s.FeatureSaveResourceId = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetGmtCreateTime(v string) *GetSampleConsistencyJobResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetGmtModifiedTime(v string) *GetSampleConsistencyJobResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetItemIdField(v string) *GetSampleConsistencyJobResponseBody {
	s.ItemIdField = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetLogs(v string) *GetSampleConsistencyJobResponseBody {
	s.Logs = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetName(v string) *GetSampleConsistencyJobResponseBody {
	s.Name = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetPartitionField(v string) *GetSampleConsistencyJobResponseBody {
	s.PartitionField = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetPartitionFieldFormat(v string) *GetSampleConsistencyJobResponseBody {
	s.PartitionFieldFormat = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetRequestId(v string) *GetSampleConsistencyJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetRequestIdField(v string) *GetSampleConsistencyJobResponseBody {
	s.RequestIdField = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetSampleConsistencyJobId(v string) *GetSampleConsistencyJobResponseBody {
	s.SampleConsistencyJobId = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetSampleRate(v string) *GetSampleConsistencyJobResponseBody {
	s.SampleRate = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetSampleTableName(v string) *GetSampleConsistencyJobResponseBody {
	s.SampleTableName = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetSceneId(v string) *GetSampleConsistencyJobResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetSceneName(v string) *GetSampleConsistencyJobResponseBody {
	s.SceneName = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetStartTime(v string) *GetSampleConsistencyJobResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetStatus(v string) *GetSampleConsistencyJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) SetUserIdField(v string) *GetSampleConsistencyJobResponseBody {
	s.UserIdField = &v
	return s
}

func (s *GetSampleConsistencyJobResponseBody) Validate() error {
	return dara.Validate(s)
}
