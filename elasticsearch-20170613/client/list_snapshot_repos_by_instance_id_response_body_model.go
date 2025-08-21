// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotReposByInstanceIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSnapshotReposByInstanceIdResponseBody
	GetRequestId() *string
	SetResult(v []*ListSnapshotReposByInstanceIdResponseBodyResult) *ListSnapshotReposByInstanceIdResponseBody
	GetResult() []*ListSnapshotReposByInstanceIdResponseBodyResult
}

type ListSnapshotReposByInstanceIdResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return results.
	Result []*ListSnapshotReposByInstanceIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListSnapshotReposByInstanceIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotReposByInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotReposByInstanceIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSnapshotReposByInstanceIdResponseBody) GetResult() []*ListSnapshotReposByInstanceIdResponseBodyResult {
	return s.Result
}

func (s *ListSnapshotReposByInstanceIdResponseBody) SetRequestId(v string) *ListSnapshotReposByInstanceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSnapshotReposByInstanceIdResponseBody) SetResult(v []*ListSnapshotReposByInstanceIdResponseBodyResult) *ListSnapshotReposByInstanceIdResponseBody {
	s.Result = v
	return s
}

func (s *ListSnapshotReposByInstanceIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSnapshotReposByInstanceIdResponseBodyResult struct {
	// Reference instance ID.
	//
	// example:
	//
	// es-cn-6ja1ro4jt000c****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The address of the repository.
	//
	// example:
	//
	// es-cn-6ja1ro4jt000c****
	RepoPath *string `json:"repoPath,omitempty" xml:"repoPath,omitempty"`
	// Reference warehouse name.
	//
	// example:
	//
	// aliyun_snapshot_from_es-cn-6ja1ro4jt000c****
	SnapWarehouse *string `json:"snapWarehouse,omitempty" xml:"snapWarehouse,omitempty"`
	// Reference warehouse status. available indicates that it is valid. unavailable indicates that it is invalid.
	//
	// example:
	//
	// available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListSnapshotReposByInstanceIdResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotReposByInstanceIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSnapshotReposByInstanceIdResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSnapshotReposByInstanceIdResponseBodyResult) GetRepoPath() *string {
	return s.RepoPath
}

func (s *ListSnapshotReposByInstanceIdResponseBodyResult) GetSnapWarehouse() *string {
	return s.SnapWarehouse
}

func (s *ListSnapshotReposByInstanceIdResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListSnapshotReposByInstanceIdResponseBodyResult) SetInstanceId(v string) *ListSnapshotReposByInstanceIdResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListSnapshotReposByInstanceIdResponseBodyResult) SetRepoPath(v string) *ListSnapshotReposByInstanceIdResponseBodyResult {
	s.RepoPath = &v
	return s
}

func (s *ListSnapshotReposByInstanceIdResponseBodyResult) SetSnapWarehouse(v string) *ListSnapshotReposByInstanceIdResponseBodyResult {
	s.SnapWarehouse = &v
	return s
}

func (s *ListSnapshotReposByInstanceIdResponseBodyResult) SetStatus(v string) *ListSnapshotReposByInstanceIdResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListSnapshotReposByInstanceIdResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
