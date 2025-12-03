// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchCommitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListSearchCommitRequest
	GetKeyword() *string
	SetOrder(v string) *ListSearchCommitRequest
	GetOrder() *string
	SetPage(v int32) *ListSearchCommitRequest
	GetPage() *int32
	SetPageSize(v int32) *ListSearchCommitRequest
	GetPageSize() *int32
	SetRepoPath(v *ListSearchCommitRequestRepoPath) *ListSearchCommitRequest
	GetRepoPath() *ListSearchCommitRequestRepoPath
	SetScope(v string) *ListSearchCommitRequest
	GetScope() *string
	SetSort(v string) *ListSearchCommitRequest
	GetSort() *string
	SetOrganizationId(v string) *ListSearchCommitRequest
	GetOrganizationId() *string
}

type ListSearchCommitRequest struct {
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
	PageSize *int32                           `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RepoPath *ListSearchCommitRequestRepoPath `json:"repoPath,omitempty" xml:"repoPath,omitempty" type:"Struct"`
	// example:
	//
	// all
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// example:
	//
	// desc
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60d54f3daccf2bbd6659f3ad
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s ListSearchCommitRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSearchCommitRequest) GoString() string {
	return s.String()
}

func (s *ListSearchCommitRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListSearchCommitRequest) GetOrder() *string {
	return s.Order
}

func (s *ListSearchCommitRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListSearchCommitRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSearchCommitRequest) GetRepoPath() *ListSearchCommitRequestRepoPath {
	return s.RepoPath
}

func (s *ListSearchCommitRequest) GetScope() *string {
	return s.Scope
}

func (s *ListSearchCommitRequest) GetSort() *string {
	return s.Sort
}

func (s *ListSearchCommitRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListSearchCommitRequest) SetKeyword(v string) *ListSearchCommitRequest {
	s.Keyword = &v
	return s
}

func (s *ListSearchCommitRequest) SetOrder(v string) *ListSearchCommitRequest {
	s.Order = &v
	return s
}

func (s *ListSearchCommitRequest) SetPage(v int32) *ListSearchCommitRequest {
	s.Page = &v
	return s
}

func (s *ListSearchCommitRequest) SetPageSize(v int32) *ListSearchCommitRequest {
	s.PageSize = &v
	return s
}

func (s *ListSearchCommitRequest) SetRepoPath(v *ListSearchCommitRequestRepoPath) *ListSearchCommitRequest {
	s.RepoPath = v
	return s
}

func (s *ListSearchCommitRequest) SetScope(v string) *ListSearchCommitRequest {
	s.Scope = &v
	return s
}

func (s *ListSearchCommitRequest) SetSort(v string) *ListSearchCommitRequest {
	s.Sort = &v
	return s
}

func (s *ListSearchCommitRequest) SetOrganizationId(v string) *ListSearchCommitRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListSearchCommitRequest) Validate() error {
	if s.RepoPath != nil {
		if err := s.RepoPath.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSearchCommitRequestRepoPath struct {
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

func (s ListSearchCommitRequestRepoPath) String() string {
	return dara.Prettify(s)
}

func (s ListSearchCommitRequestRepoPath) GoString() string {
	return s.String()
}

func (s *ListSearchCommitRequestRepoPath) GetMatchType() *string {
	return s.MatchType
}

func (s *ListSearchCommitRequestRepoPath) GetOperatorType() *string {
	return s.OperatorType
}

func (s *ListSearchCommitRequestRepoPath) GetValue() *string {
	return s.Value
}

func (s *ListSearchCommitRequestRepoPath) SetMatchType(v string) *ListSearchCommitRequestRepoPath {
	s.MatchType = &v
	return s
}

func (s *ListSearchCommitRequestRepoPath) SetOperatorType(v string) *ListSearchCommitRequestRepoPath {
	s.OperatorType = &v
	return s
}

func (s *ListSearchCommitRequestRepoPath) SetValue(v string) *ListSearchCommitRequestRepoPath {
	s.Value = &v
	return s
}

func (s *ListSearchCommitRequestRepoPath) Validate() error {
	return dara.Validate(s)
}
