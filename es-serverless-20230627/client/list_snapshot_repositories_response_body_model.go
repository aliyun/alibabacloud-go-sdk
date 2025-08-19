// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotRepositoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSnapshotRepositoriesResponseBody
	GetRequestId() *string
	SetResult(v []map[string]interface{}) *ListSnapshotRepositoriesResponseBody
	GetResult() []map[string]interface{}
}

type ListSnapshotRepositoriesResponseBody struct {
	// example:
	//
	// 56E0591D-7D62-56A2-993E-952FB2026C69
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListSnapshotRepositoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotRepositoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSnapshotRepositoriesResponseBody) GetResult() []map[string]interface{} {
	return s.Result
}

func (s *ListSnapshotRepositoriesResponseBody) SetRequestId(v string) *ListSnapshotRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSnapshotRepositoriesResponseBody) SetResult(v []map[string]interface{}) *ListSnapshotRepositoriesResponseBody {
	s.Result = v
	return s
}

func (s *ListSnapshotRepositoriesResponseBody) Validate() error {
	return dara.Validate(s)
}
