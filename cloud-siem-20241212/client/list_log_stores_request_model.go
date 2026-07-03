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
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// aliyun-cloudsiem-data-173326*******-cn-hangzhou
	LogProjectName *string `json:"LogProjectName,omitempty" xml:"LogProjectName,omitempty"`
	// The ID of the log storage region.
	//
	// example:
	//
	// cn-hangzhou
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// The ID of the user who ingests the data.
	//
	// example:
	//
	// 173326*******
	LogUserId *int64 `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	// The maximum number of entries to return on this call.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start the next query. You do not need to specify this parameter for the first query. If a subsequent query is required, set the value to the NextToken value that is returned from the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region of the Data Management center. Select the region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member whose permissions are assumed by the administrator.
	//
	// example:
	//
	// 173326*******
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
