// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvType(v string) *ListDataSourcesShrinkRequest
	GetEnvType() *string
	SetName(v string) *ListDataSourcesShrinkRequest
	GetName() *string
	SetOrder(v string) *ListDataSourcesShrinkRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListDataSourcesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataSourcesShrinkRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataSourcesShrinkRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListDataSourcesShrinkRequest
	GetSortBy() *string
	SetTags(v string) *ListDataSourcesShrinkRequest
	GetTags() *string
	SetTypesShrink(v string) *ListDataSourcesShrinkRequest
	GetTypesShrink() *string
}

type ListDataSourcesShrinkRequest struct {
	// The environment in which the data sources are used. Valid values:
	//
	// 	- Dev: development environment
	//
	// 	- Prod: production environment
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The name of the data source. Fuzzy match by data source name is supported.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The order in which you want to sort the data sources. Valid values:
	//
	// 	- Desc: descending order
	//
	// 	- Asc: ascending order
	//
	// Default value: Desc
	//
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 17820
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The field that you want to use to sort the data sources. Valid values:
	//
	// 	- CreateTime
	//
	// 	- Id
	//
	// 	- Name
	//
	// Default value: CreateTime
	//
	// example:
	//
	// Id
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The tag of the data source. This parameter specifies a filter condition.
	//
	// 	- You can specify multiple tags, which are in the logical AND relation. For example, you can query the data sources that contain the following tags: `["tag1", "tag2", "tag3"]`.
	//
	// 	- If you do not configure this parameter, tag-based filtering is not performed. You can specify up to 10 tags.
	//
	// example:
	//
	// ["tag1", "tag2", "tag3"]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The data source types. This parameter specifies a filter condition. You can specify multiple data source types.
	TypesShrink *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s ListDataSourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesShrinkRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListDataSourcesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListDataSourcesShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDataSourcesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataSourcesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataSourcesShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataSourcesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDataSourcesShrinkRequest) GetTags() *string {
	return s.Tags
}

func (s *ListDataSourcesShrinkRequest) GetTypesShrink() *string {
	return s.TypesShrink
}

func (s *ListDataSourcesShrinkRequest) SetEnvType(v string) *ListDataSourcesShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetName(v string) *ListDataSourcesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetOrder(v string) *ListDataSourcesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetPageNumber(v int32) *ListDataSourcesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetPageSize(v int32) *ListDataSourcesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetProjectId(v int64) *ListDataSourcesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetSortBy(v string) *ListDataSourcesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetTags(v string) *ListDataSourcesShrinkRequest {
	s.Tags = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetTypesShrink(v string) *ListDataSourcesShrinkRequest {
	s.TypesShrink = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
