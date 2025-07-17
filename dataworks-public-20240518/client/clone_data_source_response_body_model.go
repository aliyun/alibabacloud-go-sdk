// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CloneDataSourceResponseBody
	GetId() *int64
	SetRequestId(v string) *CloneDataSourceResponseBody
	GetRequestId() *string
}

type CloneDataSourceResponseBody struct {
	// The ID of the cloned data source.
	//
	// example:
	//
	// 19715
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request. It is used to locate logs and troubleshoot problems.
	//
	// example:
	//
	// FCD583B9-346B-5E75-82C1-4A7C192C48DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CloneDataSourceResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CloneDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneDataSourceResponseBody) SetId(v int64) *CloneDataSourceResponseBody {
	s.Id = &v
	return s
}

func (s *CloneDataSourceResponseBody) SetRequestId(v string) *CloneDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
