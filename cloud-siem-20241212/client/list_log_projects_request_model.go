// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListLogProjectsRequest
	GetLang() *string
	SetLogRegionId(v string) *ListLogProjectsRequest
	GetLogRegionId() *string
	SetLogUserId(v int64) *ListLogProjectsRequest
	GetLogUserId() *int64
	SetMaxResults(v int32) *ListLogProjectsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListLogProjectsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListLogProjectsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListLogProjectsRequest
	GetRoleFor() *int64
}

type ListLogProjectsRequest struct {
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou。
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// example:
	//
	// 173326*******。
	LogUserId *int64 `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 50。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListLogProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLogProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListLogProjectsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListLogProjectsRequest) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *ListLogProjectsRequest) GetLogUserId() *int64 {
	return s.LogUserId
}

func (s *ListLogProjectsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLogProjectsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLogProjectsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLogProjectsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListLogProjectsRequest) SetLang(v string) *ListLogProjectsRequest {
	s.Lang = &v
	return s
}

func (s *ListLogProjectsRequest) SetLogRegionId(v string) *ListLogProjectsRequest {
	s.LogRegionId = &v
	return s
}

func (s *ListLogProjectsRequest) SetLogUserId(v int64) *ListLogProjectsRequest {
	s.LogUserId = &v
	return s
}

func (s *ListLogProjectsRequest) SetMaxResults(v int32) *ListLogProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLogProjectsRequest) SetNextToken(v string) *ListLogProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *ListLogProjectsRequest) SetRegionId(v string) *ListLogProjectsRequest {
	s.RegionId = &v
	return s
}

func (s *ListLogProjectsRequest) SetRoleFor(v int64) *ListLogProjectsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListLogProjectsRequest) Validate() error {
	return dara.Validate(s)
}
