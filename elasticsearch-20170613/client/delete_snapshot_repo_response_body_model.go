// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotRepoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSnapshotRepoResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteSnapshotRepoResponseBody
	GetResult() *bool
}

type DeleteSnapshotRepoResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return results:
	//
	// 	- true: reference warehouse deleted successfully
	//
	// 	- false: reference warehouse deleted successfully failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteSnapshotRepoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotRepoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRepoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSnapshotRepoResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteSnapshotRepoResponseBody) SetRequestId(v string) *DeleteSnapshotRepoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSnapshotRepoResponseBody) SetResult(v bool) *DeleteSnapshotRepoResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteSnapshotRepoResponseBody) Validate() error {
	return dara.Validate(s)
}
