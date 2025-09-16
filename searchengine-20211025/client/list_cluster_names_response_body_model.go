// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterNamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListClusterNamesResponseBody
	GetRequestId() *string
	SetResult(v *ListClusterNamesResponseBodyResult) *ListClusterNamesResponseBody
	GetResult() *ListClusterNamesResponseBodyResult
}

type ListClusterNamesResponseBody struct {
	// id of request
	//
	// example:
	//
	// F6E3D968-529C-5C40-AFDD-133A8B8FD930
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result set.
	Result *ListClusterNamesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ListClusterNamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterNamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterNamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterNamesResponseBody) GetResult() *ListClusterNamesResponseBodyResult {
	return s.Result
}

func (s *ListClusterNamesResponseBody) SetRequestId(v string) *ListClusterNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterNamesResponseBody) SetResult(v *ListClusterNamesResponseBodyResult) *ListClusterNamesResponseBody {
	s.Result = v
	return s
}

func (s *ListClusterNamesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClusterNamesResponseBodyResult struct {
	// The description of the cluster.
	//
	// example:
	//
	// ha3_test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// 25030
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// my_index
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListClusterNamesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListClusterNamesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListClusterNamesResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListClusterNamesResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListClusterNamesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListClusterNamesResponseBodyResult) SetDescription(v string) *ListClusterNamesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListClusterNamesResponseBodyResult) SetId(v int64) *ListClusterNamesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListClusterNamesResponseBodyResult) SetName(v string) *ListClusterNamesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListClusterNamesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
