// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChainInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListChainInstanceRequest
	GetInstanceId() *string
	SetPageNo(v int32) *ListChainInstanceRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListChainInstanceRequest
	GetPageSize() *int32
	SetRepoName(v string) *ListChainInstanceRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *ListChainInstanceRequest
	GetRepoNamespaceName() *string
}

type ListChainInstanceRequest struct {
	// The operation that you want to perform. Set this parameter to **ListChainInstance**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the delivery chain started.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The time when the delivery chain is completed.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the delivery chain.
	//
	// example:
	//
	// test-namespace
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s ListChainInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChainInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListChainInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListChainInstanceRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListChainInstanceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChainInstanceRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *ListChainInstanceRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *ListChainInstanceRequest) SetInstanceId(v string) *ListChainInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListChainInstanceRequest) SetPageNo(v int32) *ListChainInstanceRequest {
	s.PageNo = &v
	return s
}

func (s *ListChainInstanceRequest) SetPageSize(v int32) *ListChainInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListChainInstanceRequest) SetRepoName(v string) *ListChainInstanceRequest {
	s.RepoName = &v
	return s
}

func (s *ListChainInstanceRequest) SetRepoNamespaceName(v string) *ListChainInstanceRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListChainInstanceRequest) Validate() error {
	return dara.Validate(s)
}
