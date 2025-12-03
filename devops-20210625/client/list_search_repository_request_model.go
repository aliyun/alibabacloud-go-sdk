// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunPk(v string) *ListSearchRepositoryRequest
	GetAliyunPk() *string
	SetKeyword(v string) *ListSearchRepositoryRequest
	GetKeyword() *string
	SetOrder(v string) *ListSearchRepositoryRequest
	GetOrder() *string
	SetPage(v int32) *ListSearchRepositoryRequest
	GetPage() *int32
	SetPageSize(v int32) *ListSearchRepositoryRequest
	GetPageSize() *int32
	SetRepoPath(v *ListSearchRepositoryRequestRepoPath) *ListSearchRepositoryRequest
	GetRepoPath() *ListSearchRepositoryRequestRepoPath
	SetScope(v string) *ListSearchRepositoryRequest
	GetScope() *string
	SetSort(v string) *ListSearchRepositoryRequest
	GetSort() *string
	SetVisibilityLevel(v int32) *ListSearchRepositoryRequest
	GetVisibilityLevel() *int32
	SetOrganizationId(v string) *ListSearchRepositoryRequest
	GetOrganizationId() *string
}

type ListSearchRepositoryRequest struct {
	// example:
	//
	// 1840004904455497
	AliyunPk *string `json:"aliyunPk,omitempty" xml:"aliyunPk,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// default
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                               `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RepoPath *ListSearchRepositoryRequestRepoPath `json:"repoPath,omitempty" xml:"repoPath,omitempty" type:"Struct"`
	// example:
	//
	// all
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// example:
	//
	// desc
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
	// example:
	//
	// 0
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61e54b0e0bb300d827e1ae27
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s ListSearchRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSearchRepositoryRequest) GoString() string {
	return s.String()
}

func (s *ListSearchRepositoryRequest) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *ListSearchRepositoryRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListSearchRepositoryRequest) GetOrder() *string {
	return s.Order
}

func (s *ListSearchRepositoryRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListSearchRepositoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSearchRepositoryRequest) GetRepoPath() *ListSearchRepositoryRequestRepoPath {
	return s.RepoPath
}

func (s *ListSearchRepositoryRequest) GetScope() *string {
	return s.Scope
}

func (s *ListSearchRepositoryRequest) GetSort() *string {
	return s.Sort
}

func (s *ListSearchRepositoryRequest) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *ListSearchRepositoryRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListSearchRepositoryRequest) SetAliyunPk(v string) *ListSearchRepositoryRequest {
	s.AliyunPk = &v
	return s
}

func (s *ListSearchRepositoryRequest) SetKeyword(v string) *ListSearchRepositoryRequest {
	s.Keyword = &v
	return s
}

func (s *ListSearchRepositoryRequest) SetOrder(v string) *ListSearchRepositoryRequest {
	s.Order = &v
	return s
}

func (s *ListSearchRepositoryRequest) SetPage(v int32) *ListSearchRepositoryRequest {
	s.Page = &v
	return s
}

func (s *ListSearchRepositoryRequest) SetPageSize(v int32) *ListSearchRepositoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListSearchRepositoryRequest) SetRepoPath(v *ListSearchRepositoryRequestRepoPath) *ListSearchRepositoryRequest {
	s.RepoPath = v
	return s
}

func (s *ListSearchRepositoryRequest) SetScope(v string) *ListSearchRepositoryRequest {
	s.Scope = &v
	return s
}

func (s *ListSearchRepositoryRequest) SetSort(v string) *ListSearchRepositoryRequest {
	s.Sort = &v
	return s
}

func (s *ListSearchRepositoryRequest) SetVisibilityLevel(v int32) *ListSearchRepositoryRequest {
	s.VisibilityLevel = &v
	return s
}

func (s *ListSearchRepositoryRequest) SetOrganizationId(v string) *ListSearchRepositoryRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListSearchRepositoryRequest) Validate() error {
	if s.RepoPath != nil {
		if err := s.RepoPath.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSearchRepositoryRequestRepoPath struct {
	// example:
	//
	// term
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// example:
	//
	// equal
	OperatorType *string `json:"operatorType,omitempty" xml:"operatorType,omitempty"`
	// example:
	//
	// orgId/test-group/spring-boot-demo
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListSearchRepositoryRequestRepoPath) String() string {
	return dara.Prettify(s)
}

func (s ListSearchRepositoryRequestRepoPath) GoString() string {
	return s.String()
}

func (s *ListSearchRepositoryRequestRepoPath) GetMatchType() *string {
	return s.MatchType
}

func (s *ListSearchRepositoryRequestRepoPath) GetOperatorType() *string {
	return s.OperatorType
}

func (s *ListSearchRepositoryRequestRepoPath) GetValue() *string {
	return s.Value
}

func (s *ListSearchRepositoryRequestRepoPath) SetMatchType(v string) *ListSearchRepositoryRequestRepoPath {
	s.MatchType = &v
	return s
}

func (s *ListSearchRepositoryRequestRepoPath) SetOperatorType(v string) *ListSearchRepositoryRequestRepoPath {
	s.OperatorType = &v
	return s
}

func (s *ListSearchRepositoryRequestRepoPath) SetValue(v string) *ListSearchRepositoryRequestRepoPath {
	s.Value = &v
	return s
}

func (s *ListSearchRepositoryRequestRepoPath) Validate() error {
	return dara.Validate(s)
}
