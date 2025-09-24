// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelVersion interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalStatus(v string) *ModelVersion
	GetApprovalStatus() *string
	SetCompressionSpec(v map[string]interface{}) *ModelVersion
	GetCompressionSpec() map[string]interface{}
	SetDistillationSpec(v map[string]interface{}) *ModelVersion
	GetDistillationSpec() map[string]interface{}
	SetEvaluationSpec(v map[string]interface{}) *ModelVersion
	GetEvaluationSpec() map[string]interface{}
	SetExtraInfo(v map[string]interface{}) *ModelVersion
	GetExtraInfo() map[string]interface{}
	SetFormatType(v string) *ModelVersion
	GetFormatType() *string
	SetFrameworkType(v string) *ModelVersion
	GetFrameworkType() *string
	SetGmtCreateTime(v string) *ModelVersion
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *ModelVersion
	GetGmtModifiedTime() *string
	SetInferenceSpec(v map[string]interface{}) *ModelVersion
	GetInferenceSpec() map[string]interface{}
	SetLabels(v []*ModelVersionLabels) *ModelVersion
	GetLabels() []*ModelVersionLabels
	SetMetrics(v map[string]interface{}) *ModelVersion
	GetMetrics() map[string]interface{}
	SetOptions(v string) *ModelVersion
	GetOptions() *string
	SetOwnerId(v string) *ModelVersion
	GetOwnerId() *string
	SetSourceId(v string) *ModelVersion
	GetSourceId() *string
	SetSourceType(v string) *ModelVersion
	GetSourceType() *string
	SetTrainingSpec(v map[string]interface{}) *ModelVersion
	GetTrainingSpec() map[string]interface{}
	SetUri(v string) *ModelVersion
	GetUri() *string
	SetUserId(v string) *ModelVersion
	GetUserId() *string
	SetVersionDescription(v string) *ModelVersion
	GetVersionDescription() *string
	SetVersionName(v string) *ModelVersion
	GetVersionName() *string
}

type ModelVersion struct {
	ApprovalStatus   *string                `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	CompressionSpec  map[string]interface{} `json:"CompressionSpec,omitempty" xml:"CompressionSpec,omitempty"`
	DistillationSpec map[string]interface{} `json:"DistillationSpec,omitempty" xml:"DistillationSpec,omitempty"`
	EvaluationSpec   map[string]interface{} `json:"EvaluationSpec,omitempty" xml:"EvaluationSpec,omitempty"`
	ExtraInfo        map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// example:
	//
	// SavedModel
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// example:
	//
	// TensorFlow
	FrameworkType *string `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtModifiedTime *string                `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	InferenceSpec   map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	Labels          []*ModelVersionLabels  `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Metrics         map[string]interface{} `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	Options         *string                `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// 155770209******
	OwnerId      *string                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SourceId     *string                `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType   *string                `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	TrainingSpec map[string]interface{} `json:"TrainingSpec,omitempty" xml:"TrainingSpec,omitempty"`
	// example:
	//
	// oss://bucket/path-to-model
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// example:
	//
	// 155770209******
	UserId             *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	// example:
	//
	// 1.0.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ModelVersion) String() string {
	return dara.Prettify(s)
}

func (s ModelVersion) GoString() string {
	return s.String()
}

func (s *ModelVersion) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *ModelVersion) GetCompressionSpec() map[string]interface{} {
	return s.CompressionSpec
}

func (s *ModelVersion) GetDistillationSpec() map[string]interface{} {
	return s.DistillationSpec
}

func (s *ModelVersion) GetEvaluationSpec() map[string]interface{} {
	return s.EvaluationSpec
}

func (s *ModelVersion) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *ModelVersion) GetFormatType() *string {
	return s.FormatType
}

func (s *ModelVersion) GetFrameworkType() *string {
	return s.FrameworkType
}

func (s *ModelVersion) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ModelVersion) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ModelVersion) GetInferenceSpec() map[string]interface{} {
	return s.InferenceSpec
}

func (s *ModelVersion) GetLabels() []*ModelVersionLabels {
	return s.Labels
}

func (s *ModelVersion) GetMetrics() map[string]interface{} {
	return s.Metrics
}

func (s *ModelVersion) GetOptions() *string {
	return s.Options
}

func (s *ModelVersion) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModelVersion) GetSourceId() *string {
	return s.SourceId
}

func (s *ModelVersion) GetSourceType() *string {
	return s.SourceType
}

func (s *ModelVersion) GetTrainingSpec() map[string]interface{} {
	return s.TrainingSpec
}

func (s *ModelVersion) GetUri() *string {
	return s.Uri
}

func (s *ModelVersion) GetUserId() *string {
	return s.UserId
}

func (s *ModelVersion) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *ModelVersion) GetVersionName() *string {
	return s.VersionName
}

func (s *ModelVersion) SetApprovalStatus(v string) *ModelVersion {
	s.ApprovalStatus = &v
	return s
}

func (s *ModelVersion) SetCompressionSpec(v map[string]interface{}) *ModelVersion {
	s.CompressionSpec = v
	return s
}

func (s *ModelVersion) SetDistillationSpec(v map[string]interface{}) *ModelVersion {
	s.DistillationSpec = v
	return s
}

func (s *ModelVersion) SetEvaluationSpec(v map[string]interface{}) *ModelVersion {
	s.EvaluationSpec = v
	return s
}

func (s *ModelVersion) SetExtraInfo(v map[string]interface{}) *ModelVersion {
	s.ExtraInfo = v
	return s
}

func (s *ModelVersion) SetFormatType(v string) *ModelVersion {
	s.FormatType = &v
	return s
}

func (s *ModelVersion) SetFrameworkType(v string) *ModelVersion {
	s.FrameworkType = &v
	return s
}

func (s *ModelVersion) SetGmtCreateTime(v string) *ModelVersion {
	s.GmtCreateTime = &v
	return s
}

func (s *ModelVersion) SetGmtModifiedTime(v string) *ModelVersion {
	s.GmtModifiedTime = &v
	return s
}

func (s *ModelVersion) SetInferenceSpec(v map[string]interface{}) *ModelVersion {
	s.InferenceSpec = v
	return s
}

func (s *ModelVersion) SetLabels(v []*ModelVersionLabels) *ModelVersion {
	s.Labels = v
	return s
}

func (s *ModelVersion) SetMetrics(v map[string]interface{}) *ModelVersion {
	s.Metrics = v
	return s
}

func (s *ModelVersion) SetOptions(v string) *ModelVersion {
	s.Options = &v
	return s
}

func (s *ModelVersion) SetOwnerId(v string) *ModelVersion {
	s.OwnerId = &v
	return s
}

func (s *ModelVersion) SetSourceId(v string) *ModelVersion {
	s.SourceId = &v
	return s
}

func (s *ModelVersion) SetSourceType(v string) *ModelVersion {
	s.SourceType = &v
	return s
}

func (s *ModelVersion) SetTrainingSpec(v map[string]interface{}) *ModelVersion {
	s.TrainingSpec = v
	return s
}

func (s *ModelVersion) SetUri(v string) *ModelVersion {
	s.Uri = &v
	return s
}

func (s *ModelVersion) SetUserId(v string) *ModelVersion {
	s.UserId = &v
	return s
}

func (s *ModelVersion) SetVersionDescription(v string) *ModelVersion {
	s.VersionDescription = &v
	return s
}

func (s *ModelVersion) SetVersionName(v string) *ModelVersion {
	s.VersionName = &v
	return s
}

func (s *ModelVersion) Validate() error {
	return dara.Validate(s)
}

type ModelVersionLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModelVersionLabels) String() string {
	return dara.Prettify(s)
}

func (s ModelVersionLabels) GoString() string {
	return s.String()
}

func (s *ModelVersionLabels) GetKey() *string {
	return s.Key
}

func (s *ModelVersionLabels) GetValue() *string {
	return s.Value
}

func (s *ModelVersionLabels) SetKey(v string) *ModelVersionLabels {
	s.Key = &v
	return s
}

func (s *ModelVersionLabels) SetValue(v string) *ModelVersionLabels {
	s.Value = &v
	return s
}

func (s *ModelVersionLabels) Validate() error {
	return dara.Validate(s)
}
