// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalStatus(v string) *ListModelVersionsRequest
	GetApprovalStatus() *string
	SetFormatType(v string) *ListModelVersionsRequest
	GetFormatType() *string
	SetFrameworkType(v string) *ListModelVersionsRequest
	GetFrameworkType() *string
	SetLabel(v string) *ListModelVersionsRequest
	GetLabel() *string
	SetOrder(v string) *ListModelVersionsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListModelVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelVersionsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListModelVersionsRequest
	GetSortBy() *string
	SetSourceId(v string) *ListModelVersionsRequest
	GetSourceId() *string
	SetSourceType(v string) *ListModelVersionsRequest
	GetSourceType() *string
	SetVersionName(v string) *ListModelVersionsRequest
	GetVersionName() *string
}

type ListModelVersionsRequest struct {
	// The approval status. This parameter is used to filter the model version list. Valid values:
	//
	// - Pending: The model version is pending approval.
	//
	// - Approved: The model version is approved for publishing.
	//
	// - Rejected: The model version is rejected for publishing.
	//
	// example:
	//
	// Approved
	ApprovalStatus *string `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	// The model format. This parameter is used to filter the model version list. Valid values:
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
	// The model framework. This parameter is used to filter the model version list. Valid values:
	//
	// - Pytorch
	//
	//   -XGBoost
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
	// The label string. This parameter is used to filter the list. Model versions that have the specified string in the key or value of their labels are returned.
	//
	// example:
	//
	// key1
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The order in which to sort the entries in the paged query. The default value is ASC.
	//
	// - ASC: ascending order.
	//
	// - DESC: descending order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number of the model version list. The value starts from 1. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for a paged query. The default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field to use for sorting in the paged query. Currently, the GmtCreateTime field is used for sorting.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The source ID.
	//
	// - If the source type is Custom, this parameter is not restricted.
	//
	// - If the source is PAIFlow or TrainingService, the format is as follows:
	//
	// ```
	//
	// region=<region_id>,workspaceId=<workspace_id>,kind=<kind>,id=<id>
	//
	// ```
	//
	// where:
	//
	// - region is the Alibaba Cloud region ID.
	//
	// - workspaceId is the workspace ID.
	//
	// - kind: the type. Valid values: PipelineRun (PAIFlow pipeline) and ServiceJob (training service).
	//
	// - id: the unique identifier.
	//
	// example:
	//
	// region=cn-shanghai,workspaceId=13**,kind=PipelineRun,id=run-sakdb****jdf
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type of the model. This parameter is used to filter the model version list. Valid values:
	//
	// - Custom (default): a custom model.
	//
	// - PAIFlow: a model from a PAI pipeline.
	//
	// - TrainingService: a model from a PAI training service.
	//
	// example:
	//
	// PAIFlow
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The model version name. This parameter is used to filter the model version list.
	//
	// example:
	//
	// 1.0.1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListModelVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListModelVersionsRequest) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *ListModelVersionsRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *ListModelVersionsRequest) GetFrameworkType() *string {
	return s.FrameworkType
}

func (s *ListModelVersionsRequest) GetLabel() *string {
	return s.Label
}

func (s *ListModelVersionsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListModelVersionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelVersionsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListModelVersionsRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ListModelVersionsRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListModelVersionsRequest) GetVersionName() *string {
	return s.VersionName
}

func (s *ListModelVersionsRequest) SetApprovalStatus(v string) *ListModelVersionsRequest {
	s.ApprovalStatus = &v
	return s
}

func (s *ListModelVersionsRequest) SetFormatType(v string) *ListModelVersionsRequest {
	s.FormatType = &v
	return s
}

func (s *ListModelVersionsRequest) SetFrameworkType(v string) *ListModelVersionsRequest {
	s.FrameworkType = &v
	return s
}

func (s *ListModelVersionsRequest) SetLabel(v string) *ListModelVersionsRequest {
	s.Label = &v
	return s
}

func (s *ListModelVersionsRequest) SetOrder(v string) *ListModelVersionsRequest {
	s.Order = &v
	return s
}

func (s *ListModelVersionsRequest) SetPageNumber(v int32) *ListModelVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelVersionsRequest) SetPageSize(v int32) *ListModelVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelVersionsRequest) SetSortBy(v string) *ListModelVersionsRequest {
	s.SortBy = &v
	return s
}

func (s *ListModelVersionsRequest) SetSourceId(v string) *ListModelVersionsRequest {
	s.SourceId = &v
	return s
}

func (s *ListModelVersionsRequest) SetSourceType(v string) *ListModelVersionsRequest {
	s.SourceType = &v
	return s
}

func (s *ListModelVersionsRequest) SetVersionName(v string) *ListModelVersionsRequest {
	s.VersionName = &v
	return s
}

func (s *ListModelVersionsRequest) Validate() error {
	return dara.Validate(s)
}
