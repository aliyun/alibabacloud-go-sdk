// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoSyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListRepoSyncTaskRequest
	GetInstanceId() *string
	SetPageNo(v int32) *ListRepoSyncTaskRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListRepoSyncTaskRequest
	GetPageSize() *int32
	SetRepoName(v string) *ListRepoSyncTaskRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *ListRepoSyncTaskRequest
	GetRepoNamespaceName() *string
	SetSyncRecordId(v string) *ListRepoSyncTaskRequest
	GetSyncRecordId() *string
	SetTag(v string) *ListRepoSyncTaskRequest
	GetTag() *string
}

type ListRepoSyncTaskRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The repository name.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// ns
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The ID of the synchronization task record, which is the same as SyncBatchTaskId in the response.
	//
	// >  If an image meets multiple synchronization rules and multiple synchronization tasks are generated for the image, these synchronization tasks use the same SyncBatchTaskId.
	//
	// example:
	//
	// crsr-7lph66uloi6h****
	SyncRecordId *string `json:"SyncRecordId,omitempty" xml:"SyncRecordId,omitempty"`
	// The image tag.
	//
	// example:
	//
	// nginx
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListRepoSyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepoSyncTaskRequest) GoString() string {
	return s.String()
}

func (s *ListRepoSyncTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRepoSyncTaskRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRepoSyncTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRepoSyncTaskRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *ListRepoSyncTaskRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *ListRepoSyncTaskRequest) GetSyncRecordId() *string {
	return s.SyncRecordId
}

func (s *ListRepoSyncTaskRequest) GetTag() *string {
	return s.Tag
}

func (s *ListRepoSyncTaskRequest) SetInstanceId(v string) *ListRepoSyncTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepoSyncTaskRequest) SetPageNo(v int32) *ListRepoSyncTaskRequest {
	s.PageNo = &v
	return s
}

func (s *ListRepoSyncTaskRequest) SetPageSize(v int32) *ListRepoSyncTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepoSyncTaskRequest) SetRepoName(v string) *ListRepoSyncTaskRequest {
	s.RepoName = &v
	return s
}

func (s *ListRepoSyncTaskRequest) SetRepoNamespaceName(v string) *ListRepoSyncTaskRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListRepoSyncTaskRequest) SetSyncRecordId(v string) *ListRepoSyncTaskRequest {
	s.SyncRecordId = &v
	return s
}

func (s *ListRepoSyncTaskRequest) SetTag(v string) *ListRepoSyncTaskRequest {
	s.Tag = &v
	return s
}

func (s *ListRepoSyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
