// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobOutputModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOutputModels(v []*ListTrainingJobOutputModelsResponseBodyOutputModels) *ListTrainingJobOutputModelsResponseBody
	GetOutputModels() []*ListTrainingJobOutputModelsResponseBodyOutputModels
}

type ListTrainingJobOutputModelsResponseBody struct {
	OutputModels []*ListTrainingJobOutputModelsResponseBodyOutputModels `json:"OutputModels,omitempty" xml:"OutputModels,omitempty" type:"Repeated"`
}

func (s ListTrainingJobOutputModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobOutputModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrainingJobOutputModelsResponseBody) GetOutputModels() []*ListTrainingJobOutputModelsResponseBodyOutputModels {
	return s.OutputModels
}

func (s *ListTrainingJobOutputModelsResponseBody) SetOutputModels(v []*ListTrainingJobOutputModelsResponseBodyOutputModels) *ListTrainingJobOutputModelsResponseBody {
	s.OutputModels = v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTrainingJobOutputModelsResponseBodyOutputModels struct {
	CompressionSpec map[string]interface{} `json:"CompressionSpec,omitempty" xml:"CompressionSpec,omitempty"`
	// example:
	//
	// {}
	EvaluationSpec map[string]interface{} `json:"EvaluationSpec,omitempty" xml:"EvaluationSpec,omitempty"`
	// example:
	//
	// {}
	InferenceSpec map[string]interface{}                                       `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	Labels        []*ListTrainingJobOutputModelsResponseBodyOutputModelsLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// {
	//
	//       "lr": 0.000001,
	//
	//       "train_loss": 2.6345
	//
	// }
	Metrics map[string]interface{} `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// example:
	//
	// model
	OutputChannelName *string `json:"OutputChannelName,omitempty" xml:"OutputChannelName,omitempty"`
	// example:
	//
	// region=cn-shanghai,workspaceId=1345,kind=PipelineRun,id=run-sakdbaskjdf
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// PAIFlow
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// {}
	TrainingSpec map[string]interface{} `json:"TrainingSpec,omitempty" xml:"TrainingSpec,omitempty"`
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou.aliyuncs.com/path/to/output/channel/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s ListTrainingJobOutputModelsResponseBodyOutputModels) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobOutputModelsResponseBodyOutputModels) GoString() string {
	return s.String()
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) GetCompressionSpec() map[string]interface{} {
	return s.CompressionSpec
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) GetEvaluationSpec() map[string]interface{} {
	return s.EvaluationSpec
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) GetInferenceSpec() map[string]interface{} {
	return s.InferenceSpec
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) GetLabels() []*ListTrainingJobOutputModelsResponseBodyOutputModelsLabels {
	return s.Labels
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) GetMetrics() map[string]interface{} {
	return s.Metrics
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) GetOutputChannelName() *string {
	return s.OutputChannelName
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) GetSourceId() *string {
	return s.SourceId
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) GetSourceType() *string {
	return s.SourceType
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) GetTrainingSpec() map[string]interface{} {
	return s.TrainingSpec
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) GetUri() *string {
	return s.Uri
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetCompressionSpec(v map[string]interface{}) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.CompressionSpec = v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetEvaluationSpec(v map[string]interface{}) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.EvaluationSpec = v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetInferenceSpec(v map[string]interface{}) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.InferenceSpec = v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetLabels(v []*ListTrainingJobOutputModelsResponseBodyOutputModelsLabels) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.Labels = v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetMetrics(v map[string]interface{}) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.Metrics = v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetOutputChannelName(v string) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.OutputChannelName = &v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetSourceId(v string) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.SourceId = &v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetSourceType(v string) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.SourceType = &v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetTrainingSpec(v map[string]interface{}) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.TrainingSpec = v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) SetUri(v string) *ListTrainingJobOutputModelsResponseBodyOutputModels {
	s.Uri = &v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModels) Validate() error {
	return dara.Validate(s)
}

type ListTrainingJobOutputModelsResponseBodyOutputModelsLabels struct {
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// StableDiffusion
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTrainingJobOutputModelsResponseBodyOutputModelsLabels) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobOutputModelsResponseBodyOutputModelsLabels) GoString() string {
	return s.String()
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModelsLabels) GetKey() *string {
	return s.Key
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModelsLabels) GetValue() *string {
	return s.Value
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModelsLabels) SetKey(v string) *ListTrainingJobOutputModelsResponseBodyOutputModelsLabels {
	s.Key = &v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModelsLabels) SetValue(v string) *ListTrainingJobOutputModelsResponseBodyOutputModelsLabels {
	s.Value = &v
	return s
}

func (s *ListTrainingJobOutputModelsResponseBodyOutputModelsLabels) Validate() error {
	return dara.Validate(s)
}
