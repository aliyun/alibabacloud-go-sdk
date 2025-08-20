// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoBuildRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListRepoBuildRecordRequest
	GetInstanceId() *string
	SetPageNo(v int32) *ListRepoBuildRecordRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListRepoBuildRecordRequest
	GetPageSize() *int32
	SetRepoId(v string) *ListRepoBuildRecordRequest
	GetRepoId() *string
}

type ListRepoBuildRecordRequest struct {
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
	// The number of entries to return on each page.
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
	// crr-tquyps22md8****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s ListRepoBuildRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepoBuildRecordRequest) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRepoBuildRecordRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRepoBuildRecordRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRepoBuildRecordRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *ListRepoBuildRecordRequest) SetInstanceId(v string) *ListRepoBuildRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepoBuildRecordRequest) SetPageNo(v int32) *ListRepoBuildRecordRequest {
	s.PageNo = &v
	return s
}

func (s *ListRepoBuildRecordRequest) SetPageSize(v int32) *ListRepoBuildRecordRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepoBuildRecordRequest) SetRepoId(v string) *ListRepoBuildRecordRequest {
	s.RepoId = &v
	return s
}

func (s *ListRepoBuildRecordRequest) Validate() error {
	return dara.Validate(s)
}
