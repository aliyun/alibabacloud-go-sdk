// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *ListImagesRequest
	GetAccessibility() *string
	SetImageUri(v string) *ListImagesRequest
	GetImageUri() *string
	SetLabels(v string) *ListImagesRequest
	GetLabels() *string
	SetName(v string) *ListImagesRequest
	GetName() *string
	SetOptions(v string) *ListImagesRequest
	GetOptions() *string
	SetOrder(v string) *ListImagesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListImagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListImagesRequest
	GetPageSize() *int32
	SetQuery(v string) *ListImagesRequest
	GetQuery() *string
	SetSortBy(v string) *ListImagesRequest
	GetSortBy() *string
	SetVerbose(v bool) *ListImagesRequest
	GetVerbose() *bool
	SetWorkspaceId(v string) *ListImagesRequest
	GetWorkspaceId() *string
}

type ListImagesRequest struct {
	// The visibility of the image. This parameter is valid only for custom images.
	//
	// - PUBLIC: The image is public.
	//
	// - PRIVATE: The image is private.
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	ImageUri      *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	// The filter conditions for labels. Separate multiple conditions with commas (,).
	//
	// The format for a single filter condition is `Key=Value`.
	//
	// The supported values for Key are:
	//
	// - system.chipType
	//
	// - system.dsw\\.cudaVersion
	//
	// - system.dsw\\.fromImageId
	//
	// - system.dsw\\.fromInstanceId
	//
	// - system.dsw\\.id
	//
	// - system.dsw\\.os
	//
	// - system.dsw\\.osVersion
	//
	// - system.dsw\\.resourceType
	//
	// - system.dsw\\.rootImageId
	//
	// - system.dsw\\.stage
	//
	// - system.dsw\\.tag
	//
	// - system.dsw\\.type
	//
	// - system.framework
	//
	// - system.origin
	//
	// - system.pythonVersion
	//
	// - system.source
	//
	// - system.supported.dlc
	//
	// - system.supported.dsw
	//
	// example:
	//
	// system.framework=XGBoost 1.6.0,system.official=true
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The name of the image. Fuzzy search is supported.
	//
	// example:
	//
	// tensorflow_2.9
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ExactMatchName
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The order in which to sort the results of a paged query. This parameter is used with SortBy. The default value is ASC.
	//
	// - ASC: ascending order.
	//
	// - DESC: descending order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number of the image list. The value starts from 1. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for a paged query. The default value is 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Performs a fuzzy search by image name and description.
	//
	// example:
	//
	// name
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The field to use for sorting in a paged query. Currently, only the GmtCreateTime field is used for sorting.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// Specifies whether to display non-essential information. Non-essential information currently includes Labels. Valid values:
	//
	// - true: Includes non-essential information.
	//
	// - false: Does not include non-essential information.
	//
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// The workspace ID. For more information about how to obtain a workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// example:
	//
	// 20******55
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListImagesRequest) GetImageUri() *string {
	return s.ImageUri
}

func (s *ListImagesRequest) GetLabels() *string {
	return s.Labels
}

func (s *ListImagesRequest) GetName() *string {
	return s.Name
}

func (s *ListImagesRequest) GetOptions() *string {
	return s.Options
}

func (s *ListImagesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListImagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListImagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImagesRequest) GetQuery() *string {
	return s.Query
}

func (s *ListImagesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListImagesRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *ListImagesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListImagesRequest) SetAccessibility(v string) *ListImagesRequest {
	s.Accessibility = &v
	return s
}

func (s *ListImagesRequest) SetImageUri(v string) *ListImagesRequest {
	s.ImageUri = &v
	return s
}

func (s *ListImagesRequest) SetLabels(v string) *ListImagesRequest {
	s.Labels = &v
	return s
}

func (s *ListImagesRequest) SetName(v string) *ListImagesRequest {
	s.Name = &v
	return s
}

func (s *ListImagesRequest) SetOptions(v string) *ListImagesRequest {
	s.Options = &v
	return s
}

func (s *ListImagesRequest) SetOrder(v string) *ListImagesRequest {
	s.Order = &v
	return s
}

func (s *ListImagesRequest) SetPageNumber(v int32) *ListImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListImagesRequest) SetPageSize(v int32) *ListImagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListImagesRequest) SetQuery(v string) *ListImagesRequest {
	s.Query = &v
	return s
}

func (s *ListImagesRequest) SetSortBy(v string) *ListImagesRequest {
	s.SortBy = &v
	return s
}

func (s *ListImagesRequest) SetVerbose(v bool) *ListImagesRequest {
	s.Verbose = &v
	return s
}

func (s *ListImagesRequest) SetWorkspaceId(v string) *ListImagesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListImagesRequest) Validate() error {
	return dara.Validate(s)
}
