// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiAppWarningByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAiAppWarningByPageRequest
	GetAppId() *string
	SetCurrentPage(v int32) *ListAiAppWarningByPageRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListAiAppWarningByPageRequest
	GetPageSize() *int32
	SetQuery(v string) *ListAiAppWarningByPageRequest
	GetQuery() *string
	SetRegionId(v string) *ListAiAppWarningByPageRequest
	GetRegionId() *string
}

type ListAiAppWarningByPageRequest struct {
	// The application ID.
	//
	// example:
	//
	// id-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query condition. This parameter is required and cannot be empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// {}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAiAppWarningByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppWarningByPageRequest) GoString() string {
	return s.String()
}

func (s *ListAiAppWarningByPageRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAiAppWarningByPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAiAppWarningByPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAiAppWarningByPageRequest) GetQuery() *string {
	return s.Query
}

func (s *ListAiAppWarningByPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAiAppWarningByPageRequest) SetAppId(v string) *ListAiAppWarningByPageRequest {
	s.AppId = &v
	return s
}

func (s *ListAiAppWarningByPageRequest) SetCurrentPage(v int32) *ListAiAppWarningByPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAiAppWarningByPageRequest) SetPageSize(v int32) *ListAiAppWarningByPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListAiAppWarningByPageRequest) SetQuery(v string) *ListAiAppWarningByPageRequest {
	s.Query = &v
	return s
}

func (s *ListAiAppWarningByPageRequest) SetRegionId(v string) *ListAiAppWarningByPageRequest {
	s.RegionId = &v
	return s
}

func (s *ListAiAppWarningByPageRequest) Validate() error {
	return dara.Validate(s)
}
