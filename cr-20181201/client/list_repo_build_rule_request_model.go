// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoBuildRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListRepoBuildRuleRequest
	GetInstanceId() *string
	SetPageNo(v int32) *ListRepoBuildRuleRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListRepoBuildRuleRequest
	GetPageSize() *int32
	SetRepoId(v string) *ListRepoBuildRuleRequest
	GetRepoId() *string
}

type ListRepoBuildRuleRequest struct {
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
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-tquyps22md8****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s ListRepoBuildRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepoBuildRuleRequest) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRepoBuildRuleRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRepoBuildRuleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRepoBuildRuleRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *ListRepoBuildRuleRequest) SetInstanceId(v string) *ListRepoBuildRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepoBuildRuleRequest) SetPageNo(v int32) *ListRepoBuildRuleRequest {
	s.PageNo = &v
	return s
}

func (s *ListRepoBuildRuleRequest) SetPageSize(v int32) *ListRepoBuildRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepoBuildRuleRequest) SetRepoId(v string) *ListRepoBuildRuleRequest {
	s.RepoId = &v
	return s
}

func (s *ListRepoBuildRuleRequest) Validate() error {
	return dara.Validate(s)
}
