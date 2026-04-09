// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiMcpServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *ListApiMcpServersRequest
	GetCreateTime() *string
	SetDescription(v string) *ListApiMcpServersRequest
	GetDescription() *string
	SetId(v string) *ListApiMcpServersRequest
	GetId() *string
	SetKeyword(v string) *ListApiMcpServersRequest
	GetKeyword() *string
	SetLanguage(v string) *ListApiMcpServersRequest
	GetLanguage() *string
	SetMaxResults(v int32) *ListApiMcpServersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListApiMcpServersRequest
	GetNextToken() *string
	SetSkip(v int32) *ListApiMcpServersRequest
	GetSkip() *int32
	SetSourceType(v string) *ListApiMcpServersRequest
	GetSourceType() *string
	SetUpdateTime(v string) *ListApiMcpServersRequest
	GetUpdateTime() *string
}

type ListApiMcpServersRequest struct {
	// example:
	//
	// 2024-10-30T02:10:13Z
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// v6ZZ7ftCzEILW***
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// oss
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// ZH_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// AAAAAZjtYxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 5
	Skip *int32 `json:"skip,omitempty" xml:"skip,omitempty"`
	// example:
	//
	// system
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// example:
	//
	// 2024-06-05T02:27:39Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListApiMcpServersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiMcpServersRequest) GoString() string {
	return s.String()
}

func (s *ListApiMcpServersRequest) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListApiMcpServersRequest) GetDescription() *string {
	return s.Description
}

func (s *ListApiMcpServersRequest) GetId() *string {
	return s.Id
}

func (s *ListApiMcpServersRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListApiMcpServersRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListApiMcpServersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApiMcpServersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApiMcpServersRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListApiMcpServersRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListApiMcpServersRequest) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListApiMcpServersRequest) SetCreateTime(v string) *ListApiMcpServersRequest {
	s.CreateTime = &v
	return s
}

func (s *ListApiMcpServersRequest) SetDescription(v string) *ListApiMcpServersRequest {
	s.Description = &v
	return s
}

func (s *ListApiMcpServersRequest) SetId(v string) *ListApiMcpServersRequest {
	s.Id = &v
	return s
}

func (s *ListApiMcpServersRequest) SetKeyword(v string) *ListApiMcpServersRequest {
	s.Keyword = &v
	return s
}

func (s *ListApiMcpServersRequest) SetLanguage(v string) *ListApiMcpServersRequest {
	s.Language = &v
	return s
}

func (s *ListApiMcpServersRequest) SetMaxResults(v int32) *ListApiMcpServersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApiMcpServersRequest) SetNextToken(v string) *ListApiMcpServersRequest {
	s.NextToken = &v
	return s
}

func (s *ListApiMcpServersRequest) SetSkip(v int32) *ListApiMcpServersRequest {
	s.Skip = &v
	return s
}

func (s *ListApiMcpServersRequest) SetSourceType(v string) *ListApiMcpServersRequest {
	s.SourceType = &v
	return s
}

func (s *ListApiMcpServersRequest) SetUpdateTime(v string) *ListApiMcpServersRequest {
	s.UpdateTime = &v
	return s
}

func (s *ListApiMcpServersRequest) Validate() error {
	return dara.Validate(s)
}
