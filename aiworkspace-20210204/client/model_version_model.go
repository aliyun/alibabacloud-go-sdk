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
	// The model format.
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
	// The model framework.
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
	// The time when the model was created, in UTC. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the model was last modified, in UTC. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The inference configurations applied to the downstream, such as the configuration of the processor or container of Elastic Algorithm Service (EAS). Example: `{ "processor": "tensorflow_gpu_1.12" }`
	//
	// example:
	//
	// {
	//
	// 	"processor": "tensorflow_gpu_1.12"
	//
	// }
	InferenceSpec map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	// The labels.
	Labels []*ModelVersionLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The model metrics.
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
	// The extended field. The value is a JSON string.
	//
	// example:
	//
	// {}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 155770209******
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The source ID.
	//
	// 	- If the source type is Custom, this field is not limited.
	//
	// 	- If the source type is PAIFlow or TrainingService, the format is:
	//
	// <!---->
	//
	//     region=<region_id>,workspaceId=<workspace_id>,kind=<kind>,id=<id>
	//
	// region is the ID of the Alibaba Cloud region. workspacceId is the ID of the workspace. kind is the type. Valid values: PipelineRun (PAIFlow pipeline) and ServiceJob (training service). id is the unique identifier.
	//
	// example:
	//
	// region=cn-shanghai,workspaceId=13**,kind=PipelineRun,id=run-sakdb****jdf
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type of the model. Default value: Custom.
	//
	// 	- Custom
	//
	// 	- PAIFlow
	//
	// 	- TrainingService
	//
	// example:
	//
	// PAIFlow
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The training configuration, used for fine-tuning and incremental training.
	//
	// example:
	//
	// {}
	TrainingSpec map[string]interface{} `json:"TrainingSpec,omitempty" xml:"TrainingSpec,omitempty"`
	// The URI of the model version, which is the location where the model is stored. The value can be the HTTP(S) address of the model, such as `https://myweb.com/mymodel.tar.gz`. If the model is stored in an Object Storage Service (OSS) bucket, the value is in the `oss://<bucket>.<endpoint>/object` format. The endpoint can be queried in the OSS console, such as `oss://mybucket.oss-cn-beijing.aliyuncs.com/mypath/`.
	//
	// example:
	//
	// oss://mybucket.oss-cn-beijing.aliyuncs.com/mypath/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 155770209******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The model version description.
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	// The model version, which is unique for the model. If you leave this parameter empty, `0.1.0` is the first version by default. Then, the minor version number incremented by 1 is used as the second version `0.2.0`.
	//
	// The version consists of a major version number, a minor version number, and a patch version number. The version numbers are separated with periods (`.`). The major and minor version numbers are digits, and the patch version number starts with a digit followed by an underscore (`_`) and a letter. such as 1.1.0 or 2.3.4_beta.
	//
	// Regular expression: `"^\\\\d+\\\\.\\\\d+\\\\.\\\\d+(_\\\\w+)?$"`
	//
	// example:
	//
	// 0.1.0
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
