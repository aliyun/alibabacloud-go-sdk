// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoSyncRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListRepoSyncRuleRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *ListRepoSyncRuleRequest
	GetNamespaceName() *string
	SetPageNo(v int32) *ListRepoSyncRuleRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListRepoSyncRuleRequest
	GetPageSize() *int32
	SetRepoName(v string) *ListRepoSyncRuleRequest
	GetRepoName() *string
	SetTargetInstanceId(v string) *ListRepoSyncRuleRequest
	GetTargetInstanceId() *string
	SetTargetRegionId(v string) *ListRepoSyncRuleRequest
	GetTargetRegionId() *string
}

type ListRepoSyncRuleRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test-namespace
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
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
	// The name of the image repository.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The ID of the destination instance.
	//
	// example:
	//
	// cri-k77rd2eo9ztt****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The region ID of the destination instance.
	//
	// example:
	//
	// cn-shenzhen
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
}

func (s ListRepoSyncRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepoSyncRuleRequest) GoString() string {
	return s.String()
}

func (s *ListRepoSyncRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRepoSyncRuleRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListRepoSyncRuleRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRepoSyncRuleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRepoSyncRuleRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *ListRepoSyncRuleRequest) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *ListRepoSyncRuleRequest) GetTargetRegionId() *string {
	return s.TargetRegionId
}

func (s *ListRepoSyncRuleRequest) SetInstanceId(v string) *ListRepoSyncRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepoSyncRuleRequest) SetNamespaceName(v string) *ListRepoSyncRuleRequest {
	s.NamespaceName = &v
	return s
}

func (s *ListRepoSyncRuleRequest) SetPageNo(v int32) *ListRepoSyncRuleRequest {
	s.PageNo = &v
	return s
}

func (s *ListRepoSyncRuleRequest) SetPageSize(v int32) *ListRepoSyncRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepoSyncRuleRequest) SetRepoName(v string) *ListRepoSyncRuleRequest {
	s.RepoName = &v
	return s
}

func (s *ListRepoSyncRuleRequest) SetTargetInstanceId(v string) *ListRepoSyncRuleRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *ListRepoSyncRuleRequest) SetTargetRegionId(v string) *ListRepoSyncRuleRequest {
	s.TargetRegionId = &v
	return s
}

func (s *ListRepoSyncRuleRequest) Validate() error {
	return dara.Validate(s)
}
