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
	// The approval status based on which the model versions are queried. Valid values:
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
	// The model format used to filter model versions. Valid values:
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
	// The framework used to filter model versions.
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
	// The label. Model versions whose label key or label value contains a specific label are filtered.
	//
	// example:
	//
	// key1
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The order in which the entries are sorted by the specific field on the returned page. Default value: ASC.
	//
	// 	- ASC
	//
	// 	- DESC
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field used to sort the results. The GmtCreateTime field is used for sorting.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
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
	// The source type used to filter model versions. Valid values:
	//
	// 	- Custom (default)
	//
	// 	- PAIFlow
	//
	// 	- TrainingService
	//
	// example:
	//
	// PAIFlow
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The model version used to filter model versions.
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
