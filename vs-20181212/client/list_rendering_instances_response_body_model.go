// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRenderingInstances(v []*ListRenderingInstancesResponseBodyRenderingInstances) *ListRenderingInstancesResponseBody
	GetRenderingInstances() []*ListRenderingInstancesResponseBodyRenderingInstances
	SetRequestId(v string) *ListRenderingInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListRenderingInstancesResponseBody
	GetTotalCount() *int64
}

type ListRenderingInstancesResponseBody struct {
	RenderingInstances []*ListRenderingInstancesResponseBodyRenderingInstances `json:"RenderingInstances,omitempty" xml:"RenderingInstances,omitempty" type:"Repeated"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRenderingInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRenderingInstancesResponseBody) GetRenderingInstances() []*ListRenderingInstancesResponseBodyRenderingInstances {
	return s.RenderingInstances
}

func (s *ListRenderingInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRenderingInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRenderingInstancesResponseBody) SetRenderingInstances(v []*ListRenderingInstancesResponseBodyRenderingInstances) *ListRenderingInstancesResponseBody {
	s.RenderingInstances = v
	return s
}

func (s *ListRenderingInstancesResponseBody) SetRequestId(v string) *ListRenderingInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRenderingInstancesResponseBody) SetTotalCount(v int64) *ListRenderingInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRenderingInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRenderingInstancesResponseBodyRenderingInstances struct {
	// example:
	//
	// 2023-11-17T02:18:04Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// example:
	//
	// crs.cp.l1
	RenderingSpec *string `json:"RenderingSpec,omitempty" xml:"RenderingSpec,omitempty"`
	StorageSize   *int32  `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s ListRenderingInstancesResponseBodyRenderingInstances) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingInstancesResponseBodyRenderingInstances) GoString() string {
	return s.String()
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) GetRenderingSpec() *string {
	return s.RenderingSpec
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) SetCreationTime(v string) *ListRenderingInstancesResponseBodyRenderingInstances {
	s.CreationTime = &v
	return s
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) SetRenderingInstanceId(v string) *ListRenderingInstancesResponseBodyRenderingInstances {
	s.RenderingInstanceId = &v
	return s
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) SetRenderingSpec(v string) *ListRenderingInstancesResponseBodyRenderingInstances {
	s.RenderingSpec = &v
	return s
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) SetStorageSize(v int32) *ListRenderingInstancesResponseBodyRenderingInstances {
	s.StorageSize = &v
	return s
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) Validate() error {
	return dara.Validate(s)
}
