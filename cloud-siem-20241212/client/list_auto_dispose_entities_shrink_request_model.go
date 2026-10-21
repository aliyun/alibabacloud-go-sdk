// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoDisposeEntitiesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoDisposeRecordIdsShrink(v string) *ListAutoDisposeEntitiesShrinkRequest
	GetAutoDisposeRecordIdsShrink() *string
	SetCurrentPage(v string) *ListAutoDisposeEntitiesShrinkRequest
	GetCurrentPage() *string
	SetDataSourceType(v string) *ListAutoDisposeEntitiesShrinkRequest
	GetDataSourceType() *string
	SetLang(v string) *ListAutoDisposeEntitiesShrinkRequest
	GetLang() *string
	SetMaxResults(v int32) *ListAutoDisposeEntitiesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAutoDisposeEntitiesShrinkRequest
	GetNextToken() *string
	SetPageSize(v string) *ListAutoDisposeEntitiesShrinkRequest
	GetPageSize() *string
	SetUuid(v string) *ListAutoDisposeEntitiesShrinkRequest
	GetUuid() *string
}

type ListAutoDisposeEntitiesShrinkRequest struct {
	// The list of decision-making disposal IDs.
	AutoDisposeRecordIdsShrink *string `json:"AutoDisposeRecordIds,omitempty" xml:"AutoDisposeRecordIds,omitempty"`
	// The current page number. The value must be greater than or equal to 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data source type.
	//
	// This parameter is required.
	//
	// example:
	//
	// alibaba_cloud_sas
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// Note: This parameter is required. If you do not specify this parameter, the service returns a 400 error: "The specified Lang can not be empty." Explicitly pass a language code.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return in this request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that indicates the position from which the current call starts reading. An empty value indicates that all data has been read.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The UUID of the playbook. This is the unique identifier of the playbook.
	//
	// example:
	//
	// b2491e39-ddf2-478a-8c07-*****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListAutoDisposeEntitiesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAutoDisposeEntitiesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAutoDisposeEntitiesShrinkRequest) GetAutoDisposeRecordIdsShrink() *string {
	return s.AutoDisposeRecordIdsShrink
}

func (s *ListAutoDisposeEntitiesShrinkRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *ListAutoDisposeEntitiesShrinkRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListAutoDisposeEntitiesShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListAutoDisposeEntitiesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAutoDisposeEntitiesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAutoDisposeEntitiesShrinkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAutoDisposeEntitiesShrinkRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListAutoDisposeEntitiesShrinkRequest) SetAutoDisposeRecordIdsShrink(v string) *ListAutoDisposeEntitiesShrinkRequest {
	s.AutoDisposeRecordIdsShrink = &v
	return s
}

func (s *ListAutoDisposeEntitiesShrinkRequest) SetCurrentPage(v string) *ListAutoDisposeEntitiesShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAutoDisposeEntitiesShrinkRequest) SetDataSourceType(v string) *ListAutoDisposeEntitiesShrinkRequest {
	s.DataSourceType = &v
	return s
}

func (s *ListAutoDisposeEntitiesShrinkRequest) SetLang(v string) *ListAutoDisposeEntitiesShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListAutoDisposeEntitiesShrinkRequest) SetMaxResults(v int32) *ListAutoDisposeEntitiesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAutoDisposeEntitiesShrinkRequest) SetNextToken(v string) *ListAutoDisposeEntitiesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListAutoDisposeEntitiesShrinkRequest) SetPageSize(v string) *ListAutoDisposeEntitiesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAutoDisposeEntitiesShrinkRequest) SetUuid(v string) *ListAutoDisposeEntitiesShrinkRequest {
	s.Uuid = &v
	return s
}

func (s *ListAutoDisposeEntitiesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
