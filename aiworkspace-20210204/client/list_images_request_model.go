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
	// The visibility of the image. Currently only takes effect for custom images.
	//
	// - PUBLIC: Public.
	//
	// - PRIVATE: Private.
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The image URI.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/pai-compression/nlp:gpu
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	// The label filter conditions. Separate multiple conditions with commas (,).
	//
	// The format of a single filter condition is `Key=Value`.
	//
	// Valid values of Key:
	//
	// - system.chipType
	//
	// - system.dsw.cudaVersion
	//
	// - system.dsw.fromImageId
	//
	// - system.dsw.fromInstanceId
	//
	// - system.dsw.id
	//
	// - system.dsw.os
	//
	// - system.dsw.osVersion
	//
	// - system.dsw.resourceType
	//
	// - system.dsw.rootImageId
	//
	// - system.dsw.stage
	//
	// - system.dsw.tag
	//
	// - system.dsw.type
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
	// The image name. Fuzzy match is supported.
	//
	// example:
	//
	// tensorflow_2.9
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The options. Currently supports ExactMatchName for exact match search on Name.
	//
	// example:
	//
	// ExactMatchName
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The sort order for the specified sort field in a paged query. Used together with SortBy. Default value: ASC.
	//
	// - ASC: Ascending order.
	//
	// - DESC: Descending order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number of the image list. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Fuzzy searches by image name and description.
	//
	// example:
	//
	// name
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The sort field for a paged query. Currently, the GmtCreateTime field is used for sorting.
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
	// The workspace ID. For information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
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
