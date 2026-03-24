// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterCheckResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckKey(v string) *ListClusterCheckResultRequest
	GetCheckKey() *string
	SetClusterId(v string) *ListClusterCheckResultRequest
	GetClusterId() *string
	SetCurrentPage(v int32) *ListClusterCheckResultRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListClusterCheckResultRequest
	GetLang() *string
	SetPageSize(v int32) *ListClusterCheckResultRequest
	GetPageSize() *int32
	SetRiskLevels(v []*string) *ListClusterCheckResultRequest
	GetRiskLevels() []*string
	SetSortType(v string) *ListClusterCheckResultRequest
	GetSortType() *string
	SetStatuses(v []*string) *ListClusterCheckResultRequest
	GetStatuses() []*string
}

type ListClusterCheckResultRequest struct {
	// Fuzzy search key for check items.
	//
	// example:
	//
	// container
	CheckKey *string `json:"CheckKey,omitempty" xml:"CheckKey,omitempty"`
	// Cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c3aaf6c8085f84791882eef200cd2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Page number for the current page in a paginated query. The default value is **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Language type for requests and responses. The default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Number of records to display per page when performing a paginated query. The default value is **20**, indicating 20 records per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of risk levels.
	RiskLevels []*string `json:"RiskLevels,omitempty" xml:"RiskLevels,omitempty" type:"Repeated"`
	// Custom sorting type. Values:
	//
	// - **RISK_LEVEL**: Risk level.
	//
	// - **STATUS**: Check item status.
	//
	// example:
	//
	// STATUS
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// List of check item statuses.
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s ListClusterCheckResultRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterCheckResultRequest) GoString() string {
	return s.String()
}

func (s *ListClusterCheckResultRequest) GetCheckKey() *string {
	return s.CheckKey
}

func (s *ListClusterCheckResultRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterCheckResultRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListClusterCheckResultRequest) GetLang() *string {
	return s.Lang
}

func (s *ListClusterCheckResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClusterCheckResultRequest) GetRiskLevels() []*string {
	return s.RiskLevels
}

func (s *ListClusterCheckResultRequest) GetSortType() *string {
	return s.SortType
}

func (s *ListClusterCheckResultRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListClusterCheckResultRequest) SetCheckKey(v string) *ListClusterCheckResultRequest {
	s.CheckKey = &v
	return s
}

func (s *ListClusterCheckResultRequest) SetClusterId(v string) *ListClusterCheckResultRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterCheckResultRequest) SetCurrentPage(v int32) *ListClusterCheckResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListClusterCheckResultRequest) SetLang(v string) *ListClusterCheckResultRequest {
	s.Lang = &v
	return s
}

func (s *ListClusterCheckResultRequest) SetPageSize(v int32) *ListClusterCheckResultRequest {
	s.PageSize = &v
	return s
}

func (s *ListClusterCheckResultRequest) SetRiskLevels(v []*string) *ListClusterCheckResultRequest {
	s.RiskLevels = v
	return s
}

func (s *ListClusterCheckResultRequest) SetSortType(v string) *ListClusterCheckResultRequest {
	s.SortType = &v
	return s
}

func (s *ListClusterCheckResultRequest) SetStatuses(v []*string) *ListClusterCheckResultRequest {
	s.Statuses = v
	return s
}

func (s *ListClusterCheckResultRequest) Validate() error {
	return dara.Validate(s)
}
