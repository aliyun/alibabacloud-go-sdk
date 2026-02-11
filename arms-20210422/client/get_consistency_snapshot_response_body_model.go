// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsistencySnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConsistencyResultJsonStr(v string) *GetConsistencySnapshotResponseBody
	GetConsistencyResultJsonStr() *string
	SetRequestId(v string) *GetConsistencySnapshotResponseBody
	GetRequestId() *string
}

type GetConsistencySnapshotResponseBody struct {
	ConsistencyResultJsonStr *string `json:"ConsistencyResultJsonStr,omitempty" xml:"ConsistencyResultJsonStr,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConsistencySnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConsistencySnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsistencySnapshotResponseBody) GetConsistencyResultJsonStr() *string {
	return s.ConsistencyResultJsonStr
}

func (s *GetConsistencySnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConsistencySnapshotResponseBody) SetConsistencyResultJsonStr(v string) *GetConsistencySnapshotResponseBody {
	s.ConsistencyResultJsonStr = &v
	return s
}

func (s *GetConsistencySnapshotResponseBody) SetRequestId(v string) *GetConsistencySnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsistencySnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}
