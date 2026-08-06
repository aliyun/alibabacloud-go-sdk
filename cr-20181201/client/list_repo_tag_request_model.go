// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListRepoTagRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListRepoTagRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRepoTagRequest
	GetNextToken() *string
	SetPageNo(v int32) *ListRepoTagRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListRepoTagRequest
	GetPageSize() *int32
	SetRepoId(v string) *ListRepoTagRequest
	GetRepoId() *string
}

type ListRepoTagRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. The maximum value is 100.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The repository ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-tquyps22md8p****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s ListRepoTagRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepoTagRequest) GoString() string {
	return s.String()
}

func (s *ListRepoTagRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRepoTagRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRepoTagRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRepoTagRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRepoTagRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRepoTagRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *ListRepoTagRequest) SetInstanceId(v string) *ListRepoTagRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepoTagRequest) SetMaxResults(v int32) *ListRepoTagRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRepoTagRequest) SetNextToken(v string) *ListRepoTagRequest {
	s.NextToken = &v
	return s
}

func (s *ListRepoTagRequest) SetPageNo(v int32) *ListRepoTagRequest {
	s.PageNo = &v
	return s
}

func (s *ListRepoTagRequest) SetPageSize(v int32) *ListRepoTagRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepoTagRequest) SetRepoId(v string) *ListRepoTagRequest {
	s.RepoId = &v
	return s
}

func (s *ListRepoTagRequest) Validate() error {
	return dara.Validate(s)
}
