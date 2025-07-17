// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAssetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDataAssetsResponseBodyPagingInfo) *ListDataAssetsResponseBody
	GetPagingInfo() *ListDataAssetsResponseBodyPagingInfo
	SetRequestId(v string) *ListDataAssetsResponseBody
	GetRequestId() *string
}

type ListDataAssetsResponseBody struct {
	// The pagination information.
	PagingInfo *ListDataAssetsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataAssetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataAssetsResponseBody) GetPagingInfo() *ListDataAssetsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDataAssetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataAssetsResponseBody) SetPagingInfo(v *ListDataAssetsResponseBodyPagingInfo) *ListDataAssetsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDataAssetsResponseBody) SetRequestId(v string) *ListDataAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataAssetsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataAssetsResponseBodyPagingInfo struct {
	// The data assets.
	DataAssets []*ListDataAssetsResponseBodyPagingInfoDataAssets `json:"DataAssets,omitempty" xml:"DataAssets,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataAssetsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDataAssetsResponseBodyPagingInfo) GetDataAssets() []*ListDataAssetsResponseBodyPagingInfoDataAssets {
	return s.DataAssets
}

func (s *ListDataAssetsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataAssetsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataAssetsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataAssetsResponseBodyPagingInfo) SetDataAssets(v []*ListDataAssetsResponseBodyPagingInfoDataAssets) *ListDataAssetsResponseBodyPagingInfo {
	s.DataAssets = v
	return s
}

func (s *ListDataAssetsResponseBodyPagingInfo) SetPageNumber(v int32) *ListDataAssetsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDataAssetsResponseBodyPagingInfo) SetPageSize(v int32) *ListDataAssetsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDataAssetsResponseBodyPagingInfo) SetTotalCount(v int32) *ListDataAssetsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDataAssetsResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListDataAssetsResponseBodyPagingInfoDataAssets struct {
	// The mappings between data assets and tags.
	DataAssetTagMappings []*ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings `json:"DataAssetTagMappings,omitempty" xml:"DataAssetTagMappings,omitempty" type:"Repeated"`
	// The environment of the workspace to which the data asset belongs. Valid values:
	//
	// 	- Dev: development environment
	//
	// 	- Prod: production environment
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The data asset ID.
	//
	// example:
	//
	// 7259557313
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the data asset.
	//
	// example:
	//
	// ali_cn_es_gfn
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 54275
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the data asset. Valid values:
	//
	// 	- ACS::DataWorks::Table
	//
	// 	- ACS::DataWorks::Task
	//
	// example:
	//
	// ACS::DataWorks::Task
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataAssetsResponseBodyPagingInfoDataAssets) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetsResponseBodyPagingInfoDataAssets) GoString() string {
	return s.String()
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssets) GetDataAssetTagMappings() []*ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings {
	return s.DataAssetTagMappings
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssets) GetEnvType() *string {
	return s.EnvType
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssets) GetId() *string {
	return s.Id
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssets) GetName() *string {
	return s.Name
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssets) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssets) GetType() *string {
	return s.Type
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssets) SetDataAssetTagMappings(v []*ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings) *ListDataAssetsResponseBodyPagingInfoDataAssets {
	s.DataAssetTagMappings = v
	return s
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssets) SetEnvType(v string) *ListDataAssetsResponseBodyPagingInfoDataAssets {
	s.EnvType = &v
	return s
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssets) SetId(v string) *ListDataAssetsResponseBodyPagingInfoDataAssets {
	s.Id = &v
	return s
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssets) SetName(v string) *ListDataAssetsResponseBodyPagingInfoDataAssets {
	s.Name = &v
	return s
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssets) SetProjectId(v int64) *ListDataAssetsResponseBodyPagingInfoDataAssets {
	s.ProjectId = &v
	return s
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssets) SetType(v string) *ListDataAssetsResponseBodyPagingInfoDataAssets {
	s.Type = &v
	return s
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssets) Validate() error {
	return dara.Validate(s)
}

type ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings struct {
	// Indicates whether the lineage-based automatic backtrack feature is enabled for the mapping.
	//
	// example:
	//
	// false
	AutoTraceEnabled *bool `json:"AutoTraceEnabled,omitempty" xml:"AutoTraceEnabled,omitempty"`
	// The creator of the mapping between the data asset and the tag.
	//
	// example:
	//
	// 12345
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The data asset ID.
	//
	// example:
	//
	// 7259557313
	DataAssetId *string `json:"DataAssetId,omitempty" xml:"DataAssetId,omitempty"`
	// The tag key.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The way in which the mapping between the data asset and the tag is created. Valid values:
	//
	// 	- System
	//
	// 	- UserDefined
	//
	// example:
	//
	// UserDefined
	TagSource *string `json:"TagSource,omitempty" xml:"TagSource,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings) GoString() string {
	return s.String()
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings) GetAutoTraceEnabled() *bool {
	return s.AutoTraceEnabled
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings) GetCreator() *string {
	return s.Creator
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings) GetDataAssetId() *string {
	return s.DataAssetId
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings) GetKey() *string {
	return s.Key
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings) GetTagSource() *string {
	return s.TagSource
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings) GetValue() *string {
	return s.Value
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings) SetAutoTraceEnabled(v bool) *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings {
	s.AutoTraceEnabled = &v
	return s
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings) SetCreator(v string) *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings {
	s.Creator = &v
	return s
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings) SetDataAssetId(v string) *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings {
	s.DataAssetId = &v
	return s
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings) SetKey(v string) *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings {
	s.Key = &v
	return s
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings) SetTagSource(v string) *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings {
	s.TagSource = &v
	return s
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings) SetValue(v string) *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings {
	s.Value = &v
	return s
}

func (s *ListDataAssetsResponseBodyPagingInfoDataAssetsDataAssetTagMappings) Validate() error {
	return dara.Validate(s)
}
