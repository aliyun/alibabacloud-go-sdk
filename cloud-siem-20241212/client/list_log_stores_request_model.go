// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogStoresRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListLogStoresRequest
	GetLang() *string
	SetLogProjectName(v string) *ListLogStoresRequest
	GetLogProjectName() *string
	SetLogRegionId(v string) *ListLogStoresRequest
	GetLogRegionId() *string
	SetLogUserId(v int64) *ListLogStoresRequest
	GetLogUserId() *int64
	SetMaxResults(v int32) *ListLogStoresRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListLogStoresRequest
	GetNextToken() *string
	SetRegionId(v string) *ListLogStoresRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListLogStoresRequest
	GetRoleFor() *int64
}

type ListLogStoresRequest struct {
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// aliyun-cloudsiem-data-173326*******-cn-hangzhou。
	LogProjectName *string `json:"LogProjectName,omitempty" xml:"LogProjectName,omitempty"`
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

func (s ListLogStoresRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLogStoresRequest) GoString() string {
	return s.String()
}

func (s *ListLogStoresRequest) GetLang() *string {
	return s.Lang
}

func (s *ListLogStoresRequest) GetLogProjectName() *string {
	return s.LogProjectName
}

func (s *ListLogStoresRequest) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *ListLogStoresRequest) GetLogUserId() *int64 {
	return s.LogUserId
}

func (s *ListLogStoresRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLogStoresRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLogStoresRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLogStoresRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListLogStoresRequest) SetLang(v string) *ListLogStoresRequest {
	s.Lang = &v
	return s
}

func (s *ListLogStoresRequest) SetLogProjectName(v string) *ListLogStoresRequest {
	s.LogProjectName = &v
	return s
}

func (s *ListLogStoresRequest) SetLogRegionId(v string) *ListLogStoresRequest {
	s.LogRegionId = &v
	return s
}

func (s *ListLogStoresRequest) SetLogUserId(v int64) *ListLogStoresRequest {
	s.LogUserId = &v
	return s
}

func (s *ListLogStoresRequest) SetMaxResults(v int32) *ListLogStoresRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLogStoresRequest) SetNextToken(v string) *ListLogStoresRequest {
	s.NextToken = &v
	return s
}

func (s *ListLogStoresRequest) SetRegionId(v string) *ListLogStoresRequest {
	s.RegionId = &v
	return s
}

func (s *ListLogStoresRequest) SetRoleFor(v int64) *ListLogStoresRequest {
	s.RoleFor = &v
	return s
}

func (s *ListLogStoresRequest) Validate() error {
	return dara.Validate(s)
}
