// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEcsSpecsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEcsSpecs(v []*EcsSpec) *ListEcsSpecsResponseBody
	GetEcsSpecs() []*EcsSpec
	SetRequestId(v string) *ListEcsSpecsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListEcsSpecsResponseBody
	GetTotalCount() *int64
}

type ListEcsSpecsResponseBody struct {
	// The list of ECS specifications.
	EcsSpecs []*EcsSpec `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of types that meet the filter conditions.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEcsSpecsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEcsSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponseBody) GetEcsSpecs() []*EcsSpec {
	return s.EcsSpecs
}

func (s *ListEcsSpecsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEcsSpecsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListEcsSpecsResponseBody) SetEcsSpecs(v []*EcsSpec) *ListEcsSpecsResponseBody {
	s.EcsSpecs = v
	return s
}

func (s *ListEcsSpecsResponseBody) SetRequestId(v string) *ListEcsSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEcsSpecsResponseBody) SetTotalCount(v int64) *ListEcsSpecsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEcsSpecsResponseBody) Validate() error {
	return dara.Validate(s)
}
