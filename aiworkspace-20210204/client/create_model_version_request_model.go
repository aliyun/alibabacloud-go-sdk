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
	// - Pending: The version is pending approval.
	//
	// - Approved: The version is approved for deployment.
	//
	// - Rejected: The version is rejected for deployment.
	//
	// example:
	//
	// Approved
	ApprovalStatus *string `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	// The compression configurations.
	//
	// example:
	//
	// {}
	CompressionSpec map[string]interface{} `json:"CompressionSpec,omitempty" xml:"CompressionSpec,omitempty"`
	// The distillation configurations.
	//
	// example:
	//
	// {}
	DistillationSpec map[string]interface{} `json:"DistillationSpec,omitempty" xml:"DistillationSpec,omitempty"`
	// The evaluation configurations.
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
	// The format of the model. Valid values:
	//
	// - OfflineModel
	//
	// - SavedModel
	//
	// - Keras H5
	//
	// - Frozen Pb
	//
	// - Caffe Prototxt
	//
	// - TorchScript
	//
	// - XGBoost
	//
	// - PMML
	//
	// - AlinkModel
	//
	// - ONNX
	//
	// example:
	//
	// SavedModel
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// The framework of the model. Valid values:
	//
	// - Pytorch
	//
	// - XGBoost
	//
	// - Keras
	//
	// - Caffe
	//
	// - Alink
	//
	// - Xflow
	//
	// - TensorFlow
	//
	// example:
	//
	// TensorFlow
	FrameworkType *string `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	// The configurations for downstream inference services, such as the processor and container for Elastic Algorithm Service (EAS). Example:
	//
	// `{ "processor": "tensorflow_gpu_1.12" }`
	//
	// example:
	//
	// {
	//
	//     "processor": "tensorflow_gpu_1.12"
	//
	// }
	InferenceSpec map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	// The list of labels.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The model metrics.
	//
	// The serialized data cannot exceed 8,192 bytes in length.
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
	// The extended fields. This parameter is a JSON string.
	//
	// example:
	//
	// {}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The source ID.
	//
	// - If SourceType is set to Custom, this parameter has no format restrictions.
	//
	// - If SourceType is PAIFlow or TrainingService, the value must be in the following format:
	//
	// ```
	//
	// region=<region_id>,workspaceId=<workspace_id>,kind=<kind>,id=<id>
	//
	// ```
	//
	// The fields are described as follows:
	//
	// - region: The ID of the Alibaba Cloud region.
	//
	// - workspaceId: The ID of the workspace.
	//
	// - kind: The type. Valid values: PipelineRun (PAI pipeline) and ServiceJob (training service).
	//
	// - id: The unique identifier.
	//
	// example:
	//
	// region=cn-shanghai,workspaceId=13**,kind=PipelineRun,id=run-sakdb****jdf
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type of the model. Valid values:
	//
	// - Custom (default): The model is custom.
	//
	// - PAIFlow: The model is from a PAI pipeline.
	//
	// - TrainingService: The model is from a PAI training service.
	//
	// example:
	//
	// PAIFlow
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The training configurations. These configurations are used for fine-tuning and incremental training.
	//
	// example:
	//
	// {}
	TrainingSpec map[string]interface{} `json:"TrainingSpec,omitempty" xml:"TrainingSpec,omitempty"`
	// The URI of the model version, which is the storage location of the model. The following types of model URIs are supported:
	//
	// - An HTTP or HTTPS URL of the model. Example: `https://myweb.com/mymodel.tar.gz`.
	//
	// - If the model is stored in Object Storage Service (OSS), the URI must be in the `oss://<bucket>.<endpoint>/object` format. For more information about endpoints, see [Endpoints](https://help.aliyun.com/document_detail/31837.html). Example: `oss://mybucket.oss-cn-beijing.aliyuncs.com/mypath/`.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://mybucket.oss-cn-beijing.aliyuncs.com/mypath/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The description of the model version.
	//
	// example:
	//
	// Sentiment analysis.
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	// The model version. The version must be unique within the model. If you do not specify this parameter, the first version defaults to **0.1.0**. The minor version number is then incremented by 1 for each subsequent version. For example, the second version defaults to **0.2.0**.
	//
	// A version number consists of a major version, a minor version, and a patch version, separated by periods (.). The major and minor versions are numbers. The patch version can start with a number and contain underscores (_) and letters. Examples: 1.1.0 and 2.3.4_beta.
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
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
