// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSnapshotResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteSnapshotResponseBody
	GetResult() *bool
}

type DeleteSnapshotResponseBody struct {
	// example:
	//
	// 16484F36-A2A3-5A05-B242-0BF2BF6AA326
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSnapshotResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteSnapshotResponseBody) SetRequestId(v string) *DeleteSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSnapshotResponseBody) SetResult(v bool) *DeleteSnapshotResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteSnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}
