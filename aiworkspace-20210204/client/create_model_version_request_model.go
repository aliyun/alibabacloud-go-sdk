// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalStatus(v string) *CreateModelVersionRequest
	GetApprovalStatus() *string
	SetCompressionSpec(v map[string]interface{}) *CreateModelVersionRequest
	GetCompressionSpec() map[string]interface{}
	SetDistillationSpec(v map[string]interface{}) *CreateModelVersionRequest
	GetDistillationSpec() map[string]interface{}
	SetEvaluationSpec(v map[string]interface{}) *CreateModelVersionRequest
	GetEvaluationSpec() map[string]interface{}
	SetExtraInfo(v map[string]interface{}) *CreateModelVersionRequest
	GetExtraInfo() map[string]interface{}
	SetFormatType(v string) *CreateModelVersionRequest
	GetFormatType() *string
	SetFrameworkType(v string) *CreateModelVersionRequest
	GetFrameworkType() *string
	SetInferenceSpec(v map[string]interface{}) *CreateModelVersionRequest
	GetInferenceSpec() map[string]interface{}
	SetLabels(v []*Label) *CreateModelVersionRequest
	GetLabels() []*Label
	SetMetrics(v map[string]interface{}) *CreateModelVersionRequest
	GetMetrics() map[string]interface{}
	SetOptions(v string) *CreateModelVersionRequest
	GetOptions() *string
	SetSourceId(v string) *CreateModelVersionRequest
	GetSourceId() *string
	SetSourceType(v string) *CreateModelVersionRequest
	GetSourceType() *string
	SetTrainingSpec(v map[string]interface{}) *CreateModelVersionRequest
	GetTrainingSpec() map[string]interface{}
	SetUri(v string) *CreateModelVersionRequest
	GetUri() *string
	SetVersionDescription(v string) *CreateModelVersionRequest
	GetVersionDescription() *string
	SetVersionName(v string) *CreateModelVersionRequest
	GetVersionName() *string
}

type CreateModelVersionRequest struct {
	// The approval status. Valid values:
	//
	// 	- Pending
	//
	// 	- Approved
	//
	// 	- Rejected
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
	CompressionSpec  map[string]interface{} `json:"CompressionSpec,omitempty" xml:"CompressionSpec,omitempty"`
	DistillationSpec map[string]interface{} `json:"DistillationSpec,omitempty" xml:"DistillationSpec,omitempty"`
	// The evaluation configuration.
	//
	// example:
	//
	// {}
	EvaluationSpec map[string]interface{} `json:"EvaluationSpec,omitempty" xml:"EvaluationSpec,omitempty"`
	// The additional information.
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
	// The model format. Valid values:
	//
	// 	- OfflineModel
	//
	// 	- SavedModel
	//
	// 	- Keras H5
	//
	// 	- Frozen Pb
	//
	// 	- Caffe Prototxt
	//
	// 	- TorchScript
	//
	// 	- XGBoost
	//
	// 	- PMML
	//
	// 	- AlinkModel
	//
	// 	- ONNX
	//
	// example:
	//
	// SavedModel
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// The model framework. Valid values:
	//
	// 	- Pytorch
	//
	// 	- XGBoost
	//
	// 	- Keras
	//
	// 	- Caffe
	//
	// 	- Alink
	//
	// 	- Xflow
	//
	// 	- TensorFlow
	//
	// example:
	//
	// TensorFlow
	FrameworkType *string `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	// Describes how to apply to downstream inference services. For example, describe the processor and container of EAS. Example: `{ "processor": "tensorflow_gpu_1.12" }`
	//
	// example:
	//
	// {
	//
	//     "processor": "tensorflow_gpu_1.12"
	//
	// }
	InferenceSpec map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	// The labels.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The metrics for the model. The length after serialization is limited to 8,192.
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
	// The extended field. This is a JSON string.
	//
	// example:
	//
	// {}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The ID of the model source.
	//
	// 	- If SourceType is set to Custom, this parameter is not limited.
	//
	// 	- If SourceType is set to PAIFlow or TrainingService, the ID of the model source is in the following format:
	//
	// <!---->
	//
	//     region=<region_id>,workspaceId=<workspace_id>,kind=<kind>,id=<id>
	//
	// Take note of the following parameters:
	//
	// 	- region indicates the region ID.
	//
	// 	- workspaceId indicates the workspace ID.
	//
	// 	- kind indicates the type. Valid values: PipelineRun (PAIFlow) and ServiceJob (training service).
	//
	// 	- id indicates the unique identifier.
	//
	// example:
	//
	// region=cn-shanghai,workspaceId=13**,kind=PipelineRun,id=run-sakdb****jdf
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The type of the model source. Valid values:
	//
	// 	- Custom (default)
	//
	// 	- PAIFlow
	//
	// 	- TrainingService: the Platform for AI (PAI) training service.
	//
	// example:
	//
	// PAIFlow
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The training configurations, which is used for fine-tuning and incremental training.
	//
	// example:
	//
	// {}
	TrainingSpec map[string]interface{} `json:"TrainingSpec,omitempty" xml:"TrainingSpec,omitempty"`
	// The URI of the model version, which is the location where the model is stored. Valid values:
	//
	// 	- The HTTP(S) address of the model. Example: `https://myweb.com/mymodel.tar.gz`.
	//
	// 	- The OSS path of the model, in the format of `oss://<bucket>.<endpoint>/object`. For information about endpoints, see [OSS regions and endpoints](https://help.aliyun.com/document_detail/31837.html). Example: `oss://mybucket.oss-cn-beijing.aliyuncs.com/mypath/`.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://mybucket.oss-cn-beijing.aliyuncs.com/mypath/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The version description.
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	// The model version, which is unique for each model. If you leave this parameter empty, the first version is **0.1.0*	- by default. After that, the minor version number is increased by 1 in sequence. For example, the second version number is **0.2.0**. A version number consists of a major version number, a minor version number, and a stage version number, separated by periods (.). The major version number and minor version number are numeric. The stage version number begins with a digit and can include numbers, underscores, and letters. For example, the version number is 1.1.0 or 2.3.4_beta.
	//
	// example:
	//
	// 0.1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateModelVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateModelVersionRequest) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *CreateModelVersionRequest) GetCompressionSpec() map[string]interface{} {
	return s.CompressionSpec
}

func (s *CreateModelVersionRequest) GetDistillationSpec() map[string]interface{} {
	return s.DistillationSpec
}

func (s *CreateModelVersionRequest) GetEvaluationSpec() map[string]interface{} {
	return s.EvaluationSpec
}

func (s *CreateModelVersionRequest) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *CreateModelVersionRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *CreateModelVersionRequest) GetFrameworkType() *string {
	return s.FrameworkType
}

func (s *CreateModelVersionRequest) GetInferenceSpec() map[string]interface{} {
	return s.InferenceSpec
}

func (s *CreateModelVersionRequest) GetLabels() []*Label {
	return s.Labels
}

func (s *CreateModelVersionRequest) GetMetrics() map[string]interface{} {
	return s.Metrics
}

func (s *CreateModelVersionRequest) GetOptions() *string {
	return s.Options
}

func (s *CreateModelVersionRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateModelVersionRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateModelVersionRequest) GetTrainingSpec() map[string]interface{} {
	return s.TrainingSpec
}

func (s *CreateModelVersionRequest) GetUri() *string {
	return s.Uri
}

func (s *CreateModelVersionRequest) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *CreateModelVersionRequest) GetVersionName() *string {
	return s.VersionName
}

func (s *CreateModelVersionRequest) SetApprovalStatus(v string) *CreateModelVersionRequest {
	s.ApprovalStatus = &v
	return s
}

func (s *CreateModelVersionRequest) SetCompressionSpec(v map[string]interface{}) *CreateModelVersionRequest {
	s.CompressionSpec = v
	return s
}

func (s *CreateModelVersionRequest) SetDistillationSpec(v map[string]interface{}) *CreateModelVersionRequest {
	s.DistillationSpec = v
	return s
}

func (s *CreateModelVersionRequest) SetEvaluationSpec(v map[string]interface{}) *CreateModelVersionRequest {
	s.EvaluationSpec = v
	return s
}

func (s *CreateModelVersionRequest) SetExtraInfo(v map[string]interface{}) *CreateModelVersionRequest {
	s.ExtraInfo = v
	return s
}

func (s *CreateModelVersionRequest) SetFormatType(v string) *CreateModelVersionRequest {
	s.FormatType = &v
	return s
}

func (s *CreateModelVersionRequest) SetFrameworkType(v string) *CreateModelVersionRequest {
	s.FrameworkType = &v
	return s
}

func (s *CreateModelVersionRequest) SetInferenceSpec(v map[string]interface{}) *CreateModelVersionRequest {
	s.InferenceSpec = v
	return s
}

func (s *CreateModelVersionRequest) SetLabels(v []*Label) *CreateModelVersionRequest {
	s.Labels = v
	return s
}

func (s *CreateModelVersionRequest) SetMetrics(v map[string]interface{}) *CreateModelVersionRequest {
	s.Metrics = v
	return s
}

func (s *CreateModelVersionRequest) SetOptions(v string) *CreateModelVersionRequest {
	s.Options = &v
	return s
}

func (s *CreateModelVersionRequest) SetSourceId(v string) *CreateModelVersionRequest {
	s.SourceId = &v
	return s
}

func (s *CreateModelVersionRequest) SetSourceType(v string) *CreateModelVersionRequest {
	s.SourceType = &v
	return s
}

func (s *CreateModelVersionRequest) SetTrainingSpec(v map[string]interface{}) *CreateModelVersionRequest {
	s.TrainingSpec = v
	return s
}

func (s *CreateModelVersionRequest) SetUri(v string) *CreateModelVersionRequest {
	s.Uri = &v
	return s
}

func (s *CreateModelVersionRequest) SetVersionDescription(v string) *CreateModelVersionRequest {
	s.VersionDescription = &v
	return s
}

func (s *CreateModelVersionRequest) SetVersionName(v string) *CreateModelVersionRequest {
	s.VersionName = &v
	return s
}

func (s *CreateModelVersionRequest) Validate() error {
	return dara.Validate(s)
}
