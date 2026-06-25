// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalStatus(v string) *UpdateModelVersionRequest
	GetApprovalStatus() *string
	SetCompressionSpec(v map[string]interface{}) *UpdateModelVersionRequest
	GetCompressionSpec() map[string]interface{}
	SetDistillationSpec(v map[string]interface{}) *UpdateModelVersionRequest
	GetDistillationSpec() map[string]interface{}
	SetEvaluationSpec(v map[string]interface{}) *UpdateModelVersionRequest
	GetEvaluationSpec() map[string]interface{}
	SetExtraInfo(v map[string]interface{}) *UpdateModelVersionRequest
	GetExtraInfo() map[string]interface{}
	SetInferenceSpec(v map[string]interface{}) *UpdateModelVersionRequest
	GetInferenceSpec() map[string]interface{}
	SetMetrics(v map[string]interface{}) *UpdateModelVersionRequest
	GetMetrics() map[string]interface{}
	SetOptions(v string) *UpdateModelVersionRequest
	GetOptions() *string
	SetSourceId(v string) *UpdateModelVersionRequest
	GetSourceId() *string
	SetSourceType(v string) *UpdateModelVersionRequest
	GetSourceType() *string
	SetTrainingSpec(v map[string]interface{}) *UpdateModelVersionRequest
	GetTrainingSpec() map[string]interface{}
	SetVersionDescription(v string) *UpdateModelVersionRequest
	GetVersionDescription() *string
}

type UpdateModelVersionRequest struct {
	// The approval status. Valid values:
	//
	// - Pending: The model is pending approval.
	//
	// - Approved: The model is approved to be published.
	//
	// - Rejected: The model is not approved to be published.
	//
	// example:
	//
	// Approved
	ApprovalStatus *string `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	// The compression configuration.
	//
	// example:
	//
	// {}
	CompressionSpec map[string]interface{} `json:"CompressionSpec,omitempty" xml:"CompressionSpec,omitempty"`
	// The distillation configuration.
	//
	// example:
	//
	// {}
	DistillationSpec map[string]interface{} `json:"DistillationSpec,omitempty" xml:"DistillationSpec,omitempty"`
	// The evaluation configuration.
	//
	// example:
	//
	// {}
	EvaluationSpec map[string]interface{} `json:"EvaluationSpec,omitempty" xml:"EvaluationSpec,omitempty"`
	// Other information.
	//
	// example:
	//
	// {
	//
	// 	"CoverUris": ["https://e***u.oss-cn-hangzhou.aliyuncs.com/st****017.preview.png"],
	//
	// 	"TrainedWords": ["albedo_overlord"]
	//
	// }
	ExtraInfo map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// Describes how to apply the model to downstream inference applications. For example, describe the processor and container for Elastic Algorithm Service (EAS). Example:
	//
	// `{ "processor": "tensorflow_gpu_1.12" }`.
	//
	// example:
	//
	// {
	//
	// 	"processor": "tensorflow_gpu_1.12"
	//
	// }
	InferenceSpec map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	// The model metrics.
	//
	// The length cannot exceed 8,192 characters after serialization.
	//
	// example:
	//
	// {
	//
	//   "Results": [{
	//
	//     "Dataset": {
	//
	//       "DatasetId": "d-sdkjanksaklerhfd"
	//
	//     },
	//
	//     "Metrics": {
	//
	//       "cer": 0.175
	//
	//     }
	//
	//   }, {
	//
	//     "Dataset": {
	//
	//       "Uri": "oss://xxxx/"
	//
	//     },
	//
	//     "Metrics": {
	//
	//       "cer": 0.172
	//
	//     }
	//
	//   }]
	//
	// }
	Metrics map[string]interface{} `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The extended field. This field is a JSON string.
	//
	// example:
	//
	// {}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The source ID.
	//
	// - If the source type is Custom, this field has no restrictions.
	//
	// - If the source is PAIFlow or TrainingService, the format is as follows:
	//
	// ```
	//
	// region=<region_id>,workspaceId=<workspace_id>,kind=<kind>,id=<id>
	//
	// ```
	//
	// The parameters are described as follows:
	//
	// - region: the Alibaba Cloud region ID.
	//
	// - workspaceId: the workspace ID.
	//
	// - kind: the type. Valid values: PipelineRun (PAI pipeline) or ServiceJob (training service).
	//
	// - id: the unique identifier.
	//
	// example:
	//
	// region=cn-shanghai,workspaceId=13**,kind=PipelineRun,id=run-sakdb****jdf
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type of the model. Valid values:
	//
	// - Custom (default): The model is a custom model.
	//
	// - PAIFlow: The model is from a PAI pipeline.
	//
	// - TrainingService: The model is from a PAI training service.
	//
	// example:
	//
	// PAIFlow
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The training configuration. This is used for fine-tuning and incremental training.
	//
	// example:
	//
	// {}
	TrainingSpec map[string]interface{} `json:"TrainingSpec,omitempty" xml:"TrainingSpec,omitempty"`
	// The description of the model version.
	//
	// example:
	//
	// General sentiment analysis.
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
}

func (s UpdateModelVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelVersionRequest) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *UpdateModelVersionRequest) GetCompressionSpec() map[string]interface{} {
	return s.CompressionSpec
}

func (s *UpdateModelVersionRequest) GetDistillationSpec() map[string]interface{} {
	return s.DistillationSpec
}

func (s *UpdateModelVersionRequest) GetEvaluationSpec() map[string]interface{} {
	return s.EvaluationSpec
}

func (s *UpdateModelVersionRequest) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *UpdateModelVersionRequest) GetInferenceSpec() map[string]interface{} {
	return s.InferenceSpec
}

func (s *UpdateModelVersionRequest) GetMetrics() map[string]interface{} {
	return s.Metrics
}

func (s *UpdateModelVersionRequest) GetOptions() *string {
	return s.Options
}

func (s *UpdateModelVersionRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateModelVersionRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateModelVersionRequest) GetTrainingSpec() map[string]interface{} {
	return s.TrainingSpec
}

func (s *UpdateModelVersionRequest) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *UpdateModelVersionRequest) SetApprovalStatus(v string) *UpdateModelVersionRequest {
	s.ApprovalStatus = &v
	return s
}

func (s *UpdateModelVersionRequest) SetCompressionSpec(v map[string]interface{}) *UpdateModelVersionRequest {
	s.CompressionSpec = v
	return s
}

func (s *UpdateModelVersionRequest) SetDistillationSpec(v map[string]interface{}) *UpdateModelVersionRequest {
	s.DistillationSpec = v
	return s
}

func (s *UpdateModelVersionRequest) SetEvaluationSpec(v map[string]interface{}) *UpdateModelVersionRequest {
	s.EvaluationSpec = v
	return s
}

func (s *UpdateModelVersionRequest) SetExtraInfo(v map[string]interface{}) *UpdateModelVersionRequest {
	s.ExtraInfo = v
	return s
}

func (s *UpdateModelVersionRequest) SetInferenceSpec(v map[string]interface{}) *UpdateModelVersionRequest {
	s.InferenceSpec = v
	return s
}

func (s *UpdateModelVersionRequest) SetMetrics(v map[string]interface{}) *UpdateModelVersionRequest {
	s.Metrics = v
	return s
}

func (s *UpdateModelVersionRequest) SetOptions(v string) *UpdateModelVersionRequest {
	s.Options = &v
	return s
}

func (s *UpdateModelVersionRequest) SetSourceId(v string) *UpdateModelVersionRequest {
	s.SourceId = &v
	return s
}

func (s *UpdateModelVersionRequest) SetSourceType(v string) *UpdateModelVersionRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateModelVersionRequest) SetTrainingSpec(v map[string]interface{}) *UpdateModelVersionRequest {
	s.TrainingSpec = v
	return s
}

func (s *UpdateModelVersionRequest) SetVersionDescription(v string) *UpdateModelVersionRequest {
	s.VersionDescription = &v
	return s
}

func (s *UpdateModelVersionRequest) Validate() error {
	return dara.Validate(s)
}
