// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSnapshotRepoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddSnapshotRepoResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddSnapshotRepoResponseBody
	GetResult() *bool
}

type AddSnapshotRepoResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return results:
	//
	// 	- true: Reference warehouse created successfully
	//
	// 	- false: Reference warehouse created failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s AddSnapshotRepoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSnapshotRepoResponseBody) GoString() string {
	return s.String()
}

func (s *AddSnapshotRepoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSnapshotRepoResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddSnapshotRepoResponseBody) SetRequestId(v string) *AddSnapshotRepoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSnapshotRepoResponseBody) SetResult(v bool) *AddSnapshotRepoResponseBody {
	s.Result = &v
	return s
}

func (s *AddSnapshotRepoResponseBody) Validate() error {
	return dara.Validate(s)
}
