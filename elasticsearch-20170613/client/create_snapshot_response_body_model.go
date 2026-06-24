// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSnapshotResponseBody
	GetRequestId() *string
	SetResult(v bool) *CreateSnapshotResponseBody
	GetResult() *bool
}

type CreateSnapshotResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result. Valid values:
	//
	// - true: The manual snapshot backup is successful.
	//
	// - false: The manual snapshot backup failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSnapshotResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CreateSnapshotResponseBody) SetRequestId(v string) *CreateSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetResult(v bool) *CreateSnapshotResponseBody {
	s.Result = &v
	return s
}

func (s *CreateSnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}
