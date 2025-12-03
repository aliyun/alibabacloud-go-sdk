// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchSourceCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilePath(v *ListSearchSourceCodeRequestFilePath) *ListSearchSourceCodeRequest
	GetFilePath() *ListSearchSourceCodeRequestFilePath
	SetIsCodeBlock(v bool) *ListSearchSourceCodeRequest
	GetIsCodeBlock() *bool
	SetKeyword(v string) *ListSearchSourceCodeRequest
	GetKeyword() *string
	SetLanguage(v string) *ListSearchSourceCodeRequest
	GetLanguage() *string
	SetOrder(v string) *ListSearchSourceCodeRequest
	GetOrder() *string
	SetPage(v int32) *ListSearchSourceCodeRequest
	GetPage() *int32
	SetPageSize(v int32) *ListSearchSourceCodeRequest
	GetPageSize() *int32
	SetRepoPath(v *ListSearchSourceCodeRequestRepoPath) *ListSearchSourceCodeRequest
	GetRepoPath() *ListSearchSourceCodeRequestRepoPath
	SetScope(v string) *ListSearchSourceCodeRequest
	GetScope() *string
	SetSort(v string) *ListSearchSourceCodeRequest
	GetSort() *string
	SetOrganizationId(v string) *ListSearchSourceCodeRequest
	GetOrganizationId() *string
}

type ListSearchSourceCodeRequest struct {
	FilePath *ListSearchSourceCodeRequestFilePath `json:"filePath,omitempty" xml:"filePath,omitempty" type:"Struct"`
	// example:
	//
	// false
	IsCodeBlock *bool `json:"isCodeBlock,omitempty" xml:"isCodeBlock,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// Java
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
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
	// 20
	PageSize *int32                               `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RepoPath *ListSearchSourceCodeRequestRepoPath `json:"repoPath,omitempty" xml:"repoPath,omitempty" type:"Struct"`
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
	// 66c0c9fffeb86b450c199fcd
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s ListSearchSourceCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSearchSourceCodeRequest) GoString() string {
	return s.String()
}

func (s *ListSearchSourceCodeRequest) GetFilePath() *ListSearchSourceCodeRequestFilePath {
	return s.FilePath
}

func (s *ListSearchSourceCodeRequest) GetIsCodeBlock() *bool {
	return s.IsCodeBlock
}

func (s *ListSearchSourceCodeRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListSearchSourceCodeRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListSearchSourceCodeRequest) GetOrder() *string {
	return s.Order
}

func (s *ListSearchSourceCodeRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListSearchSourceCodeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSearchSourceCodeRequest) GetRepoPath() *ListSearchSourceCodeRequestRepoPath {
	return s.RepoPath
}

func (s *ListSearchSourceCodeRequest) GetScope() *string {
	return s.Scope
}

func (s *ListSearchSourceCodeRequest) GetSort() *string {
	return s.Sort
}

func (s *ListSearchSourceCodeRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListSearchSourceCodeRequest) SetFilePath(v *ListSearchSourceCodeRequestFilePath) *ListSearchSourceCodeRequest {
	s.FilePath = v
	return s
}

func (s *ListSearchSourceCodeRequest) SetIsCodeBlock(v bool) *ListSearchSourceCodeRequest {
	s.IsCodeBlock = &v
	return s
}

func (s *ListSearchSourceCodeRequest) SetKeyword(v string) *ListSearchSourceCodeRequest {
	s.Keyword = &v
	return s
}

func (s *ListSearchSourceCodeRequest) SetLanguage(v string) *ListSearchSourceCodeRequest {
	s.Language = &v
	return s
}

func (s *ListSearchSourceCodeRequest) SetOrder(v string) *ListSearchSourceCodeRequest {
	s.Order = &v
	return s
}

func (s *ListSearchSourceCodeRequest) SetPage(v int32) *ListSearchSourceCodeRequest {
	s.Page = &v
	return s
}

func (s *ListSearchSourceCodeRequest) SetPageSize(v int32) *ListSearchSourceCodeRequest {
	s.PageSize = &v
	return s
}

func (s *ListSearchSourceCodeRequest) SetRepoPath(v *ListSearchSourceCodeRequestRepoPath) *ListSearchSourceCodeRequest {
	s.RepoPath = v
	return s
}

func (s *ListSearchSourceCodeRequest) SetScope(v string) *ListSearchSourceCodeRequest {
	s.Scope = &v
	return s
}

func (s *ListSearchSourceCodeRequest) SetSort(v string) *ListSearchSourceCodeRequest {
	s.Sort = &v
	return s
}

func (s *ListSearchSourceCodeRequest) SetOrganizationId(v string) *ListSearchSourceCodeRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListSearchSourceCodeRequest) Validate() error {
	if s.FilePath != nil {
		if err := s.FilePath.Validate(); err != nil {
			return err
		}
	}
	if s.RepoPath != nil {
		if err := s.RepoPath.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSearchSourceCodeRequestFilePath struct {
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
	// orgId/test-group/spring-boot-demo/test.java
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListSearchSourceCodeRequestFilePath) String() string {
	return dara.Prettify(s)
}

func (s ListSearchSourceCodeRequestFilePath) GoString() string {
	return s.String()
}

func (s *ListSearchSourceCodeRequestFilePath) GetMatchType() *string {
	return s.MatchType
}

func (s *ListSearchSourceCodeRequestFilePath) GetOperatorType() *string {
	return s.OperatorType
}

func (s *ListSearchSourceCodeRequestFilePath) GetValue() *string {
	return s.Value
}

func (s *ListSearchSourceCodeRequestFilePath) SetMatchType(v string) *ListSearchSourceCodeRequestFilePath {
	s.MatchType = &v
	return s
}

func (s *ListSearchSourceCodeRequestFilePath) SetOperatorType(v string) *ListSearchSourceCodeRequestFilePath {
	s.OperatorType = &v
	return s
}

func (s *ListSearchSourceCodeRequestFilePath) SetValue(v string) *ListSearchSourceCodeRequestFilePath {
	s.Value = &v
	return s
}

func (s *ListSearchSourceCodeRequestFilePath) Validate() error {
	return dara.Validate(s)
}

type ListSearchSourceCodeRequestRepoPath struct {
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
	// xxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListSearchSourceCodeRequestRepoPath) String() string {
	return dara.Prettify(s)
}

func (s ListSearchSourceCodeRequestRepoPath) GoString() string {
	return s.String()
}

func (s *ListSearchSourceCodeRequestRepoPath) GetMatchType() *string {
	return s.MatchType
}

func (s *ListSearchSourceCodeRequestRepoPath) GetOperatorType() *string {
	return s.OperatorType
}

func (s *ListSearchSourceCodeRequestRepoPath) GetValue() *string {
	return s.Value
}

func (s *ListSearchSourceCodeRequestRepoPath) SetMatchType(v string) *ListSearchSourceCodeRequestRepoPath {
	s.MatchType = &v
	return s
}

func (s *ListSearchSourceCodeRequestRepoPath) SetOperatorType(v string) *ListSearchSourceCodeRequestRepoPath {
	s.OperatorType = &v
	return s
}

func (s *ListSearchSourceCodeRequestRepoPath) SetValue(v string) *ListSearchSourceCodeRequestRepoPath {
	s.Value = &v
	return s
}

func (s *ListSearchSourceCodeRequestRepoPath) Validate() error {
	return dara.Validate(s)
}
