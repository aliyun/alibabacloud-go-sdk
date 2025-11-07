// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGitBranchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListGitBranchesResponseBody
	GetCount() *int32
	SetGitBranches(v []*ListGitBranchesResponseBodyGitBranches) *ListGitBranchesResponseBody
	GetGitBranches() []*ListGitBranchesResponseBodyGitBranches
	SetRequestId(v string) *ListGitBranchesResponseBody
	GetRequestId() *string
}

type ListGitBranchesResponseBody struct {
	// example:
	//
	// 1
	Count       *int32                                    `json:"Count,omitempty" xml:"Count,omitempty"`
	GitBranches []*ListGitBranchesResponseBodyGitBranches `json:"GitBranches,omitempty" xml:"GitBranches,omitempty" type:"Repeated"`
	// example:
	//
	// DBA6E6C8-F75D-41DE-AFF5-1FA03F551CA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGitBranchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGitBranchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGitBranchesResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListGitBranchesResponseBody) GetGitBranches() []*ListGitBranchesResponseBodyGitBranches {
	return s.GitBranches
}

func (s *ListGitBranchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGitBranchesResponseBody) SetCount(v int32) *ListGitBranchesResponseBody {
	s.Count = &v
	return s
}

func (s *ListGitBranchesResponseBody) SetGitBranches(v []*ListGitBranchesResponseBodyGitBranches) *ListGitBranchesResponseBody {
	s.GitBranches = v
	return s
}

func (s *ListGitBranchesResponseBody) SetRequestId(v string) *ListGitBranchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGitBranchesResponseBody) Validate() error {
	if s.GitBranches != nil {
		for _, item := range s.GitBranches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGitBranchesResponseBodyGitBranches struct {
	// example:
	//
	// master
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListGitBranchesResponseBodyGitBranches) String() string {
	return dara.Prettify(s)
}

func (s ListGitBranchesResponseBodyGitBranches) GoString() string {
	return s.String()
}

func (s *ListGitBranchesResponseBodyGitBranches) GetName() *string {
	return s.Name
}

func (s *ListGitBranchesResponseBodyGitBranches) SetName(v string) *ListGitBranchesResponseBodyGitBranches {
	s.Name = &v
	return s
}

func (s *ListGitBranchesResponseBodyGitBranches) Validate() error {
	return dara.Validate(s)
}
