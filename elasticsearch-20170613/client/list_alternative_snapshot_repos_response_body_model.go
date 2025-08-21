// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlternativeSnapshotReposResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAlternativeSnapshotReposResponseBody
	GetRequestId() *string
	SetResult(v []*ListAlternativeSnapshotReposResponseBodyResult) *ListAlternativeSnapshotReposResponseBody
	GetResult() []*ListAlternativeSnapshotReposResponseBodyResult
}

type ListAlternativeSnapshotReposResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return results.
	Result []*ListAlternativeSnapshotReposResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListAlternativeSnapshotReposResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlternativeSnapshotReposResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlternativeSnapshotReposResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlternativeSnapshotReposResponseBody) GetResult() []*ListAlternativeSnapshotReposResponseBodyResult {
	return s.Result
}

func (s *ListAlternativeSnapshotReposResponseBody) SetRequestId(v string) *ListAlternativeSnapshotReposResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlternativeSnapshotReposResponseBody) SetResult(v []*ListAlternativeSnapshotReposResponseBodyResult) *ListAlternativeSnapshotReposResponseBody {
	s.Result = v
	return s
}

func (s *ListAlternativeSnapshotReposResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAlternativeSnapshotReposResponseBodyResult struct {
	// The ID of the instance.
	//
	// example:
	//
	// es-cn-6ja1ro4jt000c****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The address of the repository.
	//
	// example:
	//
	// RepoPath
	RepoPath *string `json:"repoPath,omitempty" xml:"repoPath,omitempty"`
}

func (s ListAlternativeSnapshotReposResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAlternativeSnapshotReposResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAlternativeSnapshotReposResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAlternativeSnapshotReposResponseBodyResult) GetRepoPath() *string {
	return s.RepoPath
}

func (s *ListAlternativeSnapshotReposResponseBodyResult) SetInstanceId(v string) *ListAlternativeSnapshotReposResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListAlternativeSnapshotReposResponseBodyResult) SetRepoPath(v string) *ListAlternativeSnapshotReposResponseBodyResult {
	s.RepoPath = &v
	return s
}

func (s *ListAlternativeSnapshotReposResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
