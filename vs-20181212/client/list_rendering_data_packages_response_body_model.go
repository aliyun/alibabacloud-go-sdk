// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingDataPackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataPackages(v []*ListRenderingDataPackagesResponseBodyDataPackages) *ListRenderingDataPackagesResponseBody
	GetDataPackages() []*ListRenderingDataPackagesResponseBodyDataPackages
	SetRequestId(v string) *ListRenderingDataPackagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListRenderingDataPackagesResponseBody
	GetTotalCount() *int64
}

type ListRenderingDataPackagesResponseBody struct {
	DataPackages []*ListRenderingDataPackagesResponseBodyDataPackages `json:"DataPackages,omitempty" xml:"DataPackages,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRenderingDataPackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingDataPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRenderingDataPackagesResponseBody) GetDataPackages() []*ListRenderingDataPackagesResponseBodyDataPackages {
	return s.DataPackages
}

func (s *ListRenderingDataPackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRenderingDataPackagesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRenderingDataPackagesResponseBody) SetDataPackages(v []*ListRenderingDataPackagesResponseBodyDataPackages) *ListRenderingDataPackagesResponseBody {
	s.DataPackages = v
	return s
}

func (s *ListRenderingDataPackagesResponseBody) SetRequestId(v string) *ListRenderingDataPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRenderingDataPackagesResponseBody) SetTotalCount(v int64) *ListRenderingDataPackagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRenderingDataPackagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRenderingDataPackagesResponseBodyDataPackages struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 2024-10-15T10:23:06+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// dp-449ea3d16c0841b8bf33ec5bbc86a152
	DataPackageId *string `json:"DataPackageId,omitempty" xml:"DataPackageId,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// render-342012a227dc4ddf91f024639e43051a
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-12-06T02:03:59Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListRenderingDataPackagesResponseBodyDataPackages) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingDataPackagesResponseBodyDataPackages) GoString() string {
	return s.String()
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) GetCategory() *string {
	return s.Category
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) GetDataPackageId() *string {
	return s.DataPackageId
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) GetDescription() *string {
	return s.Description
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) GetSize() *int32 {
	return s.Size
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) GetStatus() *string {
	return s.Status
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) SetCategory(v string) *ListRenderingDataPackagesResponseBodyDataPackages {
	s.Category = &v
	return s
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) SetCreationTime(v string) *ListRenderingDataPackagesResponseBodyDataPackages {
	s.CreationTime = &v
	return s
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) SetDataPackageId(v string) *ListRenderingDataPackagesResponseBodyDataPackages {
	s.DataPackageId = &v
	return s
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) SetDescription(v string) *ListRenderingDataPackagesResponseBodyDataPackages {
	s.Description = &v
	return s
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) SetRenderingInstanceId(v string) *ListRenderingDataPackagesResponseBodyDataPackages {
	s.RenderingInstanceId = &v
	return s
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) SetSize(v int32) *ListRenderingDataPackagesResponseBodyDataPackages {
	s.Size = &v
	return s
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) SetStatus(v string) *ListRenderingDataPackagesResponseBodyDataPackages {
	s.Status = &v
	return s
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) SetUpdateTime(v string) *ListRenderingDataPackagesResponseBodyDataPackages {
	s.UpdateTime = &v
	return s
}

func (s *ListRenderingDataPackagesResponseBodyDataPackages) Validate() error {
	return dara.Validate(s)
}
