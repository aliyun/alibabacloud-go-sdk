// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleConsistencyJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v string) *CreateSampleConsistencyJobRequest
	GetDuration() *string
	SetEasModelServiceName(v string) *CreateSampleConsistencyJobRequest
	GetEasModelServiceName() *string
	SetFeatureSaveResourceId(v string) *CreateSampleConsistencyJobRequest
	GetFeatureSaveResourceId() *string
	SetInstanceId(v string) *CreateSampleConsistencyJobRequest
	GetInstanceId() *string
	SetItemIdField(v string) *CreateSampleConsistencyJobRequest
	GetItemIdField() *string
	SetName(v string) *CreateSampleConsistencyJobRequest
	GetName() *string
	SetPartitionField(v string) *CreateSampleConsistencyJobRequest
	GetPartitionField() *string
	SetPartitionFieldFormat(v string) *CreateSampleConsistencyJobRequest
	GetPartitionFieldFormat() *string
	SetRequestIdField(v string) *CreateSampleConsistencyJobRequest
	GetRequestIdField() *string
	SetSampleRate(v string) *CreateSampleConsistencyJobRequest
	GetSampleRate() *string
	SetSampleTableName(v string) *CreateSampleConsistencyJobRequest
	GetSampleTableName() *string
	SetSceneId(v string) *CreateSampleConsistencyJobRequest
	GetSceneId() *string
	SetUserIdField(v string) *CreateSampleConsistencyJobRequest
	GetUserIdField() *string
}

type CreateSampleConsistencyJobRequest struct {
	Duration              *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EasModelServiceName   *string `json:"EasModelServiceName,omitempty" xml:"EasModelServiceName,omitempty"`
	FeatureSaveResourceId *string `json:"FeatureSaveResourceId,omitempty" xml:"FeatureSaveResourceId,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ItemIdField          *string `json:"ItemIdField,omitempty" xml:"ItemIdField,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PartitionField       *string `json:"PartitionField,omitempty" xml:"PartitionField,omitempty"`
	PartitionFieldFormat *string `json:"PartitionFieldFormat,omitempty" xml:"PartitionFieldFormat,omitempty"`
	RequestIdField       *string `json:"RequestIdField,omitempty" xml:"RequestIdField,omitempty"`
	SampleRate           *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	SampleTableName      *string `json:"SampleTableName,omitempty" xml:"SampleTableName,omitempty"`
	SceneId              *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	UserIdField          *string `json:"UserIdField,omitempty" xml:"UserIdField,omitempty"`
}

func (s CreateSampleConsistencyJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleConsistencyJobRequest) GoString() string {
	return s.String()
}

func (s *CreateSampleConsistencyJobRequest) GetDuration() *string {
	return s.Duration
}

func (s *CreateSampleConsistencyJobRequest) GetEasModelServiceName() *string {
	return s.EasModelServiceName
}

func (s *CreateSampleConsistencyJobRequest) GetFeatureSaveResourceId() *string {
	return s.FeatureSaveResourceId
}

func (s *CreateSampleConsistencyJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSampleConsistencyJobRequest) GetItemIdField() *string {
	return s.ItemIdField
}

func (s *CreateSampleConsistencyJobRequest) GetName() *string {
	return s.Name
}

func (s *CreateSampleConsistencyJobRequest) GetPartitionField() *string {
	return s.PartitionField
}

func (s *CreateSampleConsistencyJobRequest) GetPartitionFieldFormat() *string {
	return s.PartitionFieldFormat
}

func (s *CreateSampleConsistencyJobRequest) GetRequestIdField() *string {
	return s.RequestIdField
}

func (s *CreateSampleConsistencyJobRequest) GetSampleRate() *string {
	return s.SampleRate
}

func (s *CreateSampleConsistencyJobRequest) GetSampleTableName() *string {
	return s.SampleTableName
}

func (s *CreateSampleConsistencyJobRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *CreateSampleConsistencyJobRequest) GetUserIdField() *string {
	return s.UserIdField
}

func (s *CreateSampleConsistencyJobRequest) SetDuration(v string) *CreateSampleConsistencyJobRequest {
	s.Duration = &v
	return s
}

func (s *CreateSampleConsistencyJobRequest) SetEasModelServiceName(v string) *CreateSampleConsistencyJobRequest {
	s.EasModelServiceName = &v
	return s
}

func (s *CreateSampleConsistencyJobRequest) SetFeatureSaveResourceId(v string) *CreateSampleConsistencyJobRequest {
	s.FeatureSaveResourceId = &v
	return s
}

func (s *CreateSampleConsistencyJobRequest) SetInstanceId(v string) *CreateSampleConsistencyJobRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSampleConsistencyJobRequest) SetItemIdField(v string) *CreateSampleConsistencyJobRequest {
	s.ItemIdField = &v
	return s
}

func (s *CreateSampleConsistencyJobRequest) SetName(v string) *CreateSampleConsistencyJobRequest {
	s.Name = &v
	return s
}

func (s *CreateSampleConsistencyJobRequest) SetPartitionField(v string) *CreateSampleConsistencyJobRequest {
	s.PartitionField = &v
	return s
}

func (s *CreateSampleConsistencyJobRequest) SetPartitionFieldFormat(v string) *CreateSampleConsistencyJobRequest {
	s.PartitionFieldFormat = &v
	return s
}

func (s *CreateSampleConsistencyJobRequest) SetRequestIdField(v string) *CreateSampleConsistencyJobRequest {
	s.RequestIdField = &v
	return s
}

func (s *CreateSampleConsistencyJobRequest) SetSampleRate(v string) *CreateSampleConsistencyJobRequest {
	s.SampleRate = &v
	return s
}

func (s *CreateSampleConsistencyJobRequest) SetSampleTableName(v string) *CreateSampleConsistencyJobRequest {
	s.SampleTableName = &v
	return s
}

func (s *CreateSampleConsistencyJobRequest) SetSceneId(v string) *CreateSampleConsistencyJobRequest {
	s.SceneId = &v
	return s
}

func (s *CreateSampleConsistencyJobRequest) SetUserIdField(v string) *CreateSampleConsistencyJobRequest {
	s.UserIdField = &v
	return s
}

func (s *CreateSampleConsistencyJobRequest) Validate() error {
	return dara.Validate(s)
}
