// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListChainRequest
	GetInstanceId() *string
	SetPageNo(v int32) *ListChainRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListChainRequest
	GetPageSize() *int32
	SetRepoName(v string) *ListChainRequest
	GetRepoName() *string
	SetRepoNamespaceName(v string) *ListChainRequest
	GetRepoNamespaceName() *string
}

type ListChainRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-4cdrlqmhn4gm****
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// repo1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// ns1
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s ListChainRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChainRequest) GoString() string {
	return s.String()
}

func (s *ListChainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListChainRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListChainRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChainRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *ListChainRequest) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *ListChainRequest) SetInstanceId(v string) *ListChainRequest {
	s.InstanceId = &v
	return s
}

func (s *ListChainRequest) SetPageNo(v int32) *ListChainRequest {
	s.PageNo = &v
	return s
}

func (s *ListChainRequest) SetPageSize(v int32) *ListChainRequest {
	s.PageSize = &v
	return s
}

func (s *ListChainRequest) SetRepoName(v string) *ListChainRequest {
	s.RepoName = &v
	return s
}

func (s *ListChainRequest) SetRepoNamespaceName(v string) *ListChainRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListChainRequest) Validate() error {
	return dara.Validate(s)
}
