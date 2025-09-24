// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalStatus(v string) *GetModelVersionResponseBody
	GetApprovalStatus() *string
	SetCompressionSpec(v map[string]interface{}) *GetModelVersionResponseBody
	GetCompressionSpec() map[string]interface{}
	SetDistillationSpec(v map[string]interface{}) *GetModelVersionResponseBody
	GetDistillationSpec() map[string]interface{}
	SetEvaluationSpec(v map[string]interface{}) *GetModelVersionResponseBody
	GetEvaluationSpec() map[string]interface{}
	SetExtraInfo(v map[string]interface{}) *GetModelVersionResponseBody
	GetExtraInfo() map[string]interface{}
	SetFormatType(v string) *GetModelVersionResponseBody
	GetFormatType() *string
	SetFrameworkType(v string) *GetModelVersionResponseBody
	GetFrameworkType() *string
	SetGmtCreateTime(v string) *GetModelVersionResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetModelVersionResponseBody
	GetGmtModifiedTime() *string
	SetInferenceSpec(v map[string]interface{}) *GetModelVersionResponseBody
	GetInferenceSpec() map[string]interface{}
	SetLabels(v []*Label) *GetModelVersionResponseBody
	GetLabels() []*Label
	SetMetrics(v map[string]interface{}) *GetModelVersionResponseBody
	GetMetrics() map[string]interface{}
	SetOptions(v string) *GetModelVersionResponseBody
	GetOptions() *string
	SetOwnerId(v string) *GetModelVersionResponseBody
	GetOwnerId() *string
	SetRequestId(v string) *GetModelVersionResponseBody
	GetRequestId() *string
	SetSourceId(v string) *GetModelVersionResponseBody
	GetSourceId() *string
	SetSourceType(v string) *GetModelVersionResponseBody
	GetSourceType() *string
	SetTrainingSpec(v map[string]interface{}) *GetModelVersionResponseBody
	GetTrainingSpec() map[string]interface{}
	SetUri(v string) *GetModelVersionResponseBody
	GetUri() *string
	SetUserId(v string) *GetModelVersionResponseBody
	GetUserId() *string
	SetVersionDescription(v string) *GetModelVersionResponseBody
	GetVersionDescription() *string
	SetVersionName(v string) *GetModelVersionResponseBody
	GetVersionName() *string
}

type GetModelVersionResponseBody struct {
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
	// 	- Pytorch -XGBoost
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
	// 2021-01-30T12:51:33.028Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the model was last modified, in UTC. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Describes how to apply to downstream inference services. For example, describes the processor and container of Elastic Algorithm Service (EAS).
	//
	// example:
	//
	// {
	//
	//     "Processor": "tensorflow_gpu_1.12"
	//
	// }
	InferenceSpec map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	// The labels.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The metrics.
	//
	// example:
	//
	// {}
	Metrics map[string]interface{} `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The extended field. The value of this parameter is a JSON string.
	//
	// example:
	//
	// {}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1234567890******
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// Take note of the following parameters:
	//
	// 	- region is the region ID.
	//
	// 	- workspaceId is the ID of the workspace.
	//
	// 	- kind is the type. Valid values: PipelineRun (PAIFlow) and ServiceJob (training service).
	//
	// 	- id is a unique identifier.
	//
	// example:
	//
	// region=cn-shanghai,workspaceId=13**,kind=PipelineRun,id=run-sakdb****jdf
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type of the model. Valid values:
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
	// The training configurations used for fine-tuning and incremental training.
	//
	// example:
	//
	// {}
	TrainingSpec map[string]interface{} `json:"TrainingSpec,omitempty" xml:"TrainingSpec,omitempty"`
	// The URI of the model version, which is the location where the model is stored. Valid values:
	//
	// 	- The HTTP(S) address of the model. Example: `https://myweb.com/mymodel.tar.gz`.
	//
	// 	- The Object Storage Service (OSS) path of the model, in the format of `oss://<bucket>.<endpoint>/object`. For endpoint, see [OSS regions and endpoints](https://help.aliyun.com/document_detail/31837.html). Example: `oss://mybucket.oss-cn-beijing.aliyuncs.com/mypath/`.
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1234567890******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The version description.
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	// The model version.
	//
	// example:
	//
	// 0.1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetModelVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelVersionResponseBody) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *GetModelVersionResponseBody) GetCompressionSpec() map[string]interface{} {
	return s.CompressionSpec
}

func (s *GetModelVersionResponseBody) GetDistillationSpec() map[string]interface{} {
	return s.DistillationSpec
}

func (s *GetModelVersionResponseBody) GetEvaluationSpec() map[string]interface{} {
	return s.EvaluationSpec
}

func (s *GetModelVersionResponseBody) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *GetModelVersionResponseBody) GetFormatType() *string {
	return s.FormatType
}

func (s *GetModelVersionResponseBody) GetFrameworkType() *string {
	return s.FrameworkType
}

func (s *GetModelVersionResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetModelVersionResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetModelVersionResponseBody) GetInferenceSpec() map[string]interface{} {
	return s.InferenceSpec
}

func (s *GetModelVersionResponseBody) GetLabels() []*Label {
	return s.Labels
}

func (s *GetModelVersionResponseBody) GetMetrics() map[string]interface{} {
	return s.Metrics
}

func (s *GetModelVersionResponseBody) GetOptions() *string {
	return s.Options
}

func (s *GetModelVersionResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetModelVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelVersionResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *GetModelVersionResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *GetModelVersionResponseBody) GetTrainingSpec() map[string]interface{} {
	return s.TrainingSpec
}

func (s *GetModelVersionResponseBody) GetUri() *string {
	return s.Uri
}

func (s *GetModelVersionResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetModelVersionResponseBody) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *GetModelVersionResponseBody) GetVersionName() *string {
	return s.VersionName
}

func (s *GetModelVersionResponseBody) SetApprovalStatus(v string) *GetModelVersionResponseBody {
	s.ApprovalStatus = &v
	return s
}

func (s *GetModelVersionResponseBody) SetCompressionSpec(v map[string]interface{}) *GetModelVersionResponseBody {
	s.CompressionSpec = v
	return s
}

func (s *GetModelVersionResponseBody) SetDistillationSpec(v map[string]interface{}) *GetModelVersionResponseBody {
	s.DistillationSpec = v
	return s
}

func (s *GetModelVersionResponseBody) SetEvaluationSpec(v map[string]interface{}) *GetModelVersionResponseBody {
	s.EvaluationSpec = v
	return s
}

func (s *GetModelVersionResponseBody) SetExtraInfo(v map[string]interface{}) *GetModelVersionResponseBody {
	s.ExtraInfo = v
	return s
}

func (s *GetModelVersionResponseBody) SetFormatType(v string) *GetModelVersionResponseBody {
	s.FormatType = &v
	return s
}

func (s *GetModelVersionResponseBody) SetFrameworkType(v string) *GetModelVersionResponseBody {
	s.FrameworkType = &v
	return s
}

func (s *GetModelVersionResponseBody) SetGmtCreateTime(v string) *GetModelVersionResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetModelVersionResponseBody) SetGmtModifiedTime(v string) *GetModelVersionResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetModelVersionResponseBody) SetInferenceSpec(v map[string]interface{}) *GetModelVersionResponseBody {
	s.InferenceSpec = v
	return s
}

func (s *GetModelVersionResponseBody) SetLabels(v []*Label) *GetModelVersionResponseBody {
	s.Labels = v
	return s
}

func (s *GetModelVersionResponseBody) SetMetrics(v map[string]interface{}) *GetModelVersionResponseBody {
	s.Metrics = v
	return s
}

func (s *GetModelVersionResponseBody) SetOptions(v string) *GetModelVersionResponseBody {
	s.Options = &v
	return s
}

func (s *GetModelVersionResponseBody) SetOwnerId(v string) *GetModelVersionResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetModelVersionResponseBody) SetRequestId(v string) *GetModelVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelVersionResponseBody) SetSourceId(v string) *GetModelVersionResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetModelVersionResponseBody) SetSourceType(v string) *GetModelVersionResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetModelVersionResponseBody) SetTrainingSpec(v map[string]interface{}) *GetModelVersionResponseBody {
	s.TrainingSpec = v
	return s
}

func (s *GetModelVersionResponseBody) SetUri(v string) *GetModelVersionResponseBody {
	s.Uri = &v
	return s
}

func (s *GetModelVersionResponseBody) SetUserId(v string) *GetModelVersionResponseBody {
	s.UserId = &v
	return s
}

func (s *GetModelVersionResponseBody) SetVersionDescription(v string) *GetModelVersionResponseBody {
	s.VersionDescription = &v
	return s
}

func (s *GetModelVersionResponseBody) SetVersionName(v string) *GetModelVersionResponseBody {
	s.VersionName = &v
	return s
}

func (s *GetModelVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
