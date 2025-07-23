// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListImagesResponseBody
	GetCurrentPage() *int32
	SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody
	GetImages() []*ListImagesResponseBodyImages
	SetPageSize(v int32) *ListImagesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListImagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListImagesResponseBody
	GetTotalCount() *int32
}

type ListImagesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The images returned.
	Images []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The number of images per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of images returned.
	//
	// example:
	//
	// 1000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListImagesResponseBody) GetImages() []*ListImagesResponseBodyImages {
	return s.Images
}

func (s *ListImagesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListImagesResponseBody) SetCurrentPage(v int32) *ListImagesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListImagesResponseBody) SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetPageSize(v int32) *ListImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) SetTotalCount(v int32) *ListImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListImagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListImagesResponseBodyImages struct {
	// The ID of the backup.
	//
	// example:
	//
	// backup-fdb897sdf****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The time when the image was copied. Unit: milliseconds. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1641275680000
	CopyTime *string `json:"CopyTime,omitempty" xml:"CopyTime,omitempty"`
	// The time when the image was generated. Unit: milliseconds. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1782849566738
	ExportTime *int64 `json:"ExportTime,omitempty" xml:"ExportTime,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// image-d79x4k11pmg19****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the hardware security module (HSM).
	//
	// example:
	//
	// hsm-cn-6ja1xknf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the backup.
	//
	// example:
	//
	// hsm-test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the source backup.
	//
	// example:
	//
	// backup-hodfhaol****
	SourceBackupUid *string `json:"SourceBackupUid,omitempty" xml:"SourceBackupUid,omitempty"`
	// The ID of the source image.
	//
	// example:
	//
	// image-ooopjygsn****
	SourceImageUid *string `json:"SourceImageUid,omitempty" xml:"SourceImageUid,omitempty"`
	// The ID of the source HSM.
	//
	// example:
	//
	// hsm-cn-wz9i2dmefudfxtmb****
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The ID of the region in which the source image resides.
	//
	// example:
	//
	// cn-shanghai
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// The status of the image. Valid values:
	//
	// 	- NEW: It is disabled.
	//
	// 	- DELETED: It is deleted.
	//
	// 	- CREATING: It is being created.
	//
	// 	- NORMAL: It is created.
	//
	// example:
	//
	// CREATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The digest of the HSM.
	//
	// example:
	//
	// 3kGeHnmQzXwSsfF0Jk9eJYhe2gP6An0/HlYIiZh1****
	VsmDigest *string `json:"VsmDigest,omitempty" xml:"VsmDigest,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) GetBackupId() *string {
	return s.BackupId
}

func (s *ListImagesResponseBodyImages) GetCopyTime() *string {
	return s.CopyTime
}

func (s *ListImagesResponseBodyImages) GetExportTime() *int64 {
	return s.ExportTime
}

func (s *ListImagesResponseBodyImages) GetImageId() *string {
	return s.ImageId
}

func (s *ListImagesResponseBodyImages) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListImagesResponseBodyImages) GetMode() *string {
	return s.Mode
}

func (s *ListImagesResponseBodyImages) GetRegionId() *string {
	return s.RegionId
}

func (s *ListImagesResponseBodyImages) GetRemark() *string {
	return s.Remark
}

func (s *ListImagesResponseBodyImages) GetSourceBackupUid() *string {
	return s.SourceBackupUid
}

func (s *ListImagesResponseBodyImages) GetSourceImageUid() *string {
	return s.SourceImageUid
}

func (s *ListImagesResponseBodyImages) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *ListImagesResponseBodyImages) GetSourceRegionId() *string {
	return s.SourceRegionId
}

func (s *ListImagesResponseBodyImages) GetStatus() *string {
	return s.Status
}

func (s *ListImagesResponseBodyImages) GetVsmDigest() *string {
	return s.VsmDigest
}

func (s *ListImagesResponseBodyImages) SetBackupId(v string) *ListImagesResponseBodyImages {
	s.BackupId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetCopyTime(v string) *ListImagesResponseBodyImages {
	s.CopyTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetExportTime(v int64) *ListImagesResponseBodyImages {
	s.ExportTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageId(v string) *ListImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetInstanceId(v string) *ListImagesResponseBodyImages {
	s.InstanceId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetMode(v string) *ListImagesResponseBodyImages {
	s.Mode = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetRegionId(v string) *ListImagesResponseBodyImages {
	s.RegionId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetRemark(v string) *ListImagesResponseBodyImages {
	s.Remark = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSourceBackupUid(v string) *ListImagesResponseBodyImages {
	s.SourceBackupUid = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSourceImageUid(v string) *ListImagesResponseBodyImages {
	s.SourceImageUid = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSourceInstanceId(v string) *ListImagesResponseBodyImages {
	s.SourceInstanceId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSourceRegionId(v string) *ListImagesResponseBodyImages {
	s.SourceRegionId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetStatus(v string) *ListImagesResponseBodyImages {
	s.Status = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetVsmDigest(v string) *ListImagesResponseBodyImages {
	s.VsmDigest = &v
	return s
}

func (s *ListImagesResponseBodyImages) Validate() error {
	return dara.Validate(s)
}
