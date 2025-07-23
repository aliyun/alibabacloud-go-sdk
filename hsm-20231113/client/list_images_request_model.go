// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *ListImagesRequest
	GetBackupId() *string
	SetCurrentPage(v int32) *ListImagesRequest
	GetCurrentPage() *int32
	SetMode(v string) *ListImagesRequest
	GetMode() *string
	SetPageSize(v int32) *ListImagesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListImagesRequest
	GetRegionId() *string
}

type ListImagesRequest struct {
	// The ID of the backup.
	//
	// This parameter is required.
	//
	// example:
	//
	// backup-fdb897sdf****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The image generation mode. Valid values:
	//
	// 	- PERIODIC: It is automatically generated.
	//
	// 	- MANUAL: It is manually generated.
	//
	// example:
	//
	// MANUAL
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The number of images per page. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *ListImagesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListImagesRequest) GetMode() *string {
	return s.Mode
}

func (s *ListImagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListImagesRequest) SetBackupId(v string) *ListImagesRequest {
	s.BackupId = &v
	return s
}

func (s *ListImagesRequest) SetCurrentPage(v int32) *ListImagesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListImagesRequest) SetMode(v string) *ListImagesRequest {
	s.Mode = &v
	return s
}

func (s *ListImagesRequest) SetPageSize(v int32) *ListImagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListImagesRequest) SetRegionId(v string) *ListImagesRequest {
	s.RegionId = &v
	return s
}

func (s *ListImagesRequest) Validate() error {
	return dara.Validate(s)
}
