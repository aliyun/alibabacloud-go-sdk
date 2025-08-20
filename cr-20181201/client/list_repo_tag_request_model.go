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
	SetPageNo(v int32) *ListRepoTagRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListRepoTagRequest
	GetPageSize() *int32
	SetRepoId(v string) *ListRepoTagRequest
	GetRepoId() *string
}

type ListRepoTagRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the repository.
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
