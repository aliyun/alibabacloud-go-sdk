// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingDataPackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListRenderingDataPackagesRequest
	GetCategory() *string
	SetDataPackageId(v string) *ListRenderingDataPackagesRequest
	GetDataPackageId() *string
	SetEndTime(v string) *ListRenderingDataPackagesRequest
	GetEndTime() *string
	SetPageNumber(v int32) *ListRenderingDataPackagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRenderingDataPackagesRequest
	GetPageSize() *int32
	SetSize(v int32) *ListRenderingDataPackagesRequest
	GetSize() *int32
	SetStartTime(v string) *ListRenderingDataPackagesRequest
	GetStartTime() *string
	SetStatus(v string) *ListRenderingDataPackagesRequest
	GetStatus() *string
}

type ListRenderingDataPackagesRequest struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// dp-449ea3d16c0841b8bf33ec5bbc86a152
	DataPackageId *string `json:"DataPackageId,omitempty" xml:"DataPackageId,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 20
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListRenderingDataPackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingDataPackagesRequest) GoString() string {
	return s.String()
}

func (s *ListRenderingDataPackagesRequest) GetCategory() *string {
	return s.Category
}

func (s *ListRenderingDataPackagesRequest) GetDataPackageId() *string {
	return s.DataPackageId
}

func (s *ListRenderingDataPackagesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListRenderingDataPackagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRenderingDataPackagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRenderingDataPackagesRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListRenderingDataPackagesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListRenderingDataPackagesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListRenderingDataPackagesRequest) SetCategory(v string) *ListRenderingDataPackagesRequest {
	s.Category = &v
	return s
}

func (s *ListRenderingDataPackagesRequest) SetDataPackageId(v string) *ListRenderingDataPackagesRequest {
	s.DataPackageId = &v
	return s
}

func (s *ListRenderingDataPackagesRequest) SetEndTime(v string) *ListRenderingDataPackagesRequest {
	s.EndTime = &v
	return s
}

func (s *ListRenderingDataPackagesRequest) SetPageNumber(v int32) *ListRenderingDataPackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRenderingDataPackagesRequest) SetPageSize(v int32) *ListRenderingDataPackagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRenderingDataPackagesRequest) SetSize(v int32) *ListRenderingDataPackagesRequest {
	s.Size = &v
	return s
}

func (s *ListRenderingDataPackagesRequest) SetStartTime(v string) *ListRenderingDataPackagesRequest {
	s.StartTime = &v
	return s
}

func (s *ListRenderingDataPackagesRequest) SetStatus(v string) *ListRenderingDataPackagesRequest {
	s.Status = &v
	return s
}

func (s *ListRenderingDataPackagesRequest) Validate() error {
	return dara.Validate(s)
}
